package main

import (
	"testing"
)

// TestEmptyGraph creates a simple graph with just a genesis block and checks for
// correctness of ordering.
func TestEmptyGraph(t *testing.T) {
	// Create the empty shape.
	//
	// O
	//
	g := NewGraph()
	ordering := g.GenesisNode().relativeOrdering()
	if len(ordering) != 1 {
		t.Fatal("There should be exactly one node in the relative ordering of a new graph")
	}
	if ordering[0].name != g.GenesisNode().name {
		t.Error("Genesis node is not ordered correctly:", ordering[0].name, g.GenesisNode().name)
	}
}

// TestBarGraph creates a graph with a genesis block and a single child and
// checks for correctness of ordering.
func TestBarGraph(t *testing.T) {
	// Create the bar shape.
	//
	// O
	// |
	// O
	//
	g := NewGraph()
	c := g.CreateNode(g.GenesisNode())
	ordering := c.relativeOrdering()
	if len(ordering) != 2 {
		t.Fatal("There should be 2 nodes in the bar graph")
	}
	if ordering[0].name != g.GenesisNode().name {
		t.Error("Genesis node is not ordered correctly:", ordering[0].name, g.GenesisNode().name)
	}
	if ordering[1].name != c.name {
		t.Error("Genesis node is not ordered correctly:", ordering[1].name, c.name)
	}
}

// TestLongBarGraph creates a graph with a genesis block and two progressive
// children and checks for correctness of ordering.
func TestLongBarGraph(t *testing.T) {
	// Create the long bar shape.
	//
	// O
	// |
	// O
	// |
	// O
	//
	g := NewGraph()
	n1 := g.CreateNode(g.GenesisNode())
	n2 := g.CreateNode(n1)

	// Check the ordering on n1.
	ordering1 := n1.relativeOrdering()
	if len(ordering1) != 2 {
		t.Fatal("There should be 2 nodes in the long bar graph")
	}
	if ordering1[0].name != g.GenesisNode().name {
		t.Error("Genesis node is not ordered correctly:", ordering1[0].name, g.GenesisNode().name)
	}
	if ordering1[1].name != n1.name {
		t.Error("Genesis node is not ordered correctly:", ordering1[1].name, n1.name)
	}

	// Check the ordering on n2.
	ordering2 := n2.relativeOrdering()
	if len(ordering2) != 3 {
		t.Fatal("There should be 3 nodes in the long bar graph")
	}
	if ordering2[0].name != g.GenesisNode().name {
		t.Error("Genesis node is not ordered correctly:", ordering2[0].name, g.GenesisNode().name)
	}
	if ordering2[1].name != n1.name {
		t.Error("Genesis node is not ordered correctly:", ordering2[1].name, n1.name)
	}
	if ordering2[2].name != n2.name {
		t.Error("Genesis node is not ordered correctly:", ordering2[2].name, n2.name)
	}
}

