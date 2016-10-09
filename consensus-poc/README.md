Jute Consensus Proof-of-Concept
===============================

This contains highly unoptimized, attempting-to-be-intuitive code which takes a
block DAG and sorts it according to the jute consensus algorithm. Example
graphs are in main.go, with the output of running the consensus-poc binary
being graphing code that you can feed into SageMath to get a visualization of
the resulting graph. The title of the SageMath graph is the sorting order of
the nodes, and the edges have weights to indicate how many votes there are for
each edge.

Included in this repo are two 'plots' directories, containing sample graphs
that can be produced by feeding output into SageMath. The 'plots' folder
contains plots with the low block time code disabled, and the 'plotsLBT' folder
contains the sample plots with the low block time code enabled.

The code in this directory has not been tested very thoroughly. All of the
sample graphs are correct, however there may be bugs when trying to build
larger or more complex graphs.
