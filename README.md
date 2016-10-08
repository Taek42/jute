 
Jute: More Scalable, More Decentralized Proof-of-Work Consensus
===============================================================

**Used in conjunction with [ppt]**

Jute is a proof-of-work consensus algorithm drawing heavily from the Bitcoin consensus algorithm. Jute assumes a network of economically rational miners where no single party or conglomerate controls more than 51% of the hashate. The jute algorithm solves the problem of selfish mining, eliminates orphans, and paves a path leading toward shorter block times and higher network throughput. This is achieved by replacing the linked-list, longest chain consensus algorithm with an algorithm that allows blocks to have many parents, creating a DAG. A sorting algorithm is applied to the DAG to get an exact ordering of blocks that is safe from reordering/reorganization except in the face of a 51% attacker.

The result is a consensus algorithm which eliminates orphan-rate based miner centralization, eliminates orphan based selfish mining attacks, allows the block time to be safely reduced by a substantial amount, and allows the throughput of the network to be safely increased by a substantial amount.

## Related work

A disclaimer: though a lot of the related work is interesting and was very helpful in guiding the design of jute, much of the work linked below has security problems or other shortcomings that are not discussed in the links provided.   

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

Blocks take time to propagate. The miner who finds a block can propagate it to all their local nodes almost instantly. The rest of the network must wait for the block to travel the network, spending more time working on an outdated and having a higher orphan rate. Miners with more hashrate get a bigger boost from their instant propagation, and they get that boost more frequently (as they find more blocks). This is a centralization pressure. When the block time is substantially higher than the network propagation time (as it is in Bitcoin), the effect is greatly reduced.

Jute all but eliminates orphans, eliminating this centralizing effect along with them. One of the chief reasons that a high block time is necessary gets eliminated, and a nontrivial mining centralization pressure is eliminated as well.

##### Selfish Mining

Miners with sufficient hashrate or sufficient network superiority are able to perform strategic block witholding and block propagation strategies such that their profitability increases by them increasing the relative orphan rates of their competitors. Because Jute mostly eliminates orphans,  orphan-based selfish mining is no longer an issue.

