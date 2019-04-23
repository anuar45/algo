package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Numbers!")

	a := 57
	b := 548579342

	fmt.Println(a, b)
	//fmt.Println(getSubInt(6, 3, b))
	//fmt.Println(getIntLen(b))
	fmt.Println(getSubNumIndex(a, b))
}

func getSubNumIndex(a, b int) int {
	aLen := getIntLen(a)
	bLen := getIntLen(b)
	for i := 0; i < bLen; i++ {
		bSub := getSubInt(i, aLen, b)
		if a == bSub {
			return i
		}
	}
	return 9999
}

func getIntLen(num int) int {
	var i int
	for i = 0; num != 0; i++ {
		num /= 10
	}
	return i
}

func getSubInt(index, length, num int) int {
	var c, d, result int
	numLen := getIntLen(num)
	index = numLen - index - length
	t := 1
	for i := 0; i <= index; i++ {
		c = num / t
		d = (num / (t * (int(math.Pow10(length))))) * (int(math.Pow10(length)))
		//fmt.Println(c, d)
		result = c - d
		t *= 10
	}
	return result
}
