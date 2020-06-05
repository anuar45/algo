package solution

func plusOne(digits []int) []int {
	r := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 && r == 1 {
			digits[i] = 0
			r = 1
			continue
		}

		if r == 1 {
			digits[i]++
			break
		}
	}

	return digits
}
