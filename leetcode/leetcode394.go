package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func get_decoded_string(text string, estring []int) string {
	decoded_msg := ""

	re := regexp.MustCompile("[0-9]+")
	num, _ := strconv.Atoi(re.FindAllString(string(text[estring[0]:estring[1]]), -1)[0])
	emsg := text[estring[0]:(estring[1] - 1)]

	emsg = strings.Split(emsg, fmt.Sprintf("%d%s", num, "["))[1]
	fmt.Println(num, emsg)

	for i := 0; i < num; i++ {
		decoded_msg += emsg
	}

	fmt.Println(decoded_msg)
	return decoded_msg
}

func main() {
	delimeter := "[0-9]+\\[[a-zA-Z]+\\]"
	text := "3[a]2[b4[F]c]"

	for {
		reg := regexp.MustCompile(delimeter)
		indexes := reg.FindAllStringIndex(text, -1)
		fmt.Println(indexes)

		if len(indexes) > 0 {
			estring := indexes[0]
			text = strings.Replace(text, text[estring[0]:estring[1]], get_decoded_string(text, estring), 1)
			fmt.Println("Updated text : ", text)

		} else {
			break
		}
	}

	fmt.Println("final text : ", text)

}
