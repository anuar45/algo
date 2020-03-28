package main

import (
	"fmt"

	"github.com/anuar45/algo/sort"
)

func main() {
	// Sorting
	input := []int{4, 9, 6, 1, 3, 7, 5, 8}
	animals := []string{"tiger", "lion", "cat", "cheetah"}
	fmt.Println(sort.BubbleSimple(input))
	fmt.Println(sort.BubbleOpt(input))
	fmt.Println(sort.Std(input))
	fmt.Println(sort.BubbleString(animals))
	fmt.Println(sort.BubbleReverse(input))
}
