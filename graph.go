package goraph

import (
	"errors"
	"fmt"
)

type Edge struct {
	toNodeID int
	weight   float64
}

func NewEdge(nodeID int, weight float64) Edge {
	return Edge{
		toNodeID: nodeID,
		weight:   weight,
	}
}

type Node struct {
	nodeID   int
	adjNodes map[int]Edge
}

func NewNode(nodeID int) *Node {
	g := new(Node)
	g.nodeID = nodeID
	g.adjNodes = make(map[int]Edge)
	return g
}

func (g *Node) AddEdge(edge Edge) {
	g.adjNodes[edge.toNodeID] = edge
}

type Graph struct {
	nodes map[int]*Node
}

func NewGraph() *Graph {
	g := new(Graph)
	g.nodes = make(map[int]*Node)
	return g
}

func (g *Graph) AddNode(node *Node) error {
	if _, ok := g.nodes[node.nodeID]; ok {
		return errors.New(fmt.Sprintf("node with id %d already exists", node.nodeID))
	}

	if node == nil {
		return errors.New("node is nil")
	}

	g.nodes[node.nodeID] = node
	return nil
}
