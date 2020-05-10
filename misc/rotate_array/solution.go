package rotate

func rotate(nums []int, k int) {

	k %= len(nums)

	reverse(nums, 0, len(nums)-1)
	//fmt.Println(nums)
	reverse(nums, 0, k-1)
	//fmt.Println(nums)
	reverse(nums, k, len(nums)-1)
	//fmt.Println(nums)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
