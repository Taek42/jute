package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
)

// nextMainNode uses the vote graph and hash of the tip block to select between
// the children of the provided parent. The selected child is the next block in
// the main chain.
func nextMainNode(parent *GraphNode, tip *GraphNode) *GraphNode {
	// Because of other nodes, the parent may have children which are not in
	// the edge graph of the tip. Separate the set of children recognized by
	// the tip node from the set of all children to the parent.
	var visibleChildren []*GraphNode
	for _, child := range parent.children {
		e := edge(child.name, parent.name)
		if _, exists := tip.relativeVoteGraph[e]; exists {
			visibleChildren = append(visibleChildren, child)
		}
	}

	// The child with the most votes on its edge to the parent wins. If
	// multiple children have the winning number of edge votes, select the
	// child with the greatest relative height. If multiple children have both
	// the winning number of edge votes and the winning relative height, select
	// between them randomly using the hash of the tip block as a seed.
	//
	// As this is a simulation, the names of the blocks are used to seed the
	// rng in lieu of their hashes.
	winningVotes := 0
	winningHeight := 0
	var winningHash [32]byte
	var winner *GraphNode
	for _, child := range visibleChildren {
		e := edge(child.name, parent.name)
		votes := tip.relativeVoteGraph[e]
		childHash := sha256.Sum256([]byte("salt" + tip.name + child.name))
		if votes > winningVotes {
			winningVotes = votes
			winningHeight = child.relativeHeight
			winningHash = childHash
			winner = child
		} else if votes == winningVotes && child.relativeHeight > winningHeight {
			winningHeight = child.relativeHeight
			winningHash = childHash
			winner = child
		} else if votes == winningVotes && child.relativeHeight == winningHeight && bytes.Compare(winningHash[:], childHash[:]) < 0 {
			winningHash = childHash
			winner = child
		}
	}
	return winner
}

// CreateNode will take a list of parent nodes, create a graph node from that
// list, and then add that node to the graph, returning the node.
func (g *Graph) CreateNode(parents ...*GraphNode) *GraphNode {
	g.nameCounter++
	tip := &GraphNode{
		name:              nodeName(strconv.Itoa(g.nameCounter)),
		parents:           parents,
		relativeHeight:    1,
		relativeVoteGraph: make(map[edgeName]int),
	}

	// Define a recursive helper function to fetch the votes of the previous
	// nodes and create a voteGraph. The optimized version of this code is O(n)
	// in the number of parents instead of O(n) in the total number of edges in
	// the graph.
	visited := make(map[nodeName]bool)
	var addEdges func(parents []*GraphNode, childName nodeName)
	addEdges = func(parents []*GraphNode, childName nodeName) {
		for _, parent := range parents {
			// Skip this parent if the parent has already voted.
			if visited[parent.name] {
				continue
			}
			visited[parent.name] = true

			// Add all of the votes from the parent to the tip.relativeVoteGraph.
			for _, vote := range parent.edgeVotes {
				tip.relativeVoteGraph[vote]++
			}

			// Add the parent-child edge with zero votes if it has not yet
			// received any votes.
			e := edge(childName, parent.name)
			tip.relativeVoteGraph[e] += 0
			addEdges(parent.parents, parent.name)
		}
	}
	// Perform a DFS on the parents and count the total number of votes for
	// each edge in the graph visible to the new node.
	addEdges(parents, tip.name)

	// Update the relative height of the node to reflect its actual relative
	// height.
	tip.relativeHeight += len(visited)

	// Add the child to all of its parents.
	for _, parent := range tip.parents {
		parent.children = append(parent.children, tip)
	}

	// Execute the jute voting algorithm, iteratively detecting the main chain
	// and then voting for it.
	current := g.genesisNode
	for len(current.children) != 0 {
		winner := nextMainNode(current, tip)
		e := edge(winner.name, current.name)
		tip.relativeVoteGraph[e]++
		tip.edgeVotes = append(tip.edgeVotes, e)
		current = winner

		// An additional rule allows us to protect against low-latency
		// minorities. The newest edge in the main chain without the extra rule
		// will always have just one vote. The extra rule is necessary if the
		// block time is substantially lower than the network propagation time.
		if g.lowBlockTime {
			// TODO
		}
	}

	// Voting complete, graph weights updated.
	return tip
}
