package puzzle

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
