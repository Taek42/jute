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

## Shortcomings with Satoshi Consensus

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

Things like [compact
blocks](https://bitcoincore.org/en/2016/06/07/compact-blocks-faq/), [the relay
network](http://bitcoinrelaynetwork.org/),
[FIBRE](http://bluematt.bitcoin.ninja/2016/07/07/relay-networks/), [weak
blocks](http://popeller.io/index.php/2016/01/19/weak-blocks-the-good-and-the-bad/),
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

## Security Requirements of Blockchains

[See slides - work in progress]

## Jute Introduction

Jute allows blocks to have multiple parents, effectively converting the
blockchain into a directed acyclic graph, or DAG. Consensus is heavily
dependent on having a precise ordering for transactions, and security is
heavily dependent on knowing that a highly confirmed transaction cannot be
re-ordered or re-organized unless an impractical amount of additional,
adversarial work is applied to the chain.

Jute is an algorithm for taking DAGs and arriving at a total ordering based on
the amount of work confirming each potential ordering. Once the majority of the
hashrate has agreed upon a particular ordering, that ordering alone receives
votes, which means that a minority hashrate is unable to cause reordering or
reorganization without being exceptionally lucky.

Jute ignores the validity of transactions in the immediate term, allowing any
block with valid proof-of-work to be accepted into the chain. Ignoring
transaction validity is important for eliminating orphans, because multiple
miners finding blocks at the same time should not be heavily penalized if there
is some transaction overlap in those blocks - it will be unavoidable for high
latency miners. This does not need to prevent SPV however. Instead of making
SPV commitments to the transcations in the blocks they are mining (where they
are not certain that the transactions will actually be valid due to potential
double spends), miners can commit to many blocks in recent history, selecting
only the transactions that they perceive as valid from each block. SPV nodes
can easily tell if a reorg has been deep enough to invalidate a particular
commitment, and thus ignores any commitments that may be invalid. The chain is
constructed such that every block will eventually have a valid SPV commitment
to it.

## The Jute Voting & Ordering Algorithms

The genesis block must be the first block in the blockchain. All blocks must
point to either the genesis block or descendents of the genesis block. Blocks
are allowed to have multiple parents. Blocks are not allowed to have parents
which are ancestors of other parents (this makes some of the code cleaner).

![Example DAG](http://i.imgur.com/IHMN24h.jpg)

The block that has no children is called 'the tip'. Though there may be
multiple tips at the same time, only one tip is considered at a time. For the
purposes of consensus, the tip with the most work in its ancestry is considered
the canonical tip, defining the ordering of blocks that is used for consensus.

A direct child-parent connection is called an edge. Each block except for the
tip block will have edges to one or more children. Each time a block is found,
it casts a vote for a set of edges. These votes are then used to determine an
ordering for the blocks.

##### The Jute Voting Algorithm

The voting algorithm starts at the genesis block and works backwards. From the
genesis block, the voting algorithm looks at all edges to children of the
genesis block, and selects the edge with the most votes. Then a vote is added
to that edge, and not to any of the other edges. This is repeated until the tip
block is reached, and a final vote is added to the edge that finds the tip
block. This means that, of all the parents of the tip block, only one edge will
get a vote.

Pseudocode:
```
current := genesisBlock
for current != tip {
    var winningChild
    var maxVotes
    for child := range current.children {
        votes := edge(child, current).numVotes
        if votes > maxVotes {
            winningChild = child
            maxVotes = votes
        }
    }
    edge(winningChild, current).numVotes++
    current = winningChild
}
```

Sometimes, multiple children of a block will have edges with the same number of
votes, and a tiebreaking solution is needed. First, the child with the most
ancestors is preferred. If multiple children have both the same number of on
their edge and the same number of ancestors, then the hash of the merging block
is used to select between the children.

Psuedocode:
```
betweenAllChildren {
    prefer (most votes)
    prefer (most ancestors)
    prefer (lowest result of H(seed || mergingBlockHash || child bock hash)
}
```

Assuming that the block time is not significantly below the network propagation
time, this voting code is sufficient. If the block time is sufficiently below
the network propagation time, a well-networked minority hashrate attacker can
begin censoring the blockchain, preventing any blocks from the honest majority
from ever getting confirmed. The following psuedocode describes an extension to
the voting which prevents this attack:

[wip]

A full golang implementation can be found in [consensus-poc/addnode.go](consensus-poc/addnode.go)

---

The code implements what is described above, however I believe there is an
optimization. When choosing between two potential main chains in the event of a
tie between chains, instead of preferring the number of ancestors of the child,
you should prefer the number of ancestors in the entire potential main chain.
This may introduce intractable computational complexity, or may introduce a
security vulnerability, I haven't thought it through fully.

##### The Jute Sorting Algorithm

[wip]

A full golang implementation can be found in [consensus-poc/sort.go](consensus-poc/sort.go)

## Intuition Around the Security of Jute

[see slides - work in progress]

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

Thanks to Andrew Poelstra for review + corrections.  
Thanks to Bob McElrath for active research on DAG based consensus, which has influenced this document.
