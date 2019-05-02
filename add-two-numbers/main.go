package main

import "fmt"

func main() {
	var l1n1 = ListNode{Val: 2, Next: nil}
	var l1n2 = ListNode{Val: 3, Next: nil}
	var l1n3 = ListNode{Val: 4, Next: nil}
	var l2n1 = ListNode{Val: 2, Next: nil}
	var l2n2 = ListNode{Val: 3, Next: nil}
	var l2n3 = ListNode{Val: 4, Next: nil}
	var l3n1 = ListNode{Val: 5, Next: nil}
	var l4n1 = ListNode{Val: 5, Next: nil}

	l1 := l1n1
	l1.Next = &l1n2
	l1.Next.Next = &l1n3
	l2 := l2n1
	l2.Next = &l2n2
	l2.Next.Next = &l2n3
	l3 := l3n1
	l4 := l4n1
	result1 := addTwoNumbers(&l1, &l2)
	result2 := addTwoNumbers(&l3, &l4)

	printList(&l1)
	printList(&l2)
	printList(result1)
	printList(&l3)
	printList(&l4)
	printList(result2)
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val)
		l = l.Next
	}
	fmt.Println()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l := new(ListNode)
	var carry, x, y int
	p, q, r := l1, l2, l
	for p != nil || q != nil || carry != 0 {
		if p == nil {
			x = 0
		} else {
			x = p.Val
		}
		if q == nil {
			y = 0
		} else {
			y = q.Val
		}

		sum := x + y + carry
		carry = sum / 10

		if sum >= 10 {
			l.Val = sum - 10
		} else {
			l.Val = sum
		}

		l.Next = new(ListNode)
		l = l.Next

		// Switch to next nodes
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	return r
}
