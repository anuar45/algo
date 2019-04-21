package main

import "fmt"

func main() {
	input := []int{4, 9, 6, 1, 3}

	fmt.Println(sortBubble(input))
}

func sortBubble(input []int) []int {
	arr := make([]int, len(input))
	copy(arr, input)
	for i := 0; i < len(arr); i++ {
		for j, k := 0, 1; k < len(arr); {
			if arr[k] < arr[j] {
				arr[j], arr[k] = arr[k], arr[j]
			}
			j++
			k++
		}
	}
	return arr
}
