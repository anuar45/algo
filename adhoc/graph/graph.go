package graph

import (
	"container/list"
	// "fmt"
)

type Node struct {
	Value   int
	Friends []*Node
}

func bfs(start *Node) []*Node {
	visited := make(map[*Node]bool)
	var visitOrder []*Node

	q := list.New()

	q.PushBack(start)

	visited[start] = true
	visitOrder = append(visitOrder, start)

	for q.Len() > 0 {
		element := q.Front()

		for _, n := range element.Value.(*Node).Friends {
			if !visited[n] {
				visited[n] = true
				q.PushBack(n)
				visitOrder = append(visitOrder, n)
			}
		}
		q.Remove(element)
	}

	return visitOrder
}

func dfs(start *Node) []*Node {
	visited := make(map[*Node]bool)
	var visitOrder []*Node

	stack := list.New()

	stack.PushFront(start)
	visited[start] = true
	//	visitOrder = append(visitOrder, start)

	for stack.Len() > 0 {
		elem := stack.Front()
		node := elem.Value.(*Node)
		stack.Remove(elem)
		visitOrder = append(visitOrder, node)

		for _, n := range node.Friends {
			if !visited[n] {
				visited[n] = true
				stack.PushFront(n)
			}
		}
	}

	return visitOrder
}
