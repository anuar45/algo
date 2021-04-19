package main

import (
	"fmt"
)

func main() {
	fmt.Println(moveZeroes([]int{1, 0, 2, 0, 0, 3}))
	fmt.Println(moveZeroes([]int{0, 1, 0, 3, 12}))

}

func moveZeroes(nums []int) []int {
	var zeroesLeft bool
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			for k := i + 1; k < len(nums); k++ {
				if nums[k] != 0 {
					nums[i], nums[k] = nums[k], nums[i]
					break
				}
				if k == len(nums)-1 {
					zeroesLeft = true
				}
			}
		}
		if zeroesLeft {
			return nums
		}
	}
	return nums
}

// 1,0,0,3,12
//
