package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("boo")
}

func findMedianSortedArraysDumb(nums1, nums2 []int) float64 {
	var result float64

	sl := append(nums1, nums2...)
	sort.Ints(sl)

	if len(sl)%2 == 1 {
		result = float64(sl[len(sl)/2])
	} else {
		result = float64(sl[len(sl)/2]+sl[len(sl)/2-1]) / 2
	}
	return result
}

func findMedianSortedArraysRecursive(nums1, nums2 []int) float64 {
	var result float64

	// sl := append(nums1, nums2...)
	// sort.Ints(sl)

	// if len(sl)%2 == 1 {
	// 	result = float64(sl[len(sl)/2])
	// } else {
	// 	result = float64(sl[len(sl)/2] + sl[len(sl)/2-1])
	// }
	return result
}
func findMedianSortedArraysQuick(nums1, nums2 []int) float64 {
	var result float64

	// sl := append(nums1, nums2...)
	// sort.Ints(sl)

	// if len(sl)%2 == 1 {
	// 	result = float64(sl[len(sl)/2])
	// } else {
	// 	result = float64(sl[len(sl)/2] + sl[len(sl)/2-1])
	// }
	return result
}
