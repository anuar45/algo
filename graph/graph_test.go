package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var graphResult = `1 -> 2 3 
2 -> 1 4 5 
3 -> 1 6 7 
4 -> 2 
5 -> 2 
6 -> 3 
7 -> 3 
`

func TestGraph(t *testing.T) {
	g := NewGraph()

	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)
	g.AddVertex(6)
	g.AddVertex(7)

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)

	assert.Equal(t, graphResult, g.String())
}

func TestGraphBFS(t *testing.T) {
	g := NewGraph()

	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)
	g.AddVertex(6)
	g.AddVertex(7)

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)

	g.TraverseBFS(1, func(vertex *Vertex) {
		fmt.Println(vertex.Value)
	})
}
