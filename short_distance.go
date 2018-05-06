package goraph

import (
	"fmt"
	"github.com/yhmin84/tbprint"
	"math"
	"sort"
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

func ShortestPathsFloyd(graph *Graph) map[int]map[int]float64 {
	var result = make(map[int]map[int]float64)
	// initialize
	//
	// sNode represents start node and eNode represents end node.
	for _, sNode := range graph.nodes {
		result[sNode.nodeID] = make(map[int]float64)
		for _, eNode := range graph.nodes {
			if sNode.nodeID == eNode.nodeID {
				result[sNode.nodeID][eNode.nodeID] = float64(0)
			} else {
				if adj, ok := sNode.adjNodes[eNode.nodeID]; !ok {
					result[sNode.nodeID][eNode.nodeID] = math.Inf(1)
				} else {
					result[sNode.nodeID][eNode.nodeID] = adj.weight
				}
			}
		}
	}

	// mNode represents a middle node.
	for _, mNode := range graph.nodes {
		for _, sNode := range graph.nodes {
			for _, eNode := range graph.nodes {
				if result[sNode.nodeID][mNode.nodeID]+result[mNode.nodeID][eNode.nodeID] < result[sNode.nodeID][eNode.nodeID] {
					result[sNode.nodeID][eNode.nodeID] = result[sNode.nodeID][mNode.nodeID] + result[mNode.nodeID][eNode.nodeID]
				}
			}
		}
	}

	return result
}

func PrintMatrix(ppDistance map[int]map[int]float64) {
	ints := make([]int, 0)
	for nodeID, _ := range ppDistance {
		ints = append(ints, nodeID)
	}
	sortedInts := sort.IntSlice(ints)
	sortedInts.Sort()

	ppDistanceStr := make([][]string, 0)
	firstRow := []string{"start/end"}
	for _, nodeID := range sortedInts {
		firstRow = append(firstRow, fmt.Sprintf("Node %d", nodeID))
	}

	ppDistanceStr = append(ppDistanceStr, firstRow)

	for _, sNodeID := range sortedInts {
		rowStr := []string{fmt.Sprintf("Node %d", sNodeID)}
		for _, eNodeID := range sortedInts {
			rowStr = append(rowStr, fmt.Sprintf("%f", ppDistance[sNodeID][eNodeID]))
		}
		ppDistanceStr = append(ppDistanceStr, rowStr)
	}
	tbprint.Print(ppDistanceStr)
}
