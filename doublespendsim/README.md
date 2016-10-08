Double Spend Simulator
======================

The double spend simulator runs a simulation where an attacker tries to commit
a double-spend to determine the probability of success for the attacker. The
attacker is given zero latency. The hashrate of the attacker, the latency of
the honest miners, the block time, and the number of confirmations can all be
configured. The result is the probability that the attacker's double spend will
be successful.

This is particularly relevant for jute, as an attacker can attempt a double
spend almost for free. Jute features a high block time, so knowing the number
of confirmations required for reasonable security under various network
conditions is very useful.

Some results of note:

33% attacker, 50 confirmations (5 minutes), 0.06% success rate.
	Iterations:                          100000
	Attacker Hashrate:                   33%
	Block Time:                          18000 (milliseconds)
	Honest Miner Block Propagation Time: 6000 (milliseconds)
	Confirmations:                       50
	Threads:                             2
	Difficulty:                          10
	Nonce Starting Point:                5004996000323410951
	Attacker Wins: 56 (0.06%)

40% attacker, 150 confirmations (15 minutes), 0.04% success rate.
	Iterations:                          100000
	Attacker Hashrate:                   40%
	Block Time:                          18000 (milliseconds)
	Honest Miner Block Propagation Time: 6000 (milliseconds)
	Confirmations:                       150
	Threads:                             2
	Difficulty:                          10
	Nonce Starting Point:                5301066859990594536
	Attacker Wins: 40 (0.040000%)

45% attacker, 500 confirmations (50 minutes), 0.08% success rate.
	Iterations:                          100000
	Attacker Hashrate:                   45%
	Block Time:                          18000 (milliseconds)
	Honest Miner Block Propagation Time: 6000 (milliseconds)
	Confirmations:                       500
	Threads:                             2
	Difficulty:                          10
	Nonce Starting Point:                4905061595660334175
	Attacker Wins: 75 (0.075000%)

Summarized, getting below a 0.1% success rate requires 50 confirmations against
a 33% hashrate attacker, 150 confirmations against a 40% hashrate attacker, and
500 confirmations against a 45% hashrate attacker.

----------------------

33% attacker, 200 confirmations (20 minutes), **1 million iterations**, 0% success rate.
	Iterations:                          1000000
	Attacker Hashrate:                   33%
	Block Time:                          18000 (milliseconds)
	Honest Miner Block Propagation Time: 6000 (milliseconds)
	Confirmations:                       200
	Threads:                             2
	Difficulty:                          10
	Nonce Starting Point:                5890706053630306446
	Attacker Wins: 0 (0%)

Summarized, its extremely unlikely that a 33% hashrate attacker can undo more
than 200 confirmations. In Bitcoin, reorgs of more than 20 minutes of history
are rare, but not surprising when they happen, and they usually happen on
accident instead of being intentionally performed by an attacker.

----------------------

40% attacker, 300 confirmations (30 minutes), **120 second propagation latency**, 0.001% success rate.
	Iterations:                          100000
	Attacker Hashrate:                   40%
	Block Time:                          **120000** (milliseconds)
	Honest Miner Block Propagation Time: 6000 (milliseconds)
	Confirmations:                       300
	Threads:                             2
	Difficulty:                          10
	Nonce Starting Point:                7693741952947876035
	Attacker Wins: 1 (0.001%)

Summarized, a 40% hashrate attacker is highly unlikely to undo more than 300
confirmations, even if there is extreme latency between the honest miners.

----------------------

45% attacker, 1200 confirmations (**10 minutes**), 0.002% success rate.
	Iterations:                          100000
	Attacker Hashrate:                   45%
	Block Time:                          18000 (milliseconds)
	Honest Miner Block Propagation Time: **500** (milliseconds)
	Confirmations:                       1200
	Threads:                             2
	Difficulty:                          10
	Nonce Starting Point:                8334072111784665621
	Attacker Wins: 2 (0.002%)

Summarized, bringing the block time very low (500 milliseconds) enables
confirmations to happen very quickly. While the extreme difference between the
block time and the propagation time favors the attacker, the sheer number of
confirmations means that the attacker is less successful. A double spend would
be extremely rare after 1200 confirmations, even if the network is configured
such that those confirmations only take 10 mintues.

It should be noted that this only represents security with regards to
consumers. When the block propagation time is this low, the high hashrate
attacker is able to execute non-double-spend related attacks which greatly
impact the profitability of honest miners. Bringing the block time down into
the hundreds of milliseconds has a lot of potential advantages, however more
problems need to be worked out before this is feasible.
