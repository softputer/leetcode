package main 

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = strings.ToUpper(s)
	r := []rune(s)
	start, end := 0, len(r)-1
	for start < end {
		if !isValid(r[start]) {
			start++
			continue
		}
		if !isValid(r[end]) {
			end--
			continue
		}
		if r[start] == r[end] {
			start++
			end--
			continue
		}
		return false
	} 
	return true
}

func isValid(r rune) bool {
	low, up := 'A', 'Z'
	low2, up2 := '0', '9' 
	if r <= up && r >= low {
		return true
	}  	
	if r <= up2 && r >= low2 {
		return true
	}  	
	return false
}

func main() {
	s := "aba"		
	fmt.Println(isPalindrome(s))
}