###### Papers and Links Related to Selfish Mining:
[Majority is Not Enough: Bitcoin Mining is Vulnerable](https://www.cs.cornell.edu/~ie53/publications/btcProcFC.pdf)  
[Optimal Selfish Mining Strategies in Bitcoin](http://fc16.ifca.ai/preproceedings/30_Sapirshtein.pdf)  
[Stubborn Mining: Generalizing Selfish Mining and Combining with an Eclipse Attack](https://eprint.iacr.org/2015/796.pdf)  
[Block Publication Incentives for Miners](https://petertodd.org/2016/block-publication-incentives-for-miners)  

##### Poorly Utilized Network

The bitcoin network spends most of it's life idle. Every 10 minutes, a 1MB block propagates, and then the majority of the nodes are quiet again for approximately the whole next 10 minutes. Miners especially must have low latency setups, which means mining in certain parts of the world becomes infeasible (or at least less profitable), and it means that the amount of initial capital needed for mining is higher. Large miners can more easily amortize initial capital costs meaning that larger, more centralized miners experience higher per-hashrate profitability.

Things like [compact blocks](https://bitcoincore.org/en/2016/06/07/compact-blocks-faq/), [the relay network](http://bitcoinrelaynetwork.org/), [FIBRE](http://bluematt.bitcoin.ninja/2016/07/07/relay-networks/), [weak blocks](http://popeller.io/index.php/2016/01/19/weak-blocks-the-good-and-the-bad/), and other fast-relay schemes all help to improve network utilization, however most of these strategies rely on pre-forwarding transactions and do not handle adversarial blocks very well.

With jute, miners can remain competitive without seeing their competitors blocks immediately. The reduced emphasis on network latency means that the block time can be lowered, and optimizations can focus more heavily on throughput. The lowered blocktime also means a more consistent load applied to the network, which itself is an optimization for throughput.

###### Papers and Links Related to Bitcoin's Network:
[On Scaling Decentralized Blockchains (A Position Paper)](http://fc16.ifca.ai/bitcoin/papers/CDE+16.pdf)

##### High Variance Mining Rewards

The long block times in Bitcoin means that smaller mining operations may only be able to find a few blocks per month or year, following a poisson distribution. This means that there is very high variance in month-to-month income for small solo miners, which makes solo mining substantially less practical without many tens of thousands of dollars of mining hardware. Bitcoin addresses this primarily thorugh the use of mining pools, which is a centralization pressure.

The lower block time of Jute means more payouts to miners, greatly reducing the minimum-hashpower requirement for practical solo mining, reducing a significant miner centralization pressure in the Bitcoin ecosystem.

##### High Variance Confirmation Times

The time between blocks follows an exponential distribution. It's common for blocks to be more than 30 minutes apart. This hurts the user experience, as the difference between 10 minutes and 30 minutes can be significant in daily life. By lowering the block time, Jute reduces the variance in confirmation times.

It should be noted that a single confirmation in Jute is much weaker than a single confirmation in Bitcoin. Instead, one must wait until a particular ordering is being endorsed by a majority of the hashrate, which can take a couple of network propagation cycles. Confirmation times with jute are still in the ballpark of 5-10 minutes, but with much lower variance.

## Security Requirements of Blockchains

The key innovation of a blockchain is immutability. And Bitcoin assumes immutability on its transactions given enough confirmations and absent an attacker with at least 51% hashrate. Aside from assuming that nobody has reached 51% hashrate, Bitcoin assumes that all miners are acting selfishly and are pursuing profits, willing to engage in unfriendly or adversarial behavior if it is profitable to do so.

And what's mor

## Jute Introduction

Jute allows blocks to have multiple parents, effectively converting the blockchain into a directed acyclic graph, or DAG. Consensus is heavily dependent on having a precise ordering for transactions, and security is heavily dependent on knowing that a highly confirmed transaction cannot be re-ordered or re-organized unless an impractical amount of additional, adversarial work is applied to the chain.

Jute is an algorithm for taking DAGs and arriving at a total ordering based on the amount of work confirming each potential ordering. Once the majority of the hashrate has agreed upon a particular ordering, that ordering alone receives votes, which means that a minority hashrate is unable to cause reordering or reorganization without being exceptionally lucky.

Jute ignores the validity of transactions in the immediate term, allowing any block with valid proof-of-work to be accepted into the chain. Ignoring transaction validity is important for eliminating orphans, because multiple miners finding blocks at the same time should not be heavily penalized if there is some transaction overlap in those blocks - it will be unavoidable for high latency miners. This does not need to prevent SPV however. Instead of making SPV commitments to the transcations in the blocks they are mining (where they are not certain that the transactions will actually be valid due to potential double spends), miners can commit to block ranges in recent history with a high enough number of confirmations that there is little to no risk of the commitment being made invalid by a double spend.

## The Jute Ordering Algorithm

The genesis block must be the first block in the blockchain. All blocks must point to either the genesis block or descendents of the genesis block. Blocks are allowed to have multiple parents. Blocks are not allowed to have parents which are ancestors of other parents (this makes some of the code cleaner).

![Example DAG](http://i.imgur.com/IHMN24h.jpg)

The block that has no children is called 'the tip'. Though there may be multiple tips at the same time, only one tip is considered at a time. For the purposes of consensus, the tip with the most work in its ancestry is considered the canonical tip, defining the ordering of blocks that is used for consensus.

A direct child-parent connection is called an edge. Each block except for the tip block will have edges to one or more children. Blocks vote for edges. The block will execute a sort on its children based on the votes of previous blocks, and then cast votes based on the sort results. Sort using the following algorithm:
``` 
1. Set 'parent' to the genesis block
2. Append 'parent' to the sorting.
3. Look at the vote counts for all edges from 'parent' to its children and select the edge with the most votes.
    3a. If there are no edges, return.
    3b. If multiple edges have the same winning number of votes, use the hash of the tip block to seed a deterministic RNG that selects between them.
4. Set 'child' to the child block of the winning edge.
5. While 'child' has ancestors that are not in the sorting:
    5a. In this loop only, view the DAG, tip, and vote counts as they were when the child was the tip block.
    5b. Look at the vote counts for all edges from blocks in the sorting to blocks not in the sorting and select the edge with the most votes.
        5b i.  If multiple edges have the same winning number of votes, select the edge with the most ancestors.
        5b ii. If multiple edges have the same winning number of ancestors, use the hash of the tip block to seed a deterministic RNG that selects between them.
    5c. Recurse to line 2, setting 'parent' equal to the child of the winning edge. When the recursion returns, iterate to line 5.
6. Set 'parent' to 'child' and iterate to line 2.
```
The final sorting will have an ordering of blocks, some of which are pairings which match real edges in the DAG. For all real edges represented in the sorting, cast one vote. And that's it! Below is a set of DAGs labeled to have the post-sorting vote counts and the blocks in sorted order:

![Five Examples](http://i.imgur.com/HFy70s8.jpg)

![Two Bonus Examples](http://i.imgur.com/uv1TqCB.jpg)

## Intuition Around the Security of Jute

Each block that is mined will only confirm one ordering of the chain, and once that confirmation has been applied, the confirmation cannot be altered by future merges into the chain. Once more than half of the hashrate is in agreement on a particular ordering, that ordering will be buried under lots of work and will be exponentially difficult for a minority hashrate adversary to undo.

What remains then is determining how the network is able to reach agreement in normal circumstances, and determining what an adversary can do to delay agreement on a block ordering. In Bitcoin, you can get situations where exactly half of the network is working on one chain, and exactly half of the network is working on another chain of equal weight. This almost always resolves within a few minutes, as Bitcoin mining has high variance. One side will get lucky and the other will not, causing one chain to get ahead and allowing the honest hashrate to come into agreement on the longest chain. This is possible in particular because the block time is much higher than the network propagation time. A minority hashrate adversary is not able to interfere very well, because interference requires finding blocks fast enough to keep the hashrate balanced between two chains. When the block time is so high, this just isn't feasible.

For Jute, with high block times the same high variance applies. As the block rate goes below the network propagation time, the amount of variance on the network reduces and blocks become a more continuous, predicatable stream. This may inhibit convergence, however it seems that obtaining balance between two halves of the hashrate would be extremely difficult even in the face of low-variance block production. An adversary certainly has a lot more flexibility with Jute, but maintaining balance would still require a lot of network superiority. An attacker could potentially save up blocks to try and maintain balance between two competing versions of the DAG sorting, which also helps, but as soon as one side has a moderate majority things converge very quickly. And, as discussed before, once convergence occurs it's exponentially difficult to reverse.

Convergence is pretty easy to measure, and therefore confirmations would not count the number of blocks, but instead would report a measurement of the convergence. In the example of 'Bonus 2', we can see that there is a very low convergence on the idea that 'B' is going to be the winning block vs. 'E' or 'C'. This is because the DAG is very wide - out of 13 blocks only 5 have voted for the 'AB' edge, and potentially there are other tips out there which haven't propagated yet which further widen the DAG and reduce or eliminate the lead held by edge 'AB'. Convergence will start to appear more strongly once the DAG is more substantially longer than it is wide. In example 5, convergence on 'AB' is quite high, as the DAG is not very wide and it's likely that all future blocks are going to continue confirming 'AB' over 'AD' or any other alternatives that may exist.

The width of a dag is closely related to the ratio of the block time vs. the network propagation time. Bitcoin as a dag would be very narrow, with less than 5% of blocks having siblings (likely less than 2%). On the other hand, if the block time were set to 1 second, the average width of the DAG would probably be 3+ blocks.

#### Incentives to Merge Blocks

Only the longest chain is considered when determining consensus. Though jute is designed to not have any orphans, this only works if people are merging eachother. If people are not merging eachother, those in the minority chain will lose out. They can gain an advantage over the majority chain by merging in the majority chain.

There's also a tiebreaking rule where in the event of a tie, the edge confirming the most blocks is selected as the winner. By merging as many blocks as possible, you advantage yourself and your children, which increases your liklihood of appearing at a low height instead of a high height.

Finally, an adversary has a much easier time executing attacks against the network if miners are not merging eachother's blocks.

The data consumed by merges should not be counted towards block size. This eliminates any counter-incentives to performing merges, and provides a stronger guarantee that the network as a whole will engage in healthy behavior. Though the incentives to perform merges are not very strong, the counter-incentives are even weaker, meaning that merging should in general happen as much as possible.

Blocks are capped to 20 parents to prevent DoS attacks trying to abuse the costless parents feature.

#### Handling Merges from Lower Difficulty Time Periods

Blocks can be created that choose any parent, and they are candidates to be merged into the main chain. This means that blocks could potentially choose parents from a bery early time period. Such blocks are not at risk of disrupting consensus, as they will not have enough confirmations on their edges to influence any changes, however it would not be fair to give these blocks the full block reward, as they were not as difficult to create. This also means that an adversary could potentially create thousands of blocks with ease due to the lower difficulty.

A rule probably needs to be added to jute enforcing that blocks are at least somewhat recent, and ensuring that payouts are fair - lower difficulty blocks should receive lower rewards.

#### Costless Attacks

Because there are no orphans, there's little to no cost for an attacker to attempt and fail to execute a double-spend of a confirmed transaction on the network. Such a double spend is not going to affect the profitability of other miners, and so is not a huge concern in terms of mining centralization. On the other hand, users like to know that their transactions are safe.

An attacker cannot choose when they get lucky. A 40% hashrate attacker is going to be able to double spend a 3-confirmation transaction only 6.5% of the time. Block times in Jute are very low, so transactions perceived as high risk would not be too inconvenienced to wait for a few extra confirmations. Network latency does play a role - you cannot be certain at all that a transaction is confirmed until the network has had at least a few full propagation cycles to reveal any potential double spends. Even without adversarial miners you need to wait at least that long. Sub-second block times will never mean sub-second confirmation times.

Things get more annoying when miners start making SPV commitments to old blocks. For SPV to be valuable, a block has to be excluded from the chain if it commits to an invalid transaction. If an attacker can reorg deep enough to invalid SPV commitments by other miners, the attacker will succeed in creating orphans. If the reorg attempt is not going well, it can be aborted early at no cost to the attacker.

As long as the SPV commitments are sufficiently far back, there is little risk of an attacker pulling this off successfully. For lower block times, sufficiently far back may only take 5-10 minutes. But if the block time needs be set higher than 5 seconds, sufficiently far back may be 30+ minutes, which reduces the utility of SPV.

#### Transaction Fee Complications

Absent transcation fees, jute is perfectly fair to miners, independent of both hashrate and network connecectivity (so long as latency is under a few minutes, which isn't that much of an ask).

Bitcoin is increasingly going to be dominated by transaction fee rewards in the future. Unfortunately, achieving perfect fairness in the face of transaction fees is substantially harder. The ideal solution would be a mining fee pool that collects all of the transaction fees and pays out evenly to miners. Unfortuantely, it's pretty easy for an adversary to exploit a mining fee pool.

When the block time is sufficiently above the network propagation time, most of the problems don't exist. A high hashrate attacker has relatively little ability to manipulate the ordering of blocks, and things will be largely fair between high latency, low hashrate miners and low latency, high hashrate miners.

As the block time comes down though, the low latency miners increasingly have the ability to see which fees have been mined, they have the ability to exploit network advantages to try and steal fees, and high hashrate miners have increased ability to manipulate the exact ordering of transactions. Not a problem when miner fees are a negligable part of your income, but reality is against us.

When the block time is substantially below the network propagation time, a low-latency high-hashrate miner has a very strong ability to choose which transactions come before which other transactions in the short term, because they can dump a huge portion of the hashrate into confirming a single set of transactions when the rest of the hashrate is distributed over a much wider graph.

If blocks are not full, the larger miners are going to be able to more or less monopolize the transaction fees by mining strategically. If they are full, larger miners will still be able to have a significant amount of control over who is getting fees, and will be able to fairly trivially steal from a small percent of smaller miners. If they use that strategically to target one miner at a time, they can iteratively put their smaller competition out of business.

----

It's not really that gloomy though. There are a handful of strategies we can emply to mitigate this. The most conceptually simple idea that I have right now involves doing post-processing at the consensus layer. If two miners mine the same transaction, you can split the fee between them. If two miners mine conflicting transactions, you can accept the transaction with the larger fee and split it between the miners. If there is not a large enough transaction log, you can perform 'fee stretching' by stipulating that at most 1/N of the fee can be collected per block, and that to get the whole fee it needs to be mined in N different blocks. This gives smaller miners a better chance.

I have a bunch of other ideas as well.

### A Summary of Potential Problems

+ As block times decrease, the largest legal block size also decreases. Below some threshold that seems to be between 5kb and 25kb (based on some... erm... interviews) complex constructions like CoinJoin and MimbleWimble become infeasible.
+ As block times decrease, variance also decreases, which may make it more difficult for the network to begin to converge on the result of a transaction. My vague intuition says that this really isn't going to be much of an issue, but I don't currently have the mathematics needed to back up that intuition. It needs to be explored more.
+ If the block time is not low enough, doing SPV securely is going to mean that commitments will need to be made long after the block containing the transaction is mined. Potentially 30 minutes or more. If the block time IS low enough, SPV may only need 5 minutes of delay, which is actually an improvement over Bitcoin. If you can assume that the strongest attacker has less than 33% hashrate, 5 minutes should be practical. If you want to defend also against 45% hashrate attackers and really high latency networks, you'll probably want 1 or 2 second block times combined with 20+ minutes of wait time (these numbers are based on some simulations I ran - code will be provided in a later draft).
+ Perhaps the worst problem relates to transaction fee fairness. Bitcoin is increasingly going to be dominated by transaction fees, so it's not something that can be overlooked. There are lots of ideas for managing the fee unfairness, but most of them introduce a fair amount of complexity or make other compromises. I have genuine hope that a simple solution exists, but at this time we do not seem to be very close to finding it.


### Practical Jute Today

+ Deploy the jute consensus algorithm
+ Set the block time to 5 seconds
+ Set the block size to 25kb
+ Set SPV to commit to block ranges after they have 600+ confirmations (50 minutes)
+ Don't allow merges of blocks if their most recent parent already has 2000+ confirmations
+ Ignore miner fee complications because they shouldn't come into play too much at 5 second confirmation times

Realistically, by the time this was implemented we'd have worked through a lot more of the issues. I have casually handwaved over things like the need to update the difficulty adjustment algorithm, the timestamp algorithm, and probably a whole host of other protocol-level things that this interferes with. Largely though I don't think those are hard things to adjust, you just need to be careful and methodical.

#### Acknowledgements

Thanks to Andrew Poelstra for review + corrections.  
Thanks to Bob McElrath for active research on DAG based consensus, which has influenced this document.