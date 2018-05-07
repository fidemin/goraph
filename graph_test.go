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

func TestCheckCycle(t *testing.T) {
	assert := assert.New(t)
	graph := NewGraph(GraphTypeUndirect)

	//graph.AddEdge(0, 1, float64(4))
	//graph.AddEdge(1, 2, float64(2))
	graph.AddEdge(2, 0, float64(3))
	graph.AddEdge(2, 3, float64(1))
	graph.AddEdge(3, 4, float64(1))
	graph.AddEdge(3, 5, float64(5))
	graph.AddEdge(4, 5, float64(6))

	assert.Equal(false, checkCycle(graph, 0, 1))

}

func TestHeapEdgeArr(t *testing.T) {
	arr := NewHeapEdgeArr()
	arr.Queue(NewHeapEdge(0, 1, float64(10)))
	arr.Queue(NewHeapEdge(2, 3, float64(5)))
	t.Log(arr)
}

func TestOrderedEdges(t *testing.T) {
	graph := NewGraph(GraphTypeUndirect)

	graph.AddEdge(0, 1, float64(4))
	graph.AddEdge(1, 2, float64(2))
	graph.AddEdge(2, 0, float64(3))
	graph.AddEdge(2, 3, float64(1))
	graph.AddEdge(3, 4, float64(1))
	graph.AddEdge(3, 5, float64(5))
	graph.AddEdge(4, 5, float64(6))

	arrEdges := orderedEdges(graph)
	t.Log(arrEdges)

	edge := arrEdges.DeQueue()
	t.Log(edge)
	t.Log(arrEdges)
}

func TestSpanningTree(t *testing.T) {
	assert := assert.New(t)

	graph := NewGraph(GraphTypeUndirect)
	graph.AddEdge(0, 1, float64(4))
	graph.AddEdge(1, 2, float64(2))
	graph.AddEdge(2, 0, float64(3))
	graph.AddEdge(2, 3, float64(1))
	graph.AddEdge(3, 4, float64(1))
	graph.AddEdge(3, 5, float64(5))
	graph.AddEdge(4, 5, float64(6))

	rGraph, minWeight, _ := SpanningTree(graph)
	t.Log(rGraph)
	assert.Equal(float64(12), minWeight)
}
