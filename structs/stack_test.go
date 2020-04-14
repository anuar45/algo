package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := NewStack()
	var v1 interface{}

	v1 = 777

	s.Push(v1)
	v2, _ := s.Pop()

	assert.Equal(t, v1, v2)
}
