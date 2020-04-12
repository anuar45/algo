package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAppend(t *testing.T) {
	l1node1 := &Node{Value: 1}
	l1node2 := &Node{Value: 2}
	l1node1.Next = l1node2
	l1node2.Prev = l1node1
	l1 := &List{l1node1}

	l2node1 := &Node{Value: 1}
	l2node2 := &Node{Value: 2}
	l2node3 := &Node{Value: 3}
	l2node1.Next = l2node2
	l2node2.Prev = l2node1
	l2node2.Next = l2node3
	l2node3.Prev = l2node2
	l2 := &List{l2node1}

	l1.Append(3)

	assert.Equal(t, l1, l2)
}

func TestListRemove(t *testing.T) {
	l1node1 := &Node{Value: 1}
	l1node2 := &Node{Value: 2}
	l1node1.Next = l1node2
	l1node2.Prev = l1node1
	l1 := &List{l1node1}

	l2node1 := &Node{Value: 1}
	l2node2 := &Node{Value: 2}
	l2node3 := &Node{Value: 3}
	l2node1.Next = l2node2
	l2node2.Prev = l2node1
	l2node2.Next = l2node3
	l2node3.Prev = l2node2
	l2 := &List{l2node1}

	l2.Remove(3)

	assert.Equal(t, l1, l2)
}
