package main

import (
	"fmt"
)

// func twoSum(nums []int, target int) []int {

// 	result := []int{}

// 	for i, n := range nums {
// 		var k = i
// 		sum := 0
// 		result = result[:0]
// 		if n <= target {
// 			for k < len(nums) {
// 				sum += nums[k]
// 				if sum < target {
// 					result = append(result, k)
// 					k++
// 				} else if sum > target {
// 					sum -= nums[k]
// 					k++
// 				} else {
// 					result = append(result, k)
// 					return result
// 				}
// 			}
// 		}
// 	}

// 	return result

// }
func twoSum(nums []int, target int) []int {
	result := make(map[int]int)
	remaining := 0

	for i, v := range nums {
		remaining = target - v
		if _, ok := result[remaining]; ok {

			return []int{result[remaining], i}
		}

		result[v] = i
	}

	return []int{}
}

func main() {
	// nums := []int{3, 2, 3}

	// target := 6

	nums := []int{2, 7, 11, 15}

	target := 9

	fmt.Println(twoSum(nums, target))
}
