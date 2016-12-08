package main

import (
	"fmt"
)


func main() {
	arr := []int{0, 1, 2, 1, 0, 2}
	sortColors(arr)
}

func sortColors(nums []int)  {
        length := len(num)
        zero, two, i := 0, length - 1, 0
        for i <= two{
                if num[i] == 0 {
        //              fmt.Println("swap", num[i], num[zero])
                        num[i], num[zero] = num[zero], num[i]
                        zero++
                        i++
                        continue
                }
                if num[i] == 2 {
        //              fmt.Println("swap", num[i], num[two])
                        num[i], num[two] = num[two], num[i]
                        two--
                        continue
                }
                i++
        }
