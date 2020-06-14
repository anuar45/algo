package solution

import (
	"container/list"
)

// CheckBalanced checks for balanced brackets in the string
func CheckBalanced(s string) bool {
	stack := list.New()

	for _, r := range s {
		if r == '(' {
			stack.PushFront(r)
		}

		if r == ')' {
			elem := stack.Front()
			if elem != nil && elem.Value.(rune) == '(' {
				stack.Remove(elem)
			} else {
				return false
			}
		}
	}

	if stack.Len() == 0 {
		return true
	}

	return false
}