// TestPentagon creates a pentagon graph and checks for correctness of
// ordering.
func TestPentagon(t *testing.T) {
	// Create the pentagon shape.
	//
	//   O
	//  / \
	// O   O
	// |   |
	// O   |
	//  \ /
	//   O
	//
	g := NewGraph()
	l1 := g.CreateNode(g.GenesisNode())
	l2 := g.CreateNode(l1)
	r1 := g.CreateNode(g.GenesisNode())
	m := g.CreateNode(l2, r1)

	// Check the ordering on l1.
	ordering1 := l1.relativeOrdering()
	if len(ordering1) != 2 {
		t.Fatal("There should be 2 nodes in the pentagon graph")
	}
	if ordering1[0].name != g.GenesisNode().name {
		t.Error("Genesis node is not ordered correctly:", ordering1[0].name, g.GenesisNode().name)
	}
	if ordering1[1].name != l1.name {
		t.Error("Genesis node is not ordered correctly:", ordering1[1].name, l1.name)
	}

	// Check the ordering on l2.
	ordering2 := l2.relativeOrdering()
	if len(ordering2) != 3 {
		t.Fatal("There should be 3 nodes in the pentagon graph")
	}
	if ordering2[0].name != g.GenesisNode().name {
		t.Error("Genesis node is not ordered correctly:", ordering2[0].name, g.GenesisNode().name)
	}
	if ordering2[1].name != l1.name {
		t.Error("Genesis node is not ordered correctly:", ordering2[1].name, l1.name)
	}
	if ordering2[2].name != l2.name {
		t.Error("Genesis node is not ordered correctly:", ordering2[2].name, l2.name)
	}

	// Check the ordering on r1.
	ordering3 := r1.relativeOrdering()
	if len(ordering3) != 2 {
		t.Fatal("There should be 2 nodes in the pentagon graph")
	}
	if ordering3[0].name != g.GenesisNode().name {
		t.Error("Genesis node is not ordered correctly:", ordering3[0].name, g.GenesisNode().name)
	}
	if ordering3[1].name != r1.name {
		t.Error("Genesis node is not ordered correctly:", ordering3[1].name, r1.name)
	}

	// Check the ordering on m.
	ordering4 := m.relativeOrdering()
	if len(ordering4) != 5 {
		t.Fatal("There should be 5 nodes in the pentagon graph")
	}
	if ordering4[0].name != g.GenesisNode().name {
		t.Error("Genesis node is not ordered correctly:", ordering4[0].name, g.GenesisNode().name)
	}
	if ordering4[1].name != l1.name {
		t.Error("Genesis node is not ordered correctly:", ordering4[1].name, l1.name)
	}
	if ordering4[2].name != l2.name {
		t.Error("Genesis node is not ordered correctly:", ordering4[2].name, l2.name)
	}
	if ordering4[3].name != r1.name {
		t.Error("Genesis node is not ordered correctly:", ordering4[3].name, r1.name)
	}
}

// TestDiamond creates a diamond graph and checks for correctness of ordering.
func TestDiamond(t *testing.T) {
	// Create the diamond shape.
	//
	//   O
	// 	/ \
	// O   O
	//  \ /
	//   O
	//
	g := NewGraph()
	l := g.CreateNode(g.GenesisNode())
	r := g.CreateNode(g.GenesisNode())
	m := g.CreateNode(l, r)

	// Check the ordering on l.
	ordering1 := l.relativeOrdering()
	if len(ordering1) != 2 {
		t.Fatal("There should be 2 nodes in the diamond graph")
	}
	if ordering1[0].name != g.GenesisNode().name {
		t.Error("Genesis node is not ordered correctly:", ordering1[0].name, g.GenesisNode().name)
	}
	if ordering1[1].name != l.name {
		t.Error("Genesis node is not ordered correctly:", ordering1[1].name, l.name)
	}

	// Check the ordering on r.
	ordering2 := r.relativeOrdering()
	if len(ordering2) != 2 {
		t.Fatal("There should be 2 nodes in the diamond graph")
	}
	if ordering2[0].name != g.GenesisNode().name {
		t.Error("Genesis node is not ordered correctly:", ordering2[0].name, g.GenesisNode().name)
	}
	if ordering2[1].name != r.name {
		t.Error("Genesis node is not ordered correctly:", ordering2[1].name, r.name)
	}

	// Check the ordering on m.
	ordering3 := m.relativeOrdering()
	if len(ordering3) != 4 {
		t.Fatal("There should be 4 nodes in the diamond graph")
	}
	if ordering3[0].name != g.GenesisNode().name {
		t.Error("Genesis node is not ordered correctly:", ordering3[0].name, g.GenesisNode().name)
	}
	if ordering3[1].name == l.name {
		// RNG favored l.
		if ordering3[2].name != r.name {
			t.Error("Ordering is incorrect")
		}
	} else if ordering3[1].name == r.name {
		// RNG favored r.
		if ordering3[2].name != l.name {
			t.Error("Ordering is incorrect")
		}
	} else {
		t.Error("Ordering is incorrect")
	}
	if ordering3[3].name != m.name {
		t.Error("Genesis node is not ordered correctly:", ordering3[3].name, m.name)
	}
}
