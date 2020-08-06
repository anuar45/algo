package solution

func moveZeros(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			moveToEnd(nums, i)
		}
	}
}

func moveToEnd(nums []int, index int) {
	for i := index; i < len(nums)-1; i++ {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}
}
