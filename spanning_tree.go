package goraph

import ()

func SpanningTree(graph *Graph) (*Graph, float64, error) {
	var rGraph = NewGraph(GraphTypeUndirect)
	var edges = orderedEdges(graph)
	var totalWeight = float64(0)

	for !edges.IsEmpty() {
		edge := edges.DeQueue()
		if checkCycle(rGraph, edge.fromNodeID, edge.toNodeID) {
			continue
		}
		rGraph.AddEdge(edge.fromNodeID, edge.toNodeID, edge.weight)
		totalWeight += edge.weight
	}
	return rGraph, totalWeight, nil
}

func orderedEdges(graph *Graph) HeapEdgeArr {
	arr := NewHeapEdgeArr()
	for _, fromNode := range graph.nodes {
		for _, toNode := range fromNode.adjNodes {
			if graph.graphType == GraphTypeUndirect {
				if fromNode.nodeID > toNode.toNodeID {
					arr.Queue(NewHeapEdge(fromNode.nodeID, toNode.toNodeID, toNode.weight))
				}
			}
		}
	}
	return arr
}

// checkCycle returns bools if new edge from fromNodeID to toNodeID added, there is cycle in graph.
func checkCycle(graph *Graph, fromNodeID int, toNodeID int) bool {
	stack := NewNodeStack()
	visited := make(map[int]bool)

	for nodeID, _ := range graph.nodes {
		visited[nodeID] = false
	}

	visited[fromNodeID] = true
	stack.Push(fromNodeID)

	for !stack.IsEmpty() {
		thisNodeID := stack.Pop()
		if thisNodeID == toNodeID {
			return true
		}

		if _, ok := graph.nodes[thisNodeID]; !ok {
			continue
		}

		for adjNodeID, _ := range graph.nodes[thisNodeID].adjNodes {
			if !visited[adjNodeID] {
				visited[adjNodeID] = true
				stack.Push(adjNodeID)
			}
		}
	}
	return false
}
