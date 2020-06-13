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

func TestBFSIter(t *testing.T) {
	graph := Graph{make(map[int][]int)}

	graph.AdjList[1] = []int{2, 3}
	graph.AdjList[2] = []int{4, 5}
	graph.AdjList[3] = []int{6, 7}

	fmt.Println(graph.BFSIter(1))

}

func TestDFSIter(t *testing.T) {
	graph := Graph{make(map[int][]int)}

	graph.AdjList[1] = []int{2, 3}
	graph.AdjList[2] = []int{4, 5}
	graph.AdjList[3] = []int{6, 7}

	fmt.Println(graph.DFSIter(1))

}

func TestDijkstra(t *testing.T) {
	graph := Graph{make(map[int][]int)}

	graph.AdjList[1] = []int{2, 3}
	graph.AdjList[2] = []int{4, 5}
	graph.AdjList[3] = []int{6, 7}

	fmt.Println(graph.Dijkstra(1, 7))

}
