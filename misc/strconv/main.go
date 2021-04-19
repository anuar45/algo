package main

import (
	"fmt"
	"strconv"
)


func main() {
	n, _ := strconv.Atoi("5")
	fmt.Printf("Atoi\t%T %d\n", n, n)
	fmt.Printf("%d\n", &n)

	fmt.Printf("%s", strconv.Itoa(5))

}