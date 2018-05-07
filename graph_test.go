package goraph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGraph(t *testing.T) {
	assert := assert.New(t)
	graph := NewGraph(GraphTypeDirect)

	// new
	graph.AddEdge(0, 1, float64(2))
	graph.AddEdge(0, 2, float64(4))
	graph.AddEdge(1, 2, float64(1))

	graph.AddEdge(2, 3, float64(7))
	graph.AddEdge(2, 0, float64(4))

	graph.AddEdge(3, 2, float64(5))
	graph.AddEdge(3, 4, float64(2))
	graph.AddEdge(3, 5, float64(6))

	graph.AddEdge(4, 5, float64(3))

	// new
	graph.AddEdge(5, 3, float64(6))

	assert.Equal(float64(2), graph.nodes[0].adjNodes[1].weight)

	distance, preNode := ShortPathDijkstra(graph, 0)
	t.Log("preNode: ", preNode)
	t.Log("distance: ", distance)

	allDistance := ShortestPathsFloyd(graph)
	PrintMatrix(allDistance)
}
