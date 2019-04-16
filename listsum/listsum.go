package main

import "fmt"

func main() {
	var l1n1 = ListNode{Val: 9, Next: nil}
	var l1n2 = ListNode{Val: 9, Next: nil}
	var l1n3 = ListNode{Val: 9, Next: nil}
	var l2n1 = ListNode{Val: 9, Next: nil}
	var l2n2 = ListNode{Val: 9, Next: nil}
	var l2n3 = ListNode{Val: 9, Next: nil}

	l1 := l1n1
	l1.Next = &l1n2
	l1.Next.Next = &l1n3
	l2 := l2n1
	l2.Next = &l2n2
	l2.Next.Next = &l2n3
	result := addTwoNumbers(&l1, &l2)

	printList(&l1)
	printList(&l2)
	printList(result)
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
		if p != nil && q != nil {
			var t ListNode
			l.Next = &t
			l = l.Next
		}
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	return r
}
