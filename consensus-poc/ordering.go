package main

import (
	"bytes"
	"crypto/sha256"
)

// nextUnorderedAncestor selects the unordered ancestor which should be ordered
// next from a list of unordered ancestors.
func nextUnorderedAncestor(edges []edgeName, coorespondingChildren []*GraphNode, tip *GraphNode) *GraphNode {
	// The child with the most votes on its edge to the parent wins. If
	// multiple children have the winning number of edge votes, select between
	// them randomly using the hash of the tip block as a seed.
	//
	// As this is a simulation, the names of the blocks are used to seed the
	// rng in lieu of their hashes.
	winningVotes := 0
	var winningHash [32]byte
	var winner *GraphNode
	for i, child := range coorespondingChildren {
		e := edges[i]
		votes := tip.relativeVoteGraph[e]
		childHash := sha256.Sum256([]byte("salt" + tip.name + child.name))
		if votes > winningVotes {
			winningVotes = votes
			winningHash = childHash
			winner = child
		} else if votes == winningVotes && bytes.Compare(winningHash[:], childHash[:]) < 0 {
			winningHash = childHash
			winner = child
		}
	}
	return winner
}

// relativeOrdering returns the sorted graph from the perspective of the
// calling node.
func (gn *GraphNode) relativeOrdering() []*GraphNode {
	// Find the genesis block.
	current := gn
	for len(current.parents) != 0 {
		current = current.parents[0]
	}
	genesis := current

	ordered := make(map[nodeName]bool)
	queued := make(map[nodeName]bool)
	var relativeOrdering []*GraphNode
	var updateOrdering func(base *GraphNode, tip *GraphNode)
	updateOrdering = func(base *GraphNode, tip *GraphNode) {
		queued[base.name] = true
		// Check that all ancestors of the base block are in the ordering. If
		// not, compare all edges from blocks in the ordering to unordered
		// ancestors and select the edge with the most votes. Set the child of
		// that edge to the new base, and set the original base to the new tip.
		for {
			var importantEdges []edgeName
			var correspondingChildren []*GraphNode
			visited := make(map[nodeName]bool)
			unvisited := base.parents
			for len(unvisited) != 0 {
				// Pop a node off of the unvisited stack.
				ancestor := unvisited[0]
				unvisited = unvisited[1:]
				if ordered[ancestor.name] {
					continue
				}

				// Find all edges that point from the ancestor to an ordered
				// node ('importantEdges'). There may not be any. If there are
				// not any, then an unordered ancestor of the unordered
				// ancestor will have an edge that points to an ordered node.
				for _, parent := range ancestor.parents {
					if ordered[parent.name] {
						e := edge(ancestor.name, parent.name)
						importantEdges = append(importantEdges, e)
						correspondingChildren = append(correspondingChildren, ancestor)
					} else if !visited[parent.name] {
						visited[parent.name] = true
						unvisited = append(unvisited, parent)
					}
				}
			}

			// If there are no important edges, all ancestors of the base block
			// have been orderd, exit the ancestor-ordering loop.
			if len(importantEdges) == 0 {
				break
			}

			// Grab a winning child from the list of important edges, and
			// recurse using the winner as the new base, and the current base
			// as the new tip.
			winner := nextUnorderedAncestor(importantEdges, correspondingChildren, tip)
			updateOrdering(winner, base)
		}

		// All ancestors of the base are now ordered. If the base is the tip,
		// there is nothing left to do.
		if base == tip {
			return
		}

		// Add the base to the ordering, then select a new primary child and recurse.
		relativeOrdering = append(relativeOrdering, base)
		ordered[base.name] = true
		_, next := primaryEdge(base, tip)
		if queued[next.name] {
			return
		}
		updateOrdering(next, tip)
	}
	updateOrdering(genesis, gn)
	relativeOrdering = append(relativeOrdering, gn) // Edge case.
	return relativeOrdering
}
