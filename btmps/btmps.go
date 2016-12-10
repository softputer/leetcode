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
	num := []int{-3}	
	result := sortedArrayToBST(num)
	fmt.Println(result)
	maxLen := maxPathSum(result)	
	fmt.Println(maxLen)
}

func maxPathSum(root *TreeNode) int {
	maxsum, _ := maxPath(root)
	return maxsum
}

func maxPath(root *TreeNode) (maxsum, single int) {
	if root == nil {
		return -9999999, 0
	}  
	leftm, lefts := maxPath(root.Left)
	rightm, rights := maxPath(root.Right)
	maxsum = max(leftm, rightm, lefts+rights+root.Val)	
	fmt.Println(maxsum)
	single = max(lefts+root.Val, rights+root.Val, 0)
	return
}

func max(a, b, c int) (result int) {
	if a > b {
		result = a
	} else {
		result = b
	}
	if c > result {
		result = c
	}
	return
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
