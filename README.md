# 15 puzzle

This is an application written in Go which attempts to solve the 15 puzzle with several state-space-search methods

The following algorithms will be used to find a solution (*not* necessarily optimal):

1. Iterative-deepening DFS
2. Bi-directional search
3. Vanilla DFS (Just for kicks :smiley:)

The following algorithms will be used to find an optimal solution:

## Performance

For the non-optimal algorithms:

1. **Vanilla DFS**:  This just fills up RAM. Even with a modest branching factor of at most 4.
