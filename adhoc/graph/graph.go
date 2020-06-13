package graph

import (
	"container/list"
	"fmt"
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

//////////////////

type Graph struct {
	AdjList map[int][]int
}

func (g *Graph) BFSIter(start int) []int {
	var order []int
	visited := make(map[int]bool)

	q := list.New()

	q.PushBack(start)

	for q.Len() > 0 {
		elem := q.Front()
		q.Remove(elem)
		vertex := elem.Value.(int)
		order = append(order, vertex)

		for _, v := range g.AdjList[vertex] {
			if !visited[v] {
				q.PushBack(v)
			}
		}
	}

	return order
}

func (g *Graph) DFSIter(start int) []int {
	var order []int
	visited := make(map[int]bool)

	stack := list.New()

	stack.PushFront(start)

	for stack.Len() > 0 {
		elem := stack.Front()
		stack.Remove(elem)
		vertex := elem.Value.(int)
		order = append(order, vertex)

		for _, v := range g.AdjList[vertex] {
			if !visited[v] {
				stack.PushFront(v)
			}
		}
	}

	return order
}

func (g *Graph) Dijkstra(start, target int) int {
	visited := make(map[int]bool)
	distance := make(map[int]int)
	w := 1
	q := list.New()

	q.PushBack(start)
	distance[start] = 0

	for q.Len() > 0 {
		elem := q.Front() // should get closest vertex
		q.Remove(elem)

		vertex := elem.Value.(int)

		visited[vertex] = true

		fmt.Println(distance)

		for _, v := range g.AdjList[vertex] {
			if visited[v] {
				continue
			}

			q.PushBack(v)

			if _, ok := distance[v]; !ok || distance[vertex]+w < distance[v] {
				distance[v] = distance[vertex] + w
			}
		}
	}

	return distance[target]

}
