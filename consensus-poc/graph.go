package main

import (
	"crypto/rand"
	"fmt"
	"strconv"
	"strings"
)

// nodeName is the string version of a node's numerical counter.
type nodeName string

// edgeName is the name of an edge connecting two nodes. The name takes the
// form "childName" + "-" + "parentName".
type edgeName string

// edge returns the name of the edge that is created when the child commits to
// the parent.
func edge(child, parent nodeName) edgeName {
	return edgeName(child + "-" + parent)
}

// GraphNode defines a node in a graph. The algorithm needs to traverse the
// graph both forward and backwards, so parents must point to all children, and
// children must point to all parents.
//
// Each node has the full list of votes for all edges as the were when this
// block was the tip block.
type GraphNode struct {
	// These values are inherent to the node and will not change.
	//
	// The edgeVotes indicates the edges that this node casts a vote for.
	name              nodeName
	edgeVotes         []edgeName
	parents           []*GraphNode
	relativeVoteGraph map[edgeName]int

	// The set of children can grow as more nodes are added to the graph.
	// Adding a child will never change the relative vote graph, or the voting
	// decisions of the parent.
	children []*GraphNode

	// Each node knows the graph salt.
	salt string
}

// Graph is the base type that is used to build out a graph of nodes.
type Graph struct {
	// nameCounter enables the graphViewer to assign unique names to each node.
	nameCounter int

	// genesisNode is the oldest node in the tree.
	genesisNode *GraphNode

	// the salt used to make sure that rng decisions are different from
	// run-to-run, especially useful during testing.
	salt string
}

// GenesisNode returns the genesis node of the graph.
func (g *Graph) GenesisNode() *GraphNode {
	return g.genesisNode
}

// NewGraph initializes a graph with a genesis node that has no children and
// returns the graph.
func NewGraph() *Graph {
	saltBase := make([]byte, 32)
	rand.Read(saltBase)
	return &Graph{
		nameCounter: 0,
		genesisNode: &GraphNode{
			name:              "0",
			edgeVotes:         make([]edgeName, 0),
			parents:           make([]*GraphNode, 0),
			relativeVoteGraph: make(map[edgeName]int),

			children: make([]*GraphNode, 0),
		},
		salt: string(saltBase),
	}
}
