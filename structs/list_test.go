package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAppend(t *testing.T) {
	l := &List{
		&Node{
			Next: &Node{
				Next:  nil,
				Value: 2,
			},
			Value: 1,
		},
	}

	want := &List{
		&Node{
			Next: &Node{
				Next: &Node{
					Next:  nil,
					Value: 3,
				},
				Value: 2,
			},
			Value: 1,
		},
	}

	Node3 := &Node{
		Next:  nil,
		Value: 3,
	}

	l.Append(Node3)

	assert.Equal(t, want, l)
}
