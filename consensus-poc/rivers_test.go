package main

import (
	"testing"
)

// TestDiamond checks that the sorting and voting results match what is
// expected after a diamond is created.
func TestDiamond(t *testing.T) {
	// Repeat the tests, but with a non-LBT graph.
	g := NewGraph()
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
		t.Error("d2 has wrong number of parents")
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

	// Check the edge strengths.
	e1 := edge(d1.name, g.GenesisNode().name)
	e2 := edge(d2.name, g.GenesisNode().name)
	e3 := edge(m.name, d1.name)
	e4 := edge(m.name, d2.name)
	if m.relativeVoteGraph[e1] == 2 {
		// RNG favored e1
		if m.relativeVoteGraph[e2] != 1 {
			t.Error("wrong number of votes for e2", m.relativeVoteGraph[e2])
		}
		if m.relativeVoteGraph[e3] != 1 {
			t.Error("wrong number of votes for e3", m.relativeVoteGraph[e3])
		}
		if m.relativeVoteGraph[e4] != 0 {
			t.Error("wrong number of votes for e4", m.relativeVoteGraph[e4])
		}
	} else {
		// RNG favored e2
		if m.relativeVoteGraph[e2] != 2 {
			t.Error("base votes seem incoorect", m.relativeVoteGraph[e1], m.relativeVoteGraph[e2])
		}
		if m.relativeVoteGraph[e1] != 1 {
			t.Error("wrong number of votes for e2", m.relativeVoteGraph[e2])
		}
		if m.relativeVoteGraph[e4] != 1 {
			t.Error("wrong number of votes for e4", m.relativeVoteGraph[e4])
		}
		if m.relativeVoteGraph[e3] != 0 {
			t.Error("wrong number of votes for e3", m.relativeVoteGraph[e3])
		}
	}

	if m.RelativeOrdering() != "0-1-2-3" && m.RelativeOrdering() != "0-2-1-3" {
		t.Error("m has incorrect relative ordering", m.RelativeOrdering())
	}
}
