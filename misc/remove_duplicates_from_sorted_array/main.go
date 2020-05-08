package main

import "fmt"

func main() {

	fmt.Println(removeDuplicates([]int{1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 5, 5}))

}

func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	i, j := 0, 1
	for j < len(nums) {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
		j++
	}
	return len(nums[:i+1])
}
