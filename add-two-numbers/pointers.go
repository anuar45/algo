package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

type ListNodePtr *struct {
	Value int
	Next  ListNodePtr
}

func main() {
	n1 := ListNode{Value: 3, Next: nil}
	n2 := ListNode{Value: 5, Next: nil}
	n3 := ListNode{Value: 7, Next: nil}

	l1 := n1
	l1.Next = &n2
	l1.Next.Next = &n3

	var l2 ListNodePtr

	var i int
	i = l1.Value

	fmt.Println(i)
	fmt.Println(l2)
	fmt.Println(&l1.Value)
	fmt.Println(l1.Next)

	fmt.Println("--------")
	printList(&l1)
}

func printList(l *ListNode) {
	p := l
	fmt.Println(p, l)
	fmt.Println(p.Value)
	fmt.Println(l.Value)

}
