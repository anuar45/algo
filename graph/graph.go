package graph

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/anuar45/algo/structs"
)

type Graph struct {
	nodes []*Node
	edges map[*Node][]*Node
}

type Node struct {
	Value interface{}
}

func NewGraph() *Graph {
	return &Graph{edges: make(map[*Node][]*Node)}
}

func (g *Graph) AddNode(v interface{}) error {
	g.nodes = append(g.nodes, &Node{v})
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

func (g *Graph) find(v interface{}) *Node {
	for _, n := range g.nodes {
		if reflect.DeepEqual(v, n.Value) {
			return n
		}
	}
	return nil
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Value)
}

func (g *Graph) String() string {
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + " -> "
		near := g.edges[g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	return s
}

func (g *Graph) TraverseBFS() {
	q := structs.NewQueue()
	q.Enqueue(g.nodes[0])

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
