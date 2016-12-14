package main

import (
	"fmt"
)

func findKth(nums1 []int, nums2 []int, k int) int {	
	if len(nums1) == 0 {
		return nums2[k-1]
	}
	if len(nums2) == 0 {
		return nums1[k-1]
	}
	if k == 1 {
		if nums1[k-1] > nums2[k-1] {
			return nums2[k-1]
		} else {
			return nums1[k-1]
		}
	}

	var a, b interface{}
	if len(nums1) >= k/2 {
		a = nums1[k/2-1]	
	} else {
		a = nil
	}
	if len(nums2) >= k/2 {
		b = nums2[k/2-1]	
	} else {
		b = nil
	}
	if b == nil || ( a != nil && a.(int) < b.(int)) {
		return findKth(nums1[k/2:], nums2, k-k/2)
	} else {
		return findKth(nums1, nums2[k/2:], k-k/2)
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	if length%2 == 1 {
		return float64(findKth(nums1, nums2, length/2 + 1))
	} else {
		return (float64)((findKth(nums1, nums2, length/2) + findKth(nums1, nums2, length/2 +1)))/2.0 
	}
}

func main() {
	a := []int{0, 1, 2, 3, 4, 9}
	b := []int{3, 4, 5, 6}
	result := findMedianSortedArrays(a, b)
	fmt.Println(result)
}
