package graph

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/anuar45/algo/structs"
)

type Graph struct {
	vertexes []*Vertex
	edges    map[*Vertex][]*Vertex
}

type Vertex struct {
	Value interface{}
}

func NewGraph() *Graph {
	return &Graph{edges: make(map[*Vertex][]*Vertex)}
}

func (g *Graph) AddVertex(v interface{}) error {
	g.vertexes = append(g.vertexes, &Vertex{v})
	return nil
}

func (g *Graph) AddEdge(v1, v2 interface{}) error {
	n1 := g.find(v1)
	n2 := g.find(v2)
	if n1 == nil || n2 == nil {
		return errors.New("cant find node with value v1 or v2")
	}

	g.edges[n1] = append(g.edges[n1], n2)
	g.edges[n2] = append(g.edges[n2], n1)
	return nil
}

func (g *Graph) find(v interface{}) *Vertex {
	for _, n := range g.vertexes {
		if reflect.DeepEqual(v, n.Value) {
			return n
		}
	}
	return nil
}

func (n *Vertex) String() string {
	return fmt.Sprintf("%v", n.Value)
}

func (g *Graph) String() string {
	s := ""
	for i := 0; i < len(g.vertexes); i++ {
		s += g.vertexes[i].String() + " -> "
		near := g.edges[g.vertexes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	return s
}

func (g *Graph) TraverseBFS(start interface{}) {
	q := structs.NewQueue()

	startVertex := g.find(start)

	q.Enqueue(startVertex)

	visited := make(map[*Vertex]bool)

	for q.Len() != 0 {
		qVal, _ := q.Dequeue()
		currVertex := qVal.(*Vertex)
		visited[currVertex] = true

		for _, v := range g.edges[currVertex] {
			if !visited[v] {
				q.Enqueue(v)
				visited[v] = true
			}
		}
	}

	fmt.Println(visited)
}

// def bfs(graph, root):
//     visited, queue = [], [root]
//     while queue:
//         vertex = queue.pop(0)
//         for w in graph[vertex]:
//             if w not in visited:
//                 print(w)
//                 visited.append(w)
//                 queue.append(w)

// graph = {0: [1,2], 1: [3,4], 2: [5,6], 3: [7,8], 4: [9,10]}
// bfs(graph, 0)
