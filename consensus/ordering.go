package main

import (
	"bytes"
	"crypto/sha256"
)

// directUnorderedAncestors returns a list of all ancestors of 'tip' that are
// not ordered, yet are direct children of ordered blocks.
func directUnorderedAncestors(ordered map[nodeName]bool, tip *GraphNode) (duas []*GraphNode) {
	unvisited := tip.parents
	for len(unvisited) != 0 {
		// Pop a node off of the unvisited stack.
		ancestor := unvisited[0]
		unvisited := unvisited[1:]
		if ordered[ancestor.name] {
			// All ancestors of this node are by definition also ordred, no
			// unordered ancestors to be found here.
			continue
		}

		// Find any edges which are pointing from an ordered node to the
		// ancestor. If there is one, this ancestor is a direct unordered
		// ancestor.
		for _, parent := range ancestor.parents {
			if ordered[parent.name] {
				duas = append(duas, ancestor)
				break
			}
		}

		// Add all of the ancestor's parents to the unvisted list.
		unvisited = append(unvisited, ancestor.parents...)
	}
	return duas
}

// nextUnorderedAncestor picks from a list of unordered ancestors the next
// ancestor in the ordering. nextUnorderedAncestor assumes that all input nodes
// are direct children of ordered nodes.
func nextUnorderedAncestor(potentials []*GraphNode, tip *GraphNode) *GraphNode {
	// Use the hash of the tip block and the hash of each child to
	// deterministically choose a winner from the set of potential winners.
	var winningHash [32]byte
	var winningNode *GraphNode
	for _, potential := range potentials {
		pHash := sha256.Sum256([]byte("salt" + tip.name + potential.name))
		if bytes.Compare(winningHash[:], pHash[:]) < 0 {
			winningHash = pHash
			winningNode = potential
		}
	}
	return winningNode
}

// relativeOrdering returns the sorted graph from the perspective of the
// calling node.
func (gn *GraphNode) relativeOrdering() []*GraphNode {
	// Find the genesis block.
	current := gn
	for len(current.parents) != 0 {
		current = current.parents[0]
	}

	// Grab an ordering, grabbing one bock in the primary chain at a time.
	var ordering []*GraphNode
	ordered := make(map[nodeName]bool)
	for {
		// Before 'current' can be added to the ordering, all ancestors must be
		// added to the ordering.
		for duas := directUnorderedAncestors(ordered, current); len(duas) != 0; {
			winner := nextUnorderedAncestor(duas, current)
			ordering = append(ordering, winner)
			ordered[winner.name] = true
		}
		ordering = append(ordering, current)
		ordered[current.name] = true

		// If the original node has been added to the graph, the ordering is
		// complete.
		if current == gn {
			break
		}
		_, current = primaryEdge(current, gn)
	}
	return ordering
}
