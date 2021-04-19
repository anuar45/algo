package main

import (
	"fmt"
	"time"
	"math/rand"
)



func main() {

	out := fanIn(routine("Joe"), routine("Ann"))

	for {
		fmt.Println(<-out)
	} 
	
}

func routine(who string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			n := rand.Intn(5)
			c <- fmt.Sprintf("%s %d", who, n)
			time.Sleep(time.Duration(n)*time.Second)
		}
	}()

	return c
}

func fanIn(in1, in2 <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for {
			select {
			case s := <- in1:
				out <- s
			case s := <- in2:
				out <- s
			}
		}
	}()

	return out
}