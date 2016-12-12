package main

import (
	"fmt"	
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var m map[int]int = make(map[int]int)
func buildTree(preorder []int, inorder []int) *TreeNode {	
	preStart, preEnd, inStart, inEnd := 0, len(preorder)-1, 0, len(inorder)-1
	for i := 0; i < len(inorder); i++ {
		m[inorder[i]] = i
	}
	return build(preorder, inorder, preStart, preEnd, inStart, inEnd)
}

func build(preorder []int, inorder []int, preStart, preEnd, inStart, inEnd int) *TreeNode {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}
	root := &TreeNode{Val:preorder[preStart]}
	
	mid := m[preorder[preStart]]	
	num := mid - inStart

	root.Left = build(preorder, inorder, preStart+1, preStart+num, inStart, mid-1)
	root.Right = build(preorder, inorder, preStart+num+1, preEnd, mid+1, inEnd)

	return root
}

func main() {
	preorder := []int{1, 2, 4, 5, 3, 6, 7}	
	inorder := []int{4, 2, 5, 1, 6, 3, 7}	
	result := buildTree(preorder, inorder)	
	fmt.Println(result)
}
