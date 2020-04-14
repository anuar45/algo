package structs

type Queue struct {
	l *List
}

func NewQueue() *Queue {
	l := NewList()
	return &Queue{l}
}

// Enqueue put item into queue
func (q *Queue) Enqueue(v interface{}) error {
	err := q.l.Insert(0, v)
	return err
}

// Dequeue retrieves item from queue
func (q *Queue) Dequeue() (interface{}, error) {
	v, err := q.l.Last()
	if err != nil {
		return v, err
	}

	err = q.l.RemoveLast()
	return v, err
}
