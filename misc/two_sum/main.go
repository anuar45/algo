package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for k := i + 1; k < len(nums); k++ {
			if nums[i]+nums[k] == target {
				return []int{i, k}
			}
		}
	}
	return []int{}
}
