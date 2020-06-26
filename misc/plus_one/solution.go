package solution

func plusOne(digits []int) []int {
	i := len(digits) - 1
	r := 1

	if digits[i] == 9 {
		for r != 0 {
			switch {
			case i == 0 && digits[i] != 9:
				digits[i]++
				r = 0
			case i == 0:
				digits[i] = 0
				digits = append([]int{1}, digits...)
				r = 0
			case digits[i] == 9:
				digits[i] = 0
				r = 1
			case digits[i] != 9:
				digits[i]++
				r = 0  
			}
			i--
		}
	} else {
		digits[i]++
	}

	return digits
}
