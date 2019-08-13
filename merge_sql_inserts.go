package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func extract_create(file *os.File) ([]string, []string) {
	scanner := bufio.NewScanner(file)

	create_scripts := ""
	script_start := false

	create_tables := []string{}
	table_names := []string{}

	for scanner.Scan() { // internally, it advances token based on sperator
		// fmt.Println(scanner.Text()) // token in unicode-char
		scannerText := scanner.Text()
		if script_start {
			match, _ := regexp.MatchString("ENGINE=innodb", scannerText)
			if match {
				script_start = false
				create_scripts += scannerText
				create_tables = append(create_tables, create_scripts)
			} else {
				create_scripts += scannerText
			}
		} else {
			match, _ := regexp.MatchString("CREATE TABLE", scannerText)
			if match {
				rexp := regexp.MustCompile("`.*?`")
				rexpmatch := rexp.FindStringSubmatch(scannerText)
				table_names = append(table_names, strings.Split(rexpmatch[0], "`")[1])
				script_start = true
				create_scripts = ""
				create_scripts += scannerText
			}
		}
	}

	return create_tables, table_names
}

func extractInsert(file *os.File, tableName string, page_limit int) (string, int) {
	scanner := bufio.NewScanner(file)
	merge_inserts := []string{}
	script_base := ""

	terminate := -1
	for scanner.Scan() { // internally, it advances token based on sperator
		// fmt.Println(scanner.Text()) // token in unicode-char
		scannerText := scanner.Text()

		match, _ := regexp.MatchString("INSERT INTO `"+tableName+"`", scannerText)
		if match {
			// fmt.Println(match, scannerText)
			terminate = 0
			sarr := strings.Split(scannerText, " VALUES ")
			script_base = sarr[0]
			merge_inserts = append(merge_inserts, strings.Split(sarr[1], ";")[0])
			page_limit--
		} else if terminate >= 0 {
			terminate++
			if terminate > 10 {
				break
			}
		}
		if page_limit <= 0 {
			break
		}
	}

	sql_script := script_base + " VALUES " + strings.Join(merge_inserts, ",") + ";"

	fmt.Println("Merged Inserts : ", len(merge_inserts))
	if len(merge_inserts) > 0 {
		fmt.Println("Merged Inserts 0 Val : ", merge_inserts[0])
		fmt.Println("Merged Inserts Last Val : ", merge_inserts[len(merge_inserts)-1])
	}

	return sql_script, len(merge_inserts)
}

func main() {
	file, err := os.Open("dump.sql")
	if err != nil {
		log.Fatal(err)
	}

	//  Extract Table names from sql
	create_tables, table_names := extract_create(file)
	file.Close()
	if len(create_tables) > 0 {
		fo, err := os.Create("create_tables.sql")
		if err != nil {
			panic(err)
		}
		// close fo on exit and check for its returned error
		defer func() {
			if err := fo.Close(); err != nil {
				panic(err)
			}
		}()

		// make a write buffer
		w := bufio.NewWriter(fo)

		// make a buffer to keep chunks that are read
		buf := []byte(strings.Join(create_tables, "\n"))
		// write a chunk
		if _, err := w.Write(buf); err != nil {
			panic(err)
		}

		if err = w.Flush(); err != nil {
			panic(err)
		}
	}

	for _, csv := range table_names {
		fmt.Println(csv)
	}

	for _, table := range table_names {
		file, err := os.Open("dump.sql")
		if err != nil {
			log.Fatal(err)
		}
		fo, err := os.Create(table + ".sql")
		if err != nil {
			panic(err)
		}
		// close fo on exit and check for its returned error
		defer func() {
			if err := fo.Close(); err != nil {
				panic(err)
			}
		}()

		pidx := 0
		qcount := 0
		fmt.Println("Table Name : ", table)
		for {
			fmt.Println("Iteration : ", pidx)
			query, rcount := extractInsert(file, table, 1000)
			if rcount > 0 {
				pidx++
				// make a write buffer
				w := bufio.NewWriter(fo)

				// make a buffer to keep chunks that are read
				buf := []byte(query)
				// write a chunk
				if _, err := w.Write(buf); err != nil {
					panic(err)
				}

				if err = w.Flush(); err != nil {
					panic(err)
				}
				qcount += rcount
			} else {
				break
			}
			fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		}

		fmt.Println(qcount)
		file.Close()
	}
}
