package main

import (
	"fmt"
	"math"
)

var (
	MAX = math.MaxInt64
)


func max_profit(prices []int) int {

	profit := 0
	buy_val := MAX 
	for _, v := range prices {
		if buy_val>v {
			buy_val=v
		}


		if buy_val<v {
			profit+=(v-buy_val)
			buy_val = v
		}
	}

	return profit
}

func main() {
	prices := []int {7,1,5,3,6,4}

	fmt.Println(max_profit(prices))
}
