Jute Ordering
=============

This is an intuitive but unoptimized implementation of the Jute ordering
algorithm. Jute is an algorithm for ordering proof-of-work blocks from an
arbitrary DAG into a strict order, such that security against 51% attacks can
be achieved, and mining fairness can be achieved.

graph.go contains definitions for the basic data structures, addnode.go
contains the code for adding a node to the DAG (including the required voting
step), and ordering.go contains code for taking an existing graph and deriving
the secure ordering according to Jute.

main.go does not contain any core code, however it does contain code that can
create orderings and produce SageMath code for generating visualizations of the
graphs.

jute\_test.go contains most of the testing that has been performed. At the
moment, it is somewhat sparse, and therefore the Jute code may contain bugs.
