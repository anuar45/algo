package solution

func singleNumber(nums []int) int {
	m := make(map[int]int)
	var one int
	for _, n := range nums {
		m[n]++
	}

	for k, v := range m {
		if v == 1 {
			one = k
			break
		}
	}
	return one
}
