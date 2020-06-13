package sort

import "fmt"

func partitition(nums []int, l, h int) int {
	pivot := nums[l]

	i, j := l+1, h

	fmt.Println(nums, l, h)
	for i < j {

		for nums[i] <= pivot {
			i++
		}

		for nums[j] > pivot {
			j--
			fmt.Println("debug j", j)
			fmt.Println("debug pivot", pivot)
		}

		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}

	}
	nums[j], nums[l] = nums[l], nums[j]
	fmt.Println("debug1", nums)
	return j
}

// Quick is simple quick sort implmentation
func Quick(nums []int, l, h int) {
	if l < h {
		j := partitition(nums, l, h)
		fmt.Println(j)

		Quick(nums, l, j-1)
		Quick(nums, j+1, h)
	}

}
