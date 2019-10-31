package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i]==nums[j] {
				nums = append(nums[:j],nums[(j+1):]...)
				j--
			}
		}
	}
	return len(nums)
}

func main(){
	nums := []int{1,1,2}


	fmt.Println(removeDuplicates(nums))

} 