package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return 1
	}
	i := 0
	for j := 1; j < length; j++{
		if nums[j] != nums[j-1] {
			i++		
		}
	}
	return i+1
}

func main() {
	nums := []int{0, 1, 1, 2}
	result := removeDuplicates(nums)
	fmt.Println(result)
}
