package solution

func intersect(nums1 []int, nums2 []int) []int {
	var small, big, intersect, maxIntersect []int

	if len(nums1) > len(nums2) {
		small = nums2
		big = nums1
	}

	for i := 0; i < len(big); i++ {
		for j := 0; j < len(small) && i+j < len(big); j++ {
			if small[j] == big[i+j] {
				intersect = append(intersect, small[j])
			} else {
				if len(intersect) > len(maxIntersect) {
					maxIntersect = intersect
				}
				intersect = nil
			}
		}
	}
	return maxIntersect
}
