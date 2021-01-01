package dag

import (
	"fmt"

	"github.com/h-fam/brain/base/go/errs"
	"github.com/h-fam/brain/base/go/key"
	"github.com/h-fam/brain/base/go/log"
)

// Node is generic node to be used in the graph.
type Node interface{}

// Graph is a generic graph handler for loading an arbitrary nodes and edges.
// This can then be turned into a DAG for futher processing.
// a node must implment the key.Key interface for uniqueness.
type Graph struct {
	nodes map[key.Key]Node
	edges map[Edge]struct{}
}

// Edge is a keyable structure for maintaining graph edges
type Edge struct {
	a key.Key
	z key.Key
}

// Equal returns the equality of e and dst.
func (e *Edge) Equal(dst interface{}) bool {
	de, ok := dst.(*Edge)
	if !ok {
		return false
	}
	if e.a != de.a {
		return false
	}
	if e.z != de.z {
		return false
	}
	return true
}

// String returns the string representation of e.
func (e *Edge) String() string {
	return fmt.Sprintf("{a: %v, z: %v}", e.a, e.z)
}

// New returns a new Graph.
func New() *Graph {
	return &Graph{
		nodes: map[key.Key]Node{},
		edges: map[Edge]struct{}{},
	}
}

// AddNode will add a node to the graph.  Duplicate nodes will return an error.
func (g *Graph) AddNode(k key.Key, n Node) error {
	if _, ok := g.nodes[k]; ok {
		return key.DuplicateError(k.String())
	}
	g.nodes[k] = n
	return nil
}

// AddEdge will add a new edge to the graph. Duplicate edges will return an error.
func (g *Graph) AddEdge(src, dst key.Key) error {
	if _, ok := g.nodes[src]; !ok {
		return fmt.Errorf("Node %q not found", src)
	}
	if _, ok := g.nodes[dst]; !ok {
		return fmt.Errorf("Node %q not found", dst)
	}
	e := Edge{a: src, z: dst}
	if _, ok := g.edges[e]; ok {
		return fmt.Errorf("Edge %v exists in graph", e)
	}
	g.edges[e] = struct{}{}
	return nil
}

// TraverseFunc is a walk function for execution against the provided node.
type TraverseFunc func(Node) error

// Traverse starts at node k and will execute fn on all nodes from k and decendents.
func (g *Graph) Traverse(k key.Key, fn TraverseFunc) error {
	s := []key.Key{k}
	visited := map[key.Key]bool{}
	var errList errs.List
	for len(s) != 0 {
		log.Info(s)
		curr := s[0]
		s = s[1:]
		err := fn(g.nodes[curr])
		if err != nil {
			errList.Add(err)
		}
		visited[curr] = true
		for e := range g.edges {
			if e.a == curr {
				if _, ok := visited[e.z]; ok {
					errList.Add(fmt.Errorf("loop detected on node %v", e.z))
					break
				}
				s = append(s, e.z)
			}
		}
	}
	return errList.Err()
}
