package dag

import (
	"fmt"
	"testing"

	"github.com/h-fam/brain/base/go/key"
)

func TestDag(t *testing.T) {
	k1 := key.NewString("foo")
	k2 := key.NewString("bar")
	g := New()
	err := g.AddEdge(k1, k2)
	if err == nil {
		t.Fatalf("AddEdge(%v, %v) failed: cannot add edge without nodes", k1, k2)
	}
	err = g.AddNode(k1, k1)
	if err != nil {
		t.Fatalf("AddNode(%v, %v) failed: %v", k1, k1, err)
	}
	err = g.AddNode(k2, k2)
	if err != nil {
		t.Fatalf("AddNode(%v, %v) failed: %v", k2, k2, err)
	}
	err = g.AddEdge(k1, k2)
	if err != nil {
		t.Fatalf("AddEdge(%v, %v) failed: %v", k1, k2, err)
	}
	err = g.AddEdge(k1, k2)
	if err == nil {
		t.Fatalf("AddEdge(%v, %v) failed: cannot have duplicate edges", k1, k2)
	}
	err = g.AddEdge(k2, k1)
	if err != nil {
		t.Fatalf("AddEdge(%v, %v) failed: %v", k2, k1, err)
	}
}

func TestTraverse(t *testing.T) {
	g := New()
	nodes := []key.Key{
		key.NewString("node0"),
		key.NewString("node1"),
		key.NewString("node2"),
		key.NewString("node3"),
		key.NewString("node4"),
		key.NewString("node5"),
		key.NewString("node6"),
		key.NewString("node7"),
		key.NewString("node8"),
		key.NewString("node9"),
	}
	for _, n := range nodes {
		g.AddNode(n, n)
	}
	g.AddEdge(nodes[0], nodes[1])
	g.AddEdge(nodes[0], nodes[2])
	g.AddEdge(nodes[0], nodes[3])
	g.AddEdge(nodes[1], nodes[4])
	g.AddEdge(nodes[1], nodes[5])
	g.AddEdge(nodes[2], nodes[6])
	g.AddEdge(nodes[2], nodes[6])
	g.AddEdge(nodes[3], nodes[7])
	g.AddEdge(nodes[4], nodes[8])
	g.AddEdge(nodes[6], nodes[9])
	err := g.Traverse(nodes[0], func(n Node) error {
		k, ok := n.(key.String)
		if !ok {
			return fmt.Errorf("invalid node type")
		}
		t.Logf("%v", k)
		return nil
	})
	if err != nil {
		t.Fatalf("Traverse failed: %v", err)
	}
}

func TestTraverseLoop(t *testing.T) {
	g := New()
	nodes := []key.Key{
		key.NewString("node0"),
		key.NewString("node1"),
		key.NewString("node2"),
		key.NewString("node3"),
		key.NewString("node4"),
		key.NewString("node5"),
		key.NewString("node6"),
		key.NewString("node7"),
		key.NewString("node8"),
		key.NewString("node9"),
	}
	for _, n := range nodes {
		g.AddNode(n, n)
	}
	g.AddEdge(nodes[0], nodes[1])
	g.AddEdge(nodes[0], nodes[2])
	g.AddEdge(nodes[0], nodes[3])
	g.AddEdge(nodes[1], nodes[4])
	g.AddEdge(nodes[1], nodes[5])
	g.AddEdge(nodes[2], nodes[6])
	g.AddEdge(nodes[2], nodes[6])
	g.AddEdge(nodes[3], nodes[7])
	g.AddEdge(nodes[4], nodes[8])
	g.AddEdge(nodes[6], nodes[0])
	err := g.Traverse(nodes[0], func(n Node) error {
		k, ok := n.(key.String)
		if !ok {
			return fmt.Errorf("invalid node type")
		}
		t.Logf("%v", k)
		return nil
	})
	if err == nil {
		t.Fatalf("TraverseLoop failed: %v", err)
	}
}
