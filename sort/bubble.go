package sort

import (
	"fmt"
	"sort"
	"time"
)

// Bubble is simple bubble sort
func Bubble(input []int) []int {
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

// BubbleOpt optimized sort
func BubbleOpt(input []int) []int {
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

// Std is standard library sort
func Std(input []int) []int {
	arr := make([]int, len(input))
	copy(arr, input)
	start := time.Now()
	sort.Ints(arr)
	bench := time.Since(start)
	fmt.Println(bench)
	return arr
}

// BubbleString is bubble string sort
func BubbleString(input []string) []string {
	arr := make([]string, len(input))
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

// BubbleReverse rbubble everse sort
func BubbleReverse(input []int) []int {
	arr := make([]int, len(input))
	copy(arr, input)
	var isSorted bool
	tStart := time.Now()
	for i := 0; i < len(arr) || !isSorted; i++ {
		isSorted = true
		for j, k := 0, 1; k < len(arr)-i; {
			if arr[k] > arr[j] {
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
