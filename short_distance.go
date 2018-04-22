package goraph

import (
	"fmt"
	"math"
)

func ShortPathDijkstra(graph *Graph, sourceNodeID int) (map[int]float64, map[int]int) {
	var set = make(map[int]bool)
	var distance = make(map[int]float64)
	var preNode = make(map[int]int)

	sourceNode := graph.nodes[sourceNodeID]

	for _, node := range graph.nodes {
		if node.nodeID != sourceNodeID {
			set[node.nodeID] = true
		}

		if node.nodeID == sourceNodeID {
			distance[node.nodeID] = float64(0)
		} else if edge, ok := sourceNode.adjNodes[node.nodeID]; ok {
			distance[node.nodeID] = edge.weight
		} else {
			distance[node.nodeID] = math.Inf(1)
		}
	}

	count := len(set)

	for i := 0; i < count; i++ {
		shortID := getMinDistanceID(set, distance)
		delete(set, shortID)

		if len(preNode) == 0 {
			preNode[shortID] = sourceNodeID
		}

		for _, edge := range graph.nodes[shortID].adjNodes {
			oldDis := distance[edge.toNodeID]
			newDis := distance[shortID] + edge.weight
			if newDis < oldDis {
				distance[edge.toNodeID] = newDis
				preNode[edge.toNodeID] = shortID
			}
		}
		fmt.Println("short ID: ", shortID)
		fmt.Println("distance: ", distance)
	}
	return distance, preNode
}

func getMinDistanceID(set map[int]bool, distance map[int]float64) int {
	shortID := 0
	shortDis := math.Inf(1)
	for id, _ := range set {
		if shortDis > distance[id] {
			shortDis = distance[id]
			shortID = id
		}
	}
	return shortID
}
