package graph

import (
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

	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)
	g.AddNode(4)
	g.AddNode(5)
	g.AddNode(6)
	g.AddNode(7)

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)

	assert.Equal(t, graphResult, g.String())
}
