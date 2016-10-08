package main

import (
	"testing"
)

// TestBridgesDiamond checks that the bridges function returns the correct set
// of nodes when applied to nodes following the diamond pattern.
func TestBridgesDiamond(t *testing.T) {
	// Create a diamond graph.
	g := NewLowBlockTimeGraph()
	d1 := g.CreateNode(g.GenesisNode())
	d2 := g.CreateNode(g.GenesisNode())
	m := g.CreateNode(d1, d2)

	// Get the list of bridges in m when d1 and d2 are used are used as the
	// main children.
	d1BridgeEdges, d1BridgeNodes := bridges(m, d1)
	d2BridgeEdges, d2BridgeNodes := bridges(m, d2)
	if len(d1BridgeEdges) != 1 || len(d1BridgeNodes) != 1 {
		t.Fatal("Bridge code returning the wrong number of bridges", len(d1BridgeEdges), len(d1BridgeNodes))
	}
	if len(d2BridgeEdges) != 1 || len(d2BridgeNodes) != 1 {
		t.Fatal("Bridge code returning the wrong number of bridges")
	}
	if d1BridgeEdges[0] != "2-0" {
		t.Error("Bridge nodes returning the wrong nodes as bridges")
	}
	if d2BridgeEdges[0] != "1-0" {
		t.Error("Bridge nodes returning the wrong nodes as bridges")
	}
}

// TestBridgesLeech checks that the bridges function returns the correct set of
// nodes when applied to the leech pattern.
func TestBridgesLeech(t *testing.T) {
	// Create a leech graph.
	g := NewLowBlockTimeGraph()
	h1 := g.CreateNode(g.GenesisNode())
	l1 := g.CreateNode(g.GenesisNode())
	h2 := g.CreateNode(h1)
	l2 := g.CreateNode(l1, h1)
	h3 := g.CreateNode(h2)
	h4 := g.CreateNode(h3)
	l3 := g.CreateNode(l2, h3)
	h5 := g.CreateNode(h4)
	l4 := g.CreateNode(l3, h4)
	m := g.CreateNode(h5, l4)

	fullBridgeEdges, fullBridges := bridges(m, h5)
	if len(fullBridgeEdges) != 4 || len(fullBridges) != 4 {
		t.Fatal("bridges returning wrong number of edges for leech pattern")
	}
}
