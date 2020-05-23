package solution

func containsDuplicate(nums []int) bool {
	var dup bool
	m := make(map[int]struct{})

	for _, v := range nums {
		if _, exists := m[v]; !exists {
			m[v] = struct{}{}
		} else {
			dup = true
		}
	}
	return dup
}
