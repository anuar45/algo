package main

import "fmt"

func main() {
	fmt.Println("Hello Graph!")

	vertex1 := &Vertex{Key:1}
	vertex2 := &Vertex{Key:2}
	vertex3 := &Vertex{Key:3}
	vertex4 := &Vertex{Key:4}
	vertex5 := &Vertex{Key:5}
	vertex6 := &Vertex{Key:6}
	vertex7 := &Vertex{Key:7}

	vertex1.AddEdge(vertex2, vertex3)
	vertex2.AddEdge(vertex4, vertex5)
	vertex3.AddEdge(vertex6, vertex7)
	
	vertex1.BFS()
	vertex1.DFS()

	// var s Stack
	// s.Push(&Vertex{Key:1})
	// s.Push(&Vertex{Key:2})
	// s.Push(&Vertex{Key:3})
	// s.Push(&Vertex{Key:4})
	// s.Push(&Vertex{Key:5})
	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())
}

type Vertex struct {
	Key int
	Edges []*Vertex
}

func (start *Vertex) AddEdge(v ...*Vertex) {
	start.Edges = append(start.Edges, v...)
}

func (start *Vertex) BFS() {
	var q Queue
	visited := make(map[*Vertex]bool)

	q.Enqueue(start)

	for q.Len() > 0 {
		curVertex := q.Dequeue()
		//visited[curVertex] = true
		fmt.Println(curVertex.Key)

		for _, v := range curVertex.Edges {
			if !visited[v] {
				q.Enqueue(v)
				visited[v] = true
			}
		}
	} 
}

func (start *Vertex) DFS() {
	var s Stack
	visited := make(map[*Vertex]bool)

	s.Push(start)
	//visited[start] = true

	for s.Len() > 0 {
		curVertex := s.Pop()
		fmt.Println(curVertex.Key)

		for _, v := range curVertex.Edges {
			if !visited[v] {
				s.Push(v)
				visited[v] = true
			}
		}
	}


	// for _ , v := range curVertex.Edges {
	// 	visited[v] = true


	// }

	// for curVertex := s.Pop(); curVertex != nil; {
	// 	visited[curVertex] = true
	// 	fmt.Println(curVertex.Key)

	// 	for _, v := range curVertex.Edges {
	// 		if !visited[v] {
	// 			s.Push(v)
	// 			visited[v] = true
	// 		}
	// 	}

	// }
}


type Queue struct {
	vertices []*Vertex
}

func (q *Queue) Enqueue(v *Vertex) {
	q.vertices = append(q.vertices, v)
}

func (q *Queue) Dequeue() *Vertex {
	if len(q.vertices) == 0 {
		return nil
	}

	v := q.vertices[0]
	q.vertices = q.vertices[1:]

	return v
}

func (q *Queue) Len() int {
	return len(q.vertices)
}

type Stack struct {
	vertices []*Vertex
}

func (s *Stack) Push(v *Vertex) {
	s.vertices = append(s.vertices, v)
}

func (s *Stack) Pop() *Vertex {
	if len(s.vertices) == 0 {
		return nil
	}

	v := s.vertices[len(s.vertices)-1]
	s.vertices = s.vertices[:len(s.vertices)-1]

	return v
}

func (s *Stack) Len() int {
	return len(s.vertices)
}