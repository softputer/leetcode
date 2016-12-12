package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		result ^= nums[i] 
	}
	return result
}

func main() {
	nums := []int{1, 1, 2, 2, 0}
	num := singleNumber(nums)
	fmt.Println(num)
}
