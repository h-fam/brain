package dag

import (
	"fmt"
	"hfam/brain/base/go/key"
)

// Graph is a generic graph handler for loading an arbitrary nodes and edges.
// This can then be turned into a DAG for futher processing.
// a node must implment the key.Key interface for uniqueness.
type Graph struct {
	nodes map[key.Key]key.Key
	edges map[key.Key]EdgeSet
}

type EdgeSet struct{}

func New() *Graph {
	return &Graph{
		nodes: map[key.Key]key.Key{},
	}
}

// AddNode will add a node to the graph.  Duplicate nodes will raise an error.
func (g *Graph) AddNode(n key.Key) error {
	if _, ok := g.nodes[n]; ok {
		return key.DuplicateError(n.String())
	}
	g.nodes[n] = n
	return nil
}

func (g *Graph) AddEdge(src, dst key.Key) error {
	if _, ok := g.nodes[src]; !ok {
		return fmt.Errorf("Node %q not found", src)
	}
	if _, ok := g.nodes[dst]; !ok {
		return fmt.Errorf("Node %q not found", dst)
	}
	return nil
}
