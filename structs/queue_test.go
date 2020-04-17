package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	var v1 interface{}

	v1 = 1

	q.Enqueue(777)
	q.Enqueue(888)
	q.Enqueue(999)
	q.Enqueue(v1)
	//fmt.Println(q)
	_, _ = q.Dequeue()
	_, _ = q.Dequeue()
	_, _ = q.Dequeue()
	v2, _ := q.Dequeue()

	assert.Equal(t, v1, v2)
	assert.Equal(t, 0, q.Len())

}
