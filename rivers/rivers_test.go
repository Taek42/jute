package main

import (
	"testing"
)

// TestDiamond checks that the sorting and voting results match what is
// expected after a diamond is created.
func TestDiamond(t *testing.T) {
	// Create a diamond graph.
	g := NewLowBlockTimeGraph()
	d1 := g.CreateNode(g.GenesisNode())
	d2 := g.CreateNode(g.GenesisNode())
	m := g.CreateNode(d1, d2)

	// Check the number of parents for each node.
	if len(g.GenesisNode().parents) != 0 {
		t.Error("genesis node has wrong number of parents")
	}
	if len(d1.parents) != 1 {
		t.Error("d1 has wrong number of parents")
	}
	if len(d2.parents) != 1 {
		t.Error("d1 has wrong number of parents")
	}
	if len(m.parents) != 2 {
		t.Error("m has wrong number of parents")
	}

	// Check the number of children.
	if len(g.GenesisNode().children) != 2 {
		t.Error("genesis node has wrong number of hcildren")
	}
	if len(d1.children) != 1 {
		t.Error("d1 has wrong number of children")
	}
	if len(d2.children) != 1 {
		t.Error("d2 has wrong number of children")
	}
	if len(m.children) != 0 {
		t.Error("m has children")
	}

	if m.RelativeOrdering() != "0-1-2-3" && m.RelativeOrdering() != "0-2-1-3" {
		t.Error("m has incorrect relative ordering", m.RelativeOrdering())
	}

	// Repeat the tests, but with a non-LBT graph. The two should have
	// identical results.
	g = NewGraph()
	d1 = g.CreateNode(g.GenesisNode())
	d2 = g.CreateNode(g.GenesisNode())
	m = g.CreateNode(d1, d2)

	// Check the number of parents for each node.
	if len(g.GenesisNode().parents) != 0 {
		t.Error("genesis node has wrong number of parents")
	}
	if len(d1.parents) != 1 {
		t.Error("d1 has wrong number of parents")
	}
	if len(d2.parents) != 1 {
		t.Error("d1 has wrong number of parents")
	}
	if len(m.parents) != 2 {
		t.Error("m has wrong number of parents")
	}

	// Check the number of children.
	if len(g.GenesisNode().children) != 2 {
		t.Error("genesis node has wrong number of hcildren")
	}
	if len(d1.children) != 1 {
		t.Error("d1 has wrong number of children")
	}
	if len(d2.children) != 1 {
		t.Error("d2 has wrong number of children")
	}
	if len(m.children) != 0 {
		t.Error("m has children")
	}

	if m.RelativeOrdering() != "0-1-2-3" && m.RelativeOrdering() != "0-2-1-3" {
		t.Error("m has incorrect relative ordering", m.RelativeOrdering())
	}
}

// TestLBTBranching checks that a graph designed to trigger the LowBlockTime
// causes the expected sorting in both the case of LBT extensions enabled and
// disabled.
//
// Note: The LBT extensions should not do any harm for high block times, and
// only exist as a branch because it adds a fair amount of complexity to the
// code. There should be no disadvantages to having it enabled all the time.
func TestLBTBranching(t *testing.T) {
	// Create the 'LBT' shaped graph (small graph which exhibits the difference
	// between the LBT code and the non-LBT code) using a graph without LBT
	// sorting enabled.
	g := NewGraph()
	w1 := g.CreateNode(g.GenesisNode())
	w2 := g.CreateNode(g.GenesisNode())
	w3 := g.CreateNode(g.GenesisNode())
	w4 := g.CreateNode(g.GenesisNode())
	n1 := g.CreateNode(g.GenesisNode())
	n2 := g.CreateNode(n1)
	m1 := g.CreateNode(w1, w2, w3, w4, n2)
	n3 := g.CreateNode(n2)
	n4 := g.CreateNode(n3)
	m2 := g.CreateNode(m1, n4)

	// Verify that the relative ordering to M has prioritized the n branch.
	m2RO := m2.RelativeOrdering()
	if len(m2RO) != 22 {
		t.Error("wrong character count for m2 ordering")
	}
	// We check a prefix and a suffix. Do to rng selection, the middle 4 nodes
	// are allowed to be in any order.
	if m2RO[:10] != "0-5-6-8-9-" {
		t.Error("wrong prefix for m2 ordering", m2RO[:10])
	}
	if m2RO[18:] != "7-10" {
		t.Error("wrong suffix for m2 ordering", m2RO[18:])
	}

	// Create the 'LBT' shaped graph (small graph which exhibits the difference
	// between the LBT code and the non-LBT code) using a graph with the LBT
	// sorting enabled.
	lg := NewLowBlockTimeGraph()
	lw1 := lg.CreateNode(lg.GenesisNode())
	lw2 := lg.CreateNode(lg.GenesisNode())
	lw3 := lg.CreateNode(lg.GenesisNode())
	lw4 := lg.CreateNode(lg.GenesisNode())
	ln1 := lg.CreateNode(lg.GenesisNode())
	ln2 := lg.CreateNode(ln1)
	lm1 := lg.CreateNode(lw1, lw2, lw3, lw4, ln2)
	ln3 := lg.CreateNode(ln2)
	ln4 := lg.CreateNode(ln3)
	lm2 := lg.CreateNode(lm1, ln4)

	lm2RO := lm2.RelativeOrdering()
	if len(m2RO) != 22 {
		t.Error("wrong character count for m2 ordering")
	}
	// We check a prefix and a suffix. Do to rng selection, the middle 4 nodes
	// are allowed to be in any order.
	if lm2RO[:6] != "0-5-6-" {
		t.Error("wrong prefix for lm2 ordering", lm2RO[:6])
	}
	if lm2RO[14:] != "7-8-9-10" {
		t.Error("wrong suffix for lm2 ordering", lm2RO[14:])
	}
}
