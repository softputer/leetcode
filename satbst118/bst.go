package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	num := []int{0, 1, 2, 3}	
	result := sortedArrayToBST(num)
	fmt.Println(result)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Left
		if result == nil {
			fmt.Println("nil")
		}
	}
}

func sortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)
	if length <= 0 {
		return nil
	}	
	var node *TreeNode = &TreeNode{Val:nums[length/2]}
	node.Left = sortedArrayToBST(nums[:length/2])
	node.Right = sortedArrayToBST(nums[length/2+1:])
	return node	
}
