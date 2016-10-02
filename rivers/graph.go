package main

import (
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
// block was the tip block. The space complexity for this is O(n^2) in the
// total number of nodes, however using a global table and some clever code we
// could optimize the space complexity to O(n).
type GraphNode struct {
	// These values are inherent to the node and will not change.
	//
	// The name must be unique.
	//
	// The edgeVotes indicates the edges that this node casts a vote for.
	//
	// The relative height is the total number of blocks that can be
	// reached by performing a depth-first-search on the parents.
	name              nodeName
	edgeVotes         []edgeName
	parents           []*GraphNode
	relativeHeight    int
	relativeVoteGraph map[edgeName]int

	// The set of children can grow as more nodes are added to the graph.
	// Adding a child will never change the relative vote graph, or the voting
	// decisions of the parent.
	children []*GraphNode
}

// RelativeOrdering sorts the graph using the supplied node as the tip, then
// prints the ordering that results.
func (gn *GraphNode) RelativeOrdering() string {
	relativeOrdering := gn.relativeOrdering()
	s := fmt.Sprint(relativeOrdering[0].name)
	for i := 1; i < len(relativeOrdering); i++ {
		s = fmt.Sprint(s, "-")
		s = fmt.Sprint(s, relativeOrdering[i].name)
	}
	return s
}

// Graph contains the genesis block, and helper functions for adding nodes to
// the graph. When a node is added, its vote will be cast and a relative
// ordering for that node will be determined.
type Graph struct {
	// nameCounter enables the graphViewer to assign unique names to each node.
	nameCounter int

	// genesisNode is the oldest node in the tree.
	genesisNode *GraphNode

	// lowBlockTime indicates whether the lowBlockTime rule is applied during
	// the voting process.
	lowBlockTime bool
}

// GenesisNode returns the genesis node of the graph.
func (g *Graph) GenesisNode() *GraphNode {
	return g.genesisNode
}

// NewGraph initializes a graph with a genesis node that has no children and
// returns the graph.
func NewGraph() *Graph {
	return &Graph{
		nameCounter: 0,
		genesisNode: &GraphNode{
			name:              "0",
			edgeVotes:         make([]edgeName, 0),
			parents:           make([]*GraphNode, 0),
			relativeHeight:    1,
			relativeVoteGraph: make(map[edgeName]int),

			children: make([]*GraphNode, 0),
		},
	}
}

// NewLowBlockTimeGraph initalizes and returns a graph that has an extra rule
// for protecting high latency miners when the block time is significantly
// lower than the network propagation time.
func NewLowBlockTimeGraph() *Graph {
	g := NewGraph()
	g.lowBlockTime = true
	return g
}

// SageGen returns a string that can be fed into Sage to create a visualization
// of the longest chain and the votes for each edge in that chain.
func (g *Graph) SageGen() string {
	// Find the node with the highest relative height.
	winningChild := g.genesisNode
	winningHeight := 1
	looseEnds := g.genesisNode.children
	for len(looseEnds) != 0 {
		le := looseEnds[0]
		looseEnds = looseEnds[1:]
		if le.relativeHeight > winningHeight {
			winningHeight = le.relativeHeight
			winningChild = le
		}
		looseEnds = append(looseEnds, le.children...)
	}

	// Now that we have the winning child, print out the code of the graph
	// of the winning child.
	s := fmt.Sprintln("G = DiGraph()")
	for edge, weight := range winningChild.relativeVoteGraph {
		// Parse the edge name into its components.
		nodes := strings.Split(string(edge), "-")
		s = fmt.Sprint(s, "G.add_edge("+nodes[0]+", "+nodes[1]+", "+strconv.Itoa(int(weight))+")\n")
	}
	relativeOrdering := winningChild.RelativeOrdering()
	s = fmt.Sprint(s, "H = G.plot(edge_labels=True, layout='acyclic', edge_color='grey')\n")
	s = fmt.Sprint(s, "H.show(title=\""+relativeOrdering+"\", figsize=(5,16))\n")
	filename := relativeOrdering + ".png"
	s = fmt.Sprintf("%sH.save(filename=\"/home/user/plots/%s\", title=\""+relativeOrdering+"\", figsize=(5,16))\n", s, filename)
	return s
}
