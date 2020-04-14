package structs

type Stack struct {
	l *List
}

func NewStack() *Stack {
	l := NewList()
	return &Stack{l}
}

func (s *Stack) Push(v interface{}) error {
	err := s.l.Append(v)
	return err
}

func (s *Stack) Pop() (interface{}, error) {
	v, err := s.l.Last()
	if err != nil {
		return nil, err
	}

	err = s.l.RemoveLast()

	return v, err
}
