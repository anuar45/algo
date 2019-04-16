package main

import "fmt"

func main() {
	var l1n1 = ListNode{Val: 5, Next: nil}
	var l1n2 = ListNode{Val: 7, Next: nil}
	var l1n3 = ListNode{Val: 9, Next: nil}
	var l2n1 = ListNode{Val: 3, Next: nil}
	var l2n2 = ListNode{Val: 5, Next: nil}
	var l2n3 = ListNode{Val: 7, Next: nil}

	l1 := l1n1
	l1.Next = &l1n2
	l1.Next.Next = &l1n3
	l2 := l2n1
	l2.Next = &l2n2
	l2.Next.Next = &l2n3
	result := addTwoNumbers(&l1, &l2)

	fmt.Println(
		l1.Val,
		l1.Next.Val,
		l1.Next.Next.Val,
	)

	fmt.Println(
		l2.Val,
		l2.Next.Val,
		l2.Next.Next.Val,
	)

	fmt.Println(
		result.Val,
		result.Next.Val,
		//result.Next.Next.Val,
	)
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l *ListNode
	var carry, x, y int
	p, q, r := l1, l2, l
	for p != nil || q != nil {
		if p.Next == nil {
			x = 0
		} else {
			x = p.Val
		}
		if q.Next == nil {
			y = 0
		} else {
			y = q.Val
		}

		sum := x + y + carry
		carry = sum / 10

		var t ListNode
		if sum >= 10 {
			t.Val = sum - 10
		} else {
			t.Val = sum
		}

		l = &t

		l = l.Next
		p = p.Next
		q = q.Next
	}
	return l
}
