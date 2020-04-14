package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	var v1 interface{}

	v1 = 777

	q.Enqueue(777)

	v2, _ := q.Dequeue()

	assert.Equal(t, v1, v2)
}
