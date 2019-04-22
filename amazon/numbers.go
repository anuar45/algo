package main

import "fmt"

func main() {
	fmt.Println("Numbers!")

	//a := 57
	b := 548579342

	fmt.Println(getSubInt(2, b))
	fmt.Println(getIntLen(b))
}

func getSubNumIndex(a, b int) int {
	aLen := getIntLen(a)
	bLen := getIntLen(b)
	var aIdx, bIdx int
	for i := 0; i < aLen; i++ {
		aIdx = getSubInt(i, a)
		for j := 0; j < bLen; j++ {
			bIdx = getSubInt(j, b)
			if aIdx == bIdx {

				break
			}
		}
	}
}

func getIntLen(num int) int {
	var i int
	for i = 0; num != 0; i++ {
		num /= 10
	}
	return i
}

func getSubInt(index, num int) int {
	var c, d, result int
	numLen := getIntLen(num)
	index = numLen - index - 1
	t := 1
	for i := 0; i <= index; i++ {
		c = num / t
		d = (num / (t * 10)) * 10
		fmt.Println(c, d)
		result = c - d
		t *= 10
	}
	return result
}
