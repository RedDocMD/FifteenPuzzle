# 15 puzzle

This is an application written in Go which attempts to solve the 15 puzzle with several state-space-search methods

The following algorithms will be used to find a solution (*not* necessarily optimal):

1. Iterative-deepening DFS
2. Bi-directional search
3. Vanilla DFS (Just for kicks :smiley:)

The following algorithms will be used to find an optimal solution:

## Performance

Measurements were made on a laptop running Ubuntu 20.04 (GNOME), with an Intel i5-8300H cpu (8 virtual cores) and 8GB of RAM.

For the non-optimal algorithms:

1. **Vanilla DFS**:  This just fills up RAM. Even with a modest branching factor of at most 4.
2. **Iterative-deepening DFS**: Memory issues are solved. But it is still very slow and even after 75 minutes of the program running, only a depth of 18 was reached.
