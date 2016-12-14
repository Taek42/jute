Jute: More Scalable, More Decentralized Proof-of-Work Consensus
===============================================================

**Used in conjunction with [this presentation](Scaling Bitcoin Milan Presentation.pdf)**

Jute is a proof-of-work consensus algorithm drawing heavily from the Bitcoin
consensus algorithm. Jute assumes a network of economically rational miners
where no single party or conglomerate controls more than 51% of the hashate.
The jute algorithm solves the problem of selfish mining, eliminates orphans,
and paves a path leading toward shorter block times and higher network
throughput. This is achieved by replacing the linked-list, longest chain
consensus algorithm with an algorithm that allows blocks to have many parents,
creating a DAG. A sorting algorithm is applied to the DAG to get an exact
ordering of blocks that is safe from reordering/reorganization except in the
face of a 51% attacker.

The result is a consensus algorithm which eliminates orphan-rate based miner
centralization, eliminates orphan based selfish mining attacks, allows the
block time to be safely reduced by a substantial amount, and allows the
throughput of the network to be safely increased by a substantial amount.

## Related work

A disclaimer: though a lot of the related work is interesting and was very
helpful in guiding the design of jute, much of the work linked below has
security problems or other shortcomings that are not discussed in the links
provided.   

[Weak Blocks - The Good and the Bad](http://popeller.io/index.php/2016/01/19/weak-blocks-the-good-and-the-bad/) [(more)](https://gist.github.com/jl2012/2db9e434ce7beb8e13354604305f6d77) [(more)](https://gist.github.com/erasmospunk/23040383b7620b525df0) [(more)](https://diyhpl.us/wiki/transcripts/scalingbitcoin/hong-kong/invertible-bloom-lookup-tables-and-weak-block-propagation-performance/) [(mandatory weak blocks)](https://lists.linuxfoundation.org/pipermail/bitcoin-dev/2015-November/011707.html)  
[Braiding the Blockchain](https://scalingbitcoin.org/hongkong2015/presentations/DAY2/2_breaking_the_chain_1_mcelrath.pdf) [(transacript)](https://diyhpl.us/wiki/transcripts/scalingbitcoin/hong-kong/braiding-the-blockchain/)  
[Inclusive Block Chain Protocols](http://fc15.ifca.ai/preproceedings/paper_101.pdf)  
[bip-soft-blocks.mediawiki](https://gist.github.com/erasmospunk/23040383b7620b525df0)  
[Block Publication Incentives for Miners](https://petertodd.org/2016/block-publication-incentives-for-miners)  
[Accelerating Bitcoin's Transaction Processing (GHOST)](https://eprint.iacr.org/2013/881.pdf)  
[Secure High Rate Transaction Processing in Bitcoin (GHOST)](http://fc15.ifca.ai/preproceedings/paper_30.pdf)  
[DAG-Coin](https://bitcointalk.org/index.php?topic=1177633.0) [(draft paper)](https://bitslog.files.wordpress.com/2015/09/dagcoin-v41.pdf) [(blog post)](https://diyhpl.us/wiki/transcripts/scalingbitcoin/hong-kong/braiding-the-blockchain/)  
[Bitcoin-NG: A Scalable Blockchain Protocol](https://arxiv.org/pdf/1510.02037.pdf)  
[Ethereum GHOST Protocol](https://blog.ethereum.org/2014/07/11/toward-a-12-second-block-time/) [(in whitepaper)](https://github.com/ethereum/wiki/wiki/White-Paper#modified-ghost-implementation)  
[Re: [Bitcoin-development] Tree-chains preliminary summary](https://www.mail-archive.com/bitcoin-development@lists.sourceforge.net/msg04388.html)  
[The Tangle](https://www.weusecoins.com/assets/pdf/library/Tangle%20-%20a%20cryptocurrency%20for%20Internet-of-Things%20industry%20-%20blockchain%20alternative.pdf)  
[Reduce Orphaning Risk and Improve Zero-Confirmation Security with Subchains](https://www.bitcoinunlimited.info/resources/subchains.pdf)

## Shortcomings with Nakamoto Consensus

##### Orphans

Blocks take time to propagate. The miner who finds a block can propagate it to
all their local nodes almost instantly. The rest of the network must wait for
the block to travel the network, spending more time working on an outdated and
having a higher orphan rate. Miners with more hashrate get a bigger boost from
their instant propagation, and they get that boost more frequently (as they
find more blocks). This is a centralization pressure. When the block time is
substantially higher than the network propagation time (as it is in Bitcoin),
the effect is greatly reduced.

Jute all but eliminates orphans, eliminating this centralizing effect along
with them. One of the chief reasons that a high block time is necessary gets
eliminated, and a nontrivial mining centralization pressure is eliminated as
well.

##### Selfish Mining

Miners with sufficient hashrate or sufficient network superiority are able to
perform strategic block witholding and block propagation strategies such that
their profitability increases by them increasing the relative orphan rates of
their competitors. Because Jute mostly eliminates orphans,  orphan-based
selfish mining is no longer an issue.

###### Papers and Links Related to Selfish Mining:
[Majority is Not Enough: Bitcoin Mining is Vulnerable](https://www.cs.cornell.edu/~ie53/publications/btcProcFC.pdf)  
[Optimal Selfish Mining Strategies in Bitcoin](http://fc16.ifca.ai/preproceedings/30_Sapirshtein.pdf)  
[Stubborn Mining: Generalizing Selfish Mining and Combining with an Eclipse Attack](https://eprint.iacr.org/2015/796.pdf)  
[Block Publication Incentives for Miners](https://petertodd.org/2016/block-publication-incentives-for-miners)  

##### Poorly Utilized Network

The bitcoin network spends most of it's life idle. Every 10 minutes, a 1MB
block propagates, and then the majority of the nodes are quiet again for
approximately the whole next 10 minutes. Miners especially must have low
latency setups, which means mining in certain parts of the world becomes
infeasible (or at least less profitable), and it means that the amount of
initial capital needed for mining is higher. Large miners can more easily
amortize initial capital costs meaning that larger, more centralized miners
experience higher per-hashrate profitability.

Things like [compact blocks](https://bitcoincore.org/en/2016/06/07/compact-blocks-faq/),
[the relay network](http://bitcoinrelaynetwork.org/),
[FIBRE](http://bluematt.bitcoin.ninja/2016/07/07/relay-networks/),
[weak blocks](http://popeller.io/index.php/2016/01/19/weak-blocks-the-good-and-the-bad/),
and other fast-relay schemes all help to improve network utilization, however
most of these strategies rely on pre-forwarding transactions and do not handle
adversarial blocks very well.

With jute, miners can remain competitive without seeing their competitors
blocks immediately. The reduced emphasis on network latency means that the
block time can be lowered, and optimizations can focus more heavily on
throughput. The lowered blocktime also means a more consistent load applied to
the network, which itself is an optimization for throughput.

###### Papers and Links Related to Bitcoin's Network:
[On Scaling Decentralized Blockchains (A Position Paper)](http://fc16.ifca.ai/bitcoin/papers/CDE+16.pdf)

##### High Variance Mining Rewards

The long block times in Bitcoin means that smaller mining operations may only
be able to find a few blocks per month or year, following a poisson
distribution. This means that there is very high variance in month-to-month
income for small solo miners, which makes solo mining substantially less
practical without many tens of thousands of dollars of mining hardware. Bitcoin
addresses this primarily thorugh the use of mining pools, which is a
centralization pressure.

The lower block time of Jute means more payouts to miners, greatly reducing the
minimum-hashpower requirement for practical solo mining, reducing a significant
miner centralization pressure in the Bitcoin ecosystem.

##### High Variance Confirmation Times

The time between blocks follows an exponential distribution. It's common for
blocks to be more than 30 minutes apart. This hurts the user experience, as the
difference between 10 minutes and 30 minutes can be significant in daily life.
By lowering the block time, Jute reduces the variance in confirmation times.

It should be noted that a single confirmation in Jute is much weaker than a
single confirmation in Bitcoin. Instead, one must wait until a particular
ordering is being endorsed by a majority of the hashrate, which can take a
couple of network propagation cycles. Confirmation times with jute are still in
the ballpark of 5-10 minutes, but with much lower variance.

## Requirements for a Competitor to Nakamoto Consensus

The goal of jute is to replace Nakamoto consensus as the primary mechanism for
acheiving decentralized consensus. To be a proper replacement, it must be able
to provide all of the benefits that Nakamoto consensus can provide, as well as
provide additional features that Nakamoto consensus is unable to provide.

##### Immutable History

Once history is added to consensus with a sufficient number of confirmations,
it must be provably difficult to alter that history. Each additional
confirmation should confirm only that history, and the network should converge
such that the entire network is confirming only that history. Additionally, the
amount of work required to alter the history must be equal to the total amount
of work confirming the history.

##### Incentive Compatibility

Miners should make the optimal amount of money by following the protocol as
prescribed. Deviations from the protocol should either be non-harmful to
network, or more ideally should result in a significant reduction in
profitability.

The incentive compatibility ideally applies to miners having up to 50% of the
hashrate.

##### Censorship Resistance

Groups controlling less than 50% of the hashrate should be unable to prevent a
transaction from being included in the network.

##### Fast Double Spend Resisitance

Confidence around the security of a transaction should be achieved consistently
in under two hours. Confidence around smaller transactions should be achieved
consistently in under half an hour.

##### Accessible Resource Requirements

The median consumer desktop connected via a median internet connection should
be able to participate as a fully validating node, even when the network is
under attack. This means that in the worst case, only a moderate amount of
strain is placed on the CPU, the memory, the hard drive, and the network
bandwidth.

##### SPV Support

Devices with limited resources, such as cell phones, should be able to verify
incoming transactions by only adding the assumption that the history with the
most work is a valid history.

##### Hashrate Profitability Fairness

Miners with more hashrate should not see higher profitability per-hashrate. Put
another way, miners should see a perfectly linear increase in profitability as
they centralize and add hashrate, as opposed to seeing superlinear
profitability gains.

##### Network Profitability Fairness

Miners with superior network connections should not see higher profitability
compared to miners with inferior network connections, so long as the inferior
miner has as much connectivity as the average consumer.

## Jute Introduction

Jute is a block based consensus algorithm where blocks are allowed to have
multiple parents, effectively converting the blockchain into a directed
acyclic graph or DAG.

The DAG is converted into a linear history through a specific sorting algorithm
(explained below). Using that algorithm as a foundation, we are able to achieve
all of the security properties described above.

Jute is an inclusive consensus protocol, which means all blocks are accepted
into the linear history, even blocks that have invalid transactions. Nodes are
able to determine which transactions are valid by looking at the final ordering
and ignoring any invalid transactions.

Miners are able to commit to the valid transactions such that SPV nodes are
able to securely tell which transactions within a block are invalidated by
conflicting history. This means that SPV nodes can achieve secuirty equivalent
to the security of SPV nodes in Nakamoto consensus, even though the history can
contain invalid transactions.

## The Jute Sorting Algorithm

The genesis block must be the first block in the blockchain. All other blocks
must have the genesis block as either a direct or indirect ancestor. Blocks are
allowed to have multiple parents. Blocks are not allowed to have parents which
are ancestors of other parents.

![Example DAG](http://i.imgur.com/IHMN24h.jpg)

A block that has no children is called a tip block. Multiple tip blocks may
exist simultaneously, resulting in multiple alternate histories. The sorting
algorithm provided below indicates which tip block is considered the canonical
history. A miner will add every known tip block as a parent of the block they
are working on.

A direct child-parent connection is called an edge. Each block except for the
tip block will have edges to one or more children.

### Weighting Edges

Jute establishes a linear ordering for blocks by identifying and leveraging
primary edges. Each block that is added to the consensus DAG gets to vote for a
set of primary edges, adding a single vote to each priamry edge.

Primary edges are chosen starting from the genesis block. All edges from the
genesis block to child blocks are considered, and the single edge with the most
votes is chosen as the primary edge. That edge receives another vote, and then
the process is repeated until the new block is reached, at which point the
voting is complete.

If multiple edges have the same winning number of votes, one is selected
randomly using a random number generator seeded by the hash of the parent block
appended to the hash of the new block.

##### Pseudocode for Primary Edge Voting:
```
var newBlock // a newly solved block extending the chain
current := genesisBlock
for current != newBlock {
    winningChildren := {}
    maxVotes := -1
    for child := range current.children {
        votes := edge(child, current).numVotes
        if votes > maxVotes {
            winningChildren = {child}
            maxVotes = votes
        } else if votes == maxVotes {
			winningChildren = append(winningChildren, child)
		}
    }
	if len(winningChildren) > 1 {
		rng := seedRNG(Hash(newBlock, current))
		winningChildren = rng.Randomize(winningChildren)
	}

    edge(winningChildren[0], current).numVotes++
    current = winningChild
}
```

A full golang implementation can be found in [consensus-poc/addnode.go](consensus-poc/addnode.go)

##### Security Intuition Around the Edge Voting

The primary edges are selected by popularity. Once the network has converged on
a particular edge as the primary edge, every block that gets produced will vote
for that same edge. That means that the gap between that edge and any competing
edge will be growing at a rate of 1 block per block that is generated by the
network.

Any miner attempting to cause an alternate history will have to produce enough
blocks to overcome the inital difference between the edges, while also keeping
up with the entire rest of the network as new blocks get produced. This
requires more than 50% hashrate.

This is equivalent to the immutability provided by Nakamoto consensus - once
all miners are confirming a certain piece of history, either luck or a full 50%
of the network hashrate is required to alter that history.

The probability that a miner with less than 50% hashrate can alter history
decreases exponentially in the number of confirmations on that history. It is
not unreasonable that a 25% hashrate miner would get lucky enough to alter
history with 3 confirmations, but it is exceedingly unlikely that a 25%
hashrate miner would be able to alter history with 1000 confirmations.

### The Jute Ordering Algorithm

Blocks are ordered in Jute based on the edge votes. The genesis block is set as
the first block in history. Then the primary edge from the genesis block to a
child is identified using the same method as above. We will label the
associated child as the primary child.

All ancestors of the primary child must be included in the ordering before the
child can be included. Once all ancestors are included, the next primary child
is identified and the algorithm is repeated until a tip block is reached and
all ancestors of that tip block have been included in the ordering.

If the primary child has multiple unordered ancestors, all edges leading from a
block in the ordering directly to an unordered ancestor are observed. The votes
are tallied on each edge as seen by the primary child (meaning that votes from
decendents of or siblings of the primary child are ignored). The edge with the
greatest number of votes is selected, tiebreaking using a random number
generator seeded by the hash of the primary child and the unordered ancestor.
Recursively, the winner is set as the new primary child, and the ordering
algorithm is applied until the original primary child is reached. If the
original primary child still has unordered ancestors at that point, the
algorithm is repeated until the primary child has no more unordered ancestors.

Psuedocode has been omitted for the ordering algorithm, as it's not very
helpful. A full golang implementation can be found in
[consensus-poc/sort.go](consensus-poc/ordering.go)

### Practical Jute Today

+ Deploy the jute consensus algorithm
+ Set the block time to 6 seconds
+ Set the block size to 50kb
+ Set SPV to commit to block ranges after they have 50+ confirmations (5 minutes)
+ Don't allow merges of blocks if their most recent parent already has 200+ confirmations
+ Ignore miner fee complications because they shouldn't come into play too much at 5 second confirmation times

Realistically, by the time this was implemented we'd have worked through a lot
more of the issues. I have casually handwaved over things like the need to
update the difficulty adjustment algorithm, the timestamp algorithm, and
probably a whole host of other protocol-level things that this interferes with.
Largely though I don't think those are hard things to adjust, you just need to
be careful and methodical.

#### Acknowledgements

Thanks to Jonathan Harvey-Buschel for contributions + research collaboration.
Thanks to Andrew Poelstra for review + corrections.  
Thanks to Bob McElrath for active research on DAG based consensus, which has influenced this document.
