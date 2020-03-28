package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var nums []int
	for _, v := range os.Args[1:] {
		arg, _ := strconv.Atoi(v)
		nums = append(nums, arg)
	}
	start := time.Now()
	r := twoSum(nums[1:], nums[0])
	bench := time.Since(start).Nanoseconds()
	fmt.Println(r, bench)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for k, v := range nums {
		if v, ok := m[v]; ok {
			return []int{k, v}
		}
		m[target-v] = k
	}
	return nil
}
