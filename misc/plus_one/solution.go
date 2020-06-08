package solution

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			if digits[i] == 9 && i == 0 {
				digits = append([]int{1}, digits...)
			}
			continue
		}

		digits[i]++
		break
	}

	return digits
}
