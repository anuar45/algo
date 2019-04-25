package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	input := []int{4, 9, 6, 1, 3, 7, 5, 8}

	fmt.Println(sortBubble(input))
	fmt.Println(sortBubbleOpt(input))
	fmt.Println(sortStd(input))

}

func sortBubble(input []int) []int {
	arr := make([]int, len(input))
	copy(arr, input)
	tStart := time.Now()
	for i := 0; i < len(arr); i++ {
		for j, k := 0, 1; k < len(arr); {
			if arr[k] < arr[j] {
				arr[j], arr[k] = arr[k], arr[j]
			}
			j++
			k++
		}
	}
	bench := time.Since(tStart)
	fmt.Println(bench)
	return arr
}

func sortBubbleOpt(input []int) []int {
	arr := make([]int, len(input))
	copy(arr, input)
	var isSorted bool
	tStart := time.Now()
	for i := 0; i < len(arr) || !isSorted; i++ {
		isSorted = true
		for j, k := 0, 1; k < len(arr)-i; {
			if arr[k] < arr[j] {
				arr[j], arr[k] = arr[k], arr[j]
				isSorted = false
			}
			j++
			k++
		}
	}
	bench := time.Since(tStart)
	fmt.Println(bench)
	return arr
}

func sortStd(input []int) []int {
	arr := make([]int, len(input))
	copy(arr, input)
	start := time.Now()
	sort.Ints(arr)
	bench := time.Since(start)
	fmt.Println(bench)
	return arr
}
