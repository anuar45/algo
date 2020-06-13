package sort

// Merge simple two-way merge
func Merge(nums1, nums2 []int) []int {
	var i1, i2, i3 int
	nums3 := make([]int, len(nums1)+len(nums2))

	for i1 < len(nums1) && i2 < len(nums2) {
		if nums1[i1] < nums2[i2] {
			nums3[i3] = nums1[i1]
			i1++
		} else {
			nums3[i3] = nums2[i2]
			i2++
		}

		i3++
	}

	for i1 < len(nums1) {
		nums3[i3] = nums1[i1]
		i3++
		i1++
	}

	for i2 < len(nums2) {
		nums3[i3] = nums2[i2]
		i3++
		i2++
	}

	return nums3
}

// MergeSort sorting using two-way merge
func MergeSort(nums []int) []int {
	var m1, m2 []int

	if len(nums) == 1 {
		return nums
	}

	mid := len(nums) / 2
	m1 = MergeSort(nums[:mid])
	m2 = MergeSort(nums[mid:])
	return Merge(m1, m2)
}
