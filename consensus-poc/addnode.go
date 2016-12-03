package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
)

// primaryEdge uses the vote graph to determine which edge of the provided
// parent is the primary edge.
func primaryEdge(parent *GraphNode, tip *GraphNode) (edgeName, *GraphNode) {
	// Other nodes may have added edges to the parent which are not actually
	// visible from the tip block. Reduce the set of considered edges to only
	// those edges that are visible from the tip block.
	var visibleChildren []*GraphNode
	for _, child := range parent.children {
		e := edge(child.name, parent.name)
		if _, exists := tip.relativeVoteGraph[e]; exists {
			visibleChildren = append(visibleChildren, child)
		}
	}

	// The child with the most votes on its edge to the parent wins. If
	// multiple children have the winning number of edge votes, select between
	// them randomly using the hash of the tip block as a seed.
	//
	// As this is a simulation, the names of the blocks are used to seed the
	// rng in lieu of their hashes.
	winningVotes := 0
	var winningHash [32]byte
	var winner *GraphNode
	for _, child := range visibleChildren {
		e := edge(child.name, parent.name)
		votes := tip.relativeVoteGraph[e]
		childHash := sha256.Sum256([]byte(tip.salt + string(tip.name+child.name)))
		if votes > winningVotes {
			winningVotes = votes
			winningHash = childHash
			winner = child
		} else if votes == winningVotes && bytes.Compare(winningHash[:], childHash[:]) < 0 {
			winningHash = childHash
			winner = child
		}
	}
	return edge(winner.name, parent.name), winner
}

// CreateNode will take a list of parent nodes, create a graph node from that
// list, and then add that node to the graph, returning the node.
func (g *Graph) CreateNode(parents ...*GraphNode) *GraphNode {
	g.nameCounter++
	tip := &GraphNode{
		name:              nodeName(strconv.Itoa(g.nameCounter)),
		parents:           parents,
		relativeVoteGraph: make(map[edgeName]int),

		salt: g.salt,
	}

	// Perform a DFS on all ancestors of the tip to build the relative vote
	// graph. The relative vote graph of the tip block is the number of votes
	// that each ancestor edge has given the tip block's ancestry.
	visited := make(map[nodeName]bool)
	var addEdges func(parents []*GraphNode, childName nodeName)
	addEdges = func(parents []*GraphNode, childName nodeName) {
		for _, parent := range parents {
			// Add the parent-child edge.
			e := edge(childName, parent.name)
			tip.relativeVoteGraph[e] += 0

			// Skip this parent if the parent has already voted.
			if visited[parent.name] {
				continue
			}
			visited[parent.name] = true

			// Add all of the votes from the parent to the tip.relativeVoteGraph.
			for _, vote := range parent.edgeVotes {
				tip.relativeVoteGraph[vote]++
			}
			addEdges(parent.parents, parent.name)
		}
	}
	addEdges(tip.parents, tip.name)

	// Add the tip block as a child to each of its parents.
	for _, parent := range tip.parents {
		parent.children = append(parent.children, tip)
	}

	// Discover the primary edges for this tip block and vote for them.
	current := g.genesisNode
	for len(current.children) != 0 {
		e, winningChild := primaryEdge(current, tip)
		tip.relativeVoteGraph[e]++
		tip.edgeVotes = append(tip.edgeVotes, e)
		current = winningChild
	}
	return tip
}
