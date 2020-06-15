# 15 puzzle

This is an application written in Go which attempts to solve the 15 puzzle with several state-space-search methods

## Implementations

The following algorithms will be used to find a solution:

1. DFS
2. IDFS (Iterative-deepening DFS)
3. A\*
4. IDA\* (Iterative Deepening A\*)

The heuristics implemented were:

1. Max Manhattan Distance (The maximum among the Manhattan distances of all the tiles)
2. Summed Manhattan Distance (The sum of the Manhattan distances of all the tiles)
3. Weighted Manhattan Distance (The Manhattan distances were weighted by the row)
4. Manhattan Distance + Linear Conflict (In short *Combined Manhattan*)
5. Inversion Distance (A heuristic based on the ordering of the tiles)

## Performance

Measurements were made on a laptop running Arch Linux (5.7.2-arch1-1), with an Intel i5-8300H cpu (8 virtual cores) and 8GB of RAM.

### Uninformed algorithms

#### Vanilla DFS

This just fills up RAM. Even with a modest branching factor of at most 4.

#### Iterative-deepening DFS

Memory issues are solved. But it is still very slow and even after 75 minutes of the program running, only a depth of 18 was reached.

### Informed algorithms

#### A\*

This algorithm fills up memory slower than DFS. Nevertheless, it runs out of memory except in the very simple test case.

#### IDA\*

This algorithm fares better than the others in terms of memory usage and speed. However the performance heavily depends on the heuristic used.

Based on measurements made on the input from file **third_input**, the following order was obtained:

    Combined Manhattan (20.746s) < Summed Manhattan (1.250s) < Inversion Distance (0.011s)

All of them found out a solution of 15 moves each.

The performance of *Combined Manhattan* is not fully reflected in the above number. While the other two heuristic values are calculated in constant time from the heuristic value of the previous node, I haven't found a way to do so for the Linear Conflict. Hence, the entire heuristic is recalculated from scratch for each node.

## Problems

Even with the best heuristic (Inversion distance) and IDA*, not all legal initial positions are being solved. The harder initial states seem to take forever to solve.

