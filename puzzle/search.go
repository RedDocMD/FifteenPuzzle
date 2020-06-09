package puzzle

import (
	"fmt"
)

// DepthFirstSearch performs vanilla DFS
func DepthFirstSearch(start *Board) *Board {
	var open []*Board
	closed := make(map[*Board]bool)
	open = append(open, start)
	for len(open) > 0 {
		n := len(open) - 1
		node := open[n]
		open = open[:n]
		if node.Solved() {
			return node
		}
		if closed[node] {
			continue
		}
		var next *Board

		next = node.NextBoard(ShiftDown)
		if next != nil {
			open = append(open, next)
		}

		next = node.NextBoard(ShiftUp)
		if next != nil {
			open = append(open, next)
		}

		next = node.NextBoard(ShiftLeft)
		if next != nil {
			open = append(open, next)
		}

		next = node.NextBoard(ShiftRight)
		if next != nil {
			open = append(open, next)
		}

		closed[node] = true
	}
	return nil
}

// Depth limited DFS return flags
const (
	SUCCESS = iota
	FAILURE = iota
	CUTOFF  = iota
)

// IterativeDeepeningSearch performs iterative deepening DFS
func IterativeDeepeningSearch(start *Board) *Board {
	for limit := 1; ; limit++ {
		fmt.Println("Current depth limit is", limit)
		result, goal := recursiveDLS(start, limit)
		switch result {
		case SUCCESS:
			return goal
		case FAILURE:
			return nil
		}
	}
}

func recursiveDLS(node *Board, limit int) (int, *Board) {
	if node.Solved() {
		return SUCCESS, node
	}
	if limit == 0 {
		return CUTOFF, nil
	}
	cutoff := false
	actions := actions()
	for i := range actions {
		action := actions[i]
		next := node.NextBoard(action)
		if next != nil {
			result, goal := recursiveDLS(next, limit-1)
			switch result {
			case SUCCESS:
				return result, goal
			case CUTOFF:
				cutoff = true
			}
		}
	}
	if cutoff {
		return CUTOFF, nil
	}
	return FAILURE, nil
}

// IterativeDeepeningAStar performs IDA* search algorithm
func IterativeDeepeningAStar(start *Board) *Board {
	limit := 1
	for {
		fmt.Println("Current f cutoff is", limit)
		result, nextCutoff, goal := recursiveDAstar(start, limit)
		switch result {
		case SUCCESS:
			return goal
		case FAILURE:
			return nil
		case CUTOFF:
			limit = nextCutoff
		}
	}
}

func recursiveDAstar(node *Board, fLimit int) (int, int, *Board) {
	if node.Solved() {
		return SUCCESS, fLimit, node
	}
	f := node.Depth() + node.Heuristic()
	if f > fLimit {
		return CUTOFF, f, nil
	}
	cutoff := false
	cutoffLimit := int(1e9)
	actions := actions()
	for i := range actions {
		action := actions[i]
		next := node.NextBoard(action)
		if next != nil {
			result, nextCutoff, goal := recursiveDAstar(next, fLimit)
			switch result {
			case SUCCESS:
				return result, fLimit, goal
			case CUTOFF:
				cutoff = true
				if nextCutoff < cutoffLimit {
					cutoffLimit = nextCutoff
				}
			}
		}
	}
	if cutoff {
		return CUTOFF, cutoffLimit, nil
	}
	return FAILURE, fLimit, nil
}
