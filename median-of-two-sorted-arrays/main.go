package main

import "fmt"

func main() {
	a := []int{4, 5, 2, 1}
	b := []int{1, 3, 9}

	fmt.Println(findMedianSortedArrays(a, b))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sum1 := sumSlice(nums1)
	sum2 := sumSlice(nums2)
	sum := float64(sum1+sum2) / float64(len(nums1)+len(nums2))

	fmt.Println(sum1)
	fmt.Println(sum2)
	fmt.Println(len(nums1))
	fmt.Println(len(nums2))

	return sum
}

func sumSlice(s []int) int {
	var sum int
	for _, v := range s {
		sum += v
	}
	return sum
}
