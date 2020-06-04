package solution

func intersect(nums1 []int, nums2 []int) []int {
	var intersect []int

	m1 := make(map[int]int)
	for _, n := range nums1 {
		m1[n]++
	}

	for _, n := range nums2 {
		if m1[n] != 0 {
			intersect = append(intersect, n)
			m1[n]--
		}
	}

	return intersect
}
