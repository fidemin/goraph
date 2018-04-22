package goraph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGraph(t *testing.T) {
	assert := assert.New(t)
	graph := NewGraph()

	node0 := NewNode(0)
	node0.AddEdge(NewEdge(1, float64(2)))
	node0.AddEdge(NewEdge(2, float64(4)))
	graph.AddNode(node0)

	node1 := NewNode(1)
	node1.AddEdge(NewEdge(2, float64(1)))
	graph.AddNode(node1)

	node2 := NewNode(2)
	node2.AddEdge(NewEdge(3, float64(7)))
	node2.AddEdge(NewEdge(0, float64(4)))
	graph.AddNode(node2)

	node3 := NewNode(3)
	node3.AddEdge(NewEdge(4, float64(2)))
	node3.AddEdge(NewEdge(5, float64(6)))
	graph.AddNode(node3)

	node4 := NewNode(4)
	node4.AddEdge(NewEdge(5, float64(3)))
	graph.AddNode(node4)

	node5 := NewNode(5)
	node5.AddEdge(NewEdge(3, float64(6)))
	graph.AddNode(node5)
	assert.Equal(float64(2), graph.nodes[0].adjNodes[1].weight)

	distance, preNode := ShortPathDijkstra(graph, 0)
	t.Log("preNode: ", preNode)
	t.Log("distance: ", distance)
}
