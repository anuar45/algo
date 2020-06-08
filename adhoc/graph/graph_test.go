package graph

import (
	"fmt"
	"testing"
)

var (
	node1 = &Node{Value: 1}
	node2 = &Node{Value: 2}
	node3 = &Node{Value: 3}
	node4 = &Node{Value: 4}
	node5 = &Node{Value: 5}
	node6 = &Node{Value: 6}
	node7 = &Node{Value: 7}
)

func TestBFS(t *testing.T) {

	node1.Friends = []*Node{node2, node3}
	node2.Friends = []*Node{node4, node5}
	node3.Friends = []*Node{node6, node7}

	gotBfs := bfs(node1)

	for _, node := range gotBfs {
		fmt.Print(node.Value)
	}

}

func TestDFS(t *testing.T) {
	node1.Friends = []*Node{node2, node3}
	node2.Friends = []*Node{node4, node5}
	node3.Friends = []*Node{node6, node7}

	gotDfs := dfs(node1)

	for _, node := range gotDfs {
		fmt.Print(node.Value)
	}
}
