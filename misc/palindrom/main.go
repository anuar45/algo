package main

import "fmt"

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome("0P"))
	fmt.Println(isPalindrome("......a....."))
}

func isPalindrome(s string) bool {
	srunes := []rune(s)

	for i := 0; i < len(srunes); {
		if srunes[i] >= 'A' && srunes[i] <= 'Z' {
			srunes[i] += 'a' - 'A'
			continue
		}
		i++
	}

	for i := 0; i < len(srunes); {
		if !(srunes[i] >= 'a' && srunes[i] <= 'z') && !(srunes[i] >= '0' && srunes[i] <= '9') {
			srunes = append(srunes[:i], srunes[i+1:]...)
			continue
		}
		i++
	}

	if len(srunes) == 1 {
		return true
	}

	for i, j := 0, len(srunes)-1; i < j; {
		if srunes[i] != srunes[j] {
			return false
		}
		i++
		j--
	}

	return true
}
