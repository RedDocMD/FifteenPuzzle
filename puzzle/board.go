package puzzle

import (
	"errors"
	"fmt"
)

// Board represents the n puzzle board at any given time
type Board struct {
	tiles         [][]int8
	size          int8
	parent        *Board
	depth         int8
	heuristic     int
	heuristicType int8
}

// Defines possible actions on this state to generate the next
const (
	ShiftLeft  = iota
	ShiftRight = iota
	ShiftDown  = iota
	ShiftUp    = iota
	MaxActions = iota // Must be last element
)

// Defines the possible heuristic types
const (
	WeightedManhattan = iota
	SummedManhattan   = iota
	MaxManahttan      = iota
	InversionDistance = iota
	CombinedManhattan = iota
	MaxHeuristicTypes = iota // Must be last element
)

// NewBoard returns a Board pointer for the given tile config
// Returns nil if tiles is an invalid board config
func NewBoard(tiles [][]int8, size int8, heuristicType int8) *Board {
	f := Board{tiles, size, nil, 0, -1, heuristicType}
	return &f
}

// Depth returns the depth of the node (0 for the start node)
func (board *Board) Depth() int {
	return int(board.depth)
}

// NextBoard generates the next Board corresponding to the given action or nil of not possible
func (board *Board) NextBoard(action int8) *Board {
	newTiles := make([][]int8, board.size)
	for i := range newTiles {
		newTiles[i] = make([]int8, board.size)
		copy(newTiles[i], board.tiles[i])
	}
	newBoard := NewBoard(newTiles, board.size, board.heuristicType)
	newBoard.parent = board
	newBoard.depth = board.depth + 1

	i, j, _ := board.findZero()
	switch action {
	case ShiftUp:
		if i == 0 {
			newBoard = nil
		} else {
			newBoard.tiles[i][j] = newBoard.tiles[i-1][j]
			newBoard.tiles[i-1][j] = 0
		}
	case ShiftDown:
		if i == board.size-1 {
			newBoard = nil
		} else {
			newBoard.tiles[i][j] = newBoard.tiles[i+1][j]
			newBoard.tiles[i+1][j] = 0
		}
	case ShiftLeft:
		if j == 0 {
			newBoard = nil
		} else {
			newBoard.tiles[i][j] = newBoard.tiles[i][j-1]
			newBoard.tiles[i][j-1] = 0
		}
	case ShiftRight:
		if j == board.size-1 {
			newBoard = nil
		} else {
			newBoard.tiles[i][j] = newBoard.tiles[i][j+1]
			newBoard.tiles[i][j+1] = 0
		}
	}
	if newBoard != nil {
		switch newBoard.heuristicType {
		case InversionDistance:
			newBoard.heuristic = board.inversionDistanceFromMove(action, i, j)
		case SummedManhattan:
			newBoard.heuristic = newBoard.summedManhattanFromMove(board.heuristic, action, i, j)
		}
	}
	return newBoard
}

func (board *Board) findZero() (int8, int8, error) {
	var i, j int8
	for i = 0; i < board.size; i++ {
		for j = 0; j < board.size; j++ {
			if board.tiles[i][j] == 0 {
				return i, j, nil
			}
		}
	}
	return int8(0), int8(0), errors.New("Invalid board")
}

// Solved checks if the board is in the position which is the solved state
func (board *Board) Solved() bool {
	size := board.size
	for i := int8(0); i < size*size-1; i++ {
		if board.tiles[i/size][i%size] != i+1 {
			return false
		}
	}
	return board.tiles[size-1][size-1] == 0
}

// String returns string representation of Board
func (board Board) String() string {
	var str string = ""
	var i, j int8
	for i = 0; i < board.size; i++ {
		for j = 0; j < board.size; j++ {
			if board.tiles[i][j] != 0 {
				str += fmt.Sprintf("%3d", board.tiles[i][j])
			} else {
				str += "   "
			}
		}
		str += "\n"
	}
	return str
}

// PrintPath prints the path from the first node to this board
func (board *Board) PrintPath() {
	if board == nil {
		fmt.Println("Cannot print path from nil board")
	} else {
		fmt.Println("There are", board.depth, "transitions")
		count := int(board.depth + 1)
		boards := make([]*Board, 0)
		pt := board
		for pt != nil {
			boards = append(boards, pt)
			pt = pt.parent
		}
		for i := count - 1; i >= 0; i-- {
			fmt.Println(boards[i])
		}
	}
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// Heuristic returns the Manhattan distance heuristic for the given node
func (board *Board) Heuristic() int {
	if board.heuristic == -1 {
		var val int
		switch board.heuristicType {
		case WeightedManhattan:
			val = board.weightedManhattan()
		case SummedManhattan:
			val = board.summedManhattan()
		case MaxManahttan:
			val = board.maxManhattan()
		case InversionDistance:
			val = board.inversionDistance()
		case CombinedManhattan:
			val = board.combinedManhattan()
		}
		board.heuristic = val
		return val
	} else {
		return board.heuristic
	}
}

func (board *Board) weightedManhattan() int {
	sum := 0
	weight := int(board.size)
	for i := 0; i < int(board.size); i++ {
		for j := 0; j < int(board.size); j++ {
			val := int(board.tiles[i][j]) + 1
			ii := val / int(board.size)
			jj := val % int(board.size)
			sum += (abs(ii-i) + abs(jj-j)) * weight
		}
		weight--
	}
	return sum
}

func (board *Board) summedManhattan() int {
	sum := 0
	for i := 0; i < int(board.size); i++ {
		for j := 0; j < int(board.size); j++ {
			val := int(board.tiles[i][j]) + 1
			ii := val / int(board.size)
			jj := val % int(board.size)
			sum += abs(ii-i) + abs(jj-j)
		}
	}
	return sum
}

func (board *Board) summedManhattanFromMove(old int, action int8, i int8, j int8) int {
	// Call only on the NEW board
	heuristic := old
	val := int(board.tiles[i][j])
	ii := val / int(board.size)
	jj := val % int(board.size)
	heuristic += abs(ii-int(i)) + abs(jj-int(j))

	switch action {
	case ShiftDown:
		heuristic -= abs(ii-int(i+1)) + abs(jj-int(j))
	case ShiftUp:
		heuristic -= abs(ii-int(i-1)) + abs(jj-int(j))
	case ShiftLeft:
		heuristic -= abs(ii-int(i)) + abs(jj-int(j-1))
	case ShiftRight:
		heuristic -= abs(ii-int(i)) + abs(jj-int(j+1))
	}

	return heuristic
}

func (board *Board) linearConflict() int {
	lc := 0
	for i := 0; i < int(board.size); i++ {
		for j := 0; j < int(board.size)-1; j++ {
			val1 := int(board.tiles[i][j]) + 1
			val2 := int(board.tiles[i][j+1]) + 1
			if1 := val1 / int(board.size)
			if2 := val2 / int(board.size)
			jf1 := val1 % int(board.size)
			jf2 := val2 % int(board.size)
			if if1 == if2 {
				if jf1-jf2 == 1 {
					lc++
				}
			}
		}
	}
	for j := 0; j < int(board.size); j++ {
		for i := 0; i < int(board.size)-1; i++ {
			val1 := int(board.tiles[i][j]) + 1
			val2 := int(board.tiles[i+1][j]) + 1
			if1 := val1 / int(board.size)
			if2 := val2 / int(board.size)
			jf1 := val1 % int(board.size)
			jf2 := val2 % int(board.size)
			if jf1 == jf2 {
				if if1-if2 == 1 {
					lc++
				}
			}
		}
	}
	return 2 * lc
}

func (board *Board) combinedManhattan() int {
	manhattan := board.summedManhattan()
	linearConflict := board.linearConflict()
	return manhattan + linearConflict
}

func (board *Board) maxManhattan() int {
	max := 0
	for i := 0; i < int(board.size); i++ {
		for j := 0; j < int(board.size); j++ {
			val := int(board.tiles[i][j]) + 1
			ii := val / int(board.size)
			jj := val % int(board.size)
			diff := abs(ii-i) + abs(jj-j)
			if diff > max {
				max = diff
			}
		}
	}
	return max
}

func (board *Board) inversionDistance() int {
	unpacked := make([]int, board.size*board.size)
	idx := 0
	for i := 0; i < int(board.size); i++ {
		for j := 0; j < int(board.size); j++ {
			unpacked[idx] = int(board.tiles[i][j])
			idx++
		}
	}

	inv := 0
	for i := 0; i < int(board.size*board.size); i++ {
		if unpacked[i] != 0 {
			for j := 0; j < i; j++ {
				if unpacked[i] < unpacked[j] {
					inv++
				}
			}
		}
	}
	vertical := inv/3 + 1

	idx = 0
	for i := 0; i < int(board.size); i++ {
		for j := 0; j < int(board.size); j++ {
			unpacked[idx] = j*int(board.size) + i
			idx++
		}
	}

	inv = 0
	for i := 0; i < int(board.size); i++ {
		for j := 0; j < int(board.size); j++ {
			val := int(board.tiles[i][j]) - 1
			if val != -1 {
				idx = 0
				for k := range unpacked {
					if unpacked[k] == val {
						idx = k
						break
					}
				}
				inv += abs(idx - (j*int(board.size) + i))
			}
		}
	}
	horizontal := inv/3 + 1
	// horizontal := 0

	return vertical + horizontal
}

var lookup [16]int8 = [16]int8{1, 5, 9, 13, 2, 6, 10, 14, 3, 7, 11, 15, 4, 8, 12, 0}

func findInLookup(num int8) int8 {
	for i := range lookup {
		if lookup[i] == num {
			return int8(i)
		}
	}
	return -1
}

func (board *Board) inversionDistanceFromMove(action int8, i int8, j int8) int {
	heuristic := board.heuristic
	size := board.size
	count := 0
	if heuristic == -1 {
		panic(errors.New("Call board.Heuristic() at least once"))
	}
	switch action {
	case ShiftUp:
		idx := i*size + j - 3
		el := board.tiles[i-1][j]
		if el > board.tiles[idx/size][idx%size] {
			count++
		} else {
			count--
		}
		idx++
		if el > board.tiles[idx/size][idx%size] {
			count++
		} else {
			count--
		}
		idx++
		if el > board.tiles[idx/size][idx%size] {
			count++
		} else {
			count--
		}
	case ShiftDown:
		idx := i*size + j + 3
		el := board.tiles[i+1][j]
		if el < board.tiles[idx/size][idx%size] {
			count++
		} else {
			count--
		}
		idx--
		if el < board.tiles[idx/size][idx%size] {
			count++
		} else {
			count--
		}
		idx--
		if el < board.tiles[idx/size][idx%size] {
			count++
		} else {
			count--
		}
	case ShiftLeft:
		idx := i + j*size - 3
		el := board.tiles[i][j-1]
		elpos := findInLookup(el)
		if elpos > findInLookup(board.tiles[idx%size][idx/size]) {
			count++
		} else {
			count--
		}
		idx++
		if el > findInLookup(board.tiles[idx%size][idx/size]) {
			count++
		} else {
			count--
		}
		idx++
		if elpos > findInLookup(board.tiles[idx%size][idx/size]) {
			count++
		} else {
			count--
		}
	case ShiftRight:
		idx := i + j*size + 3
		el := board.tiles[i][j+1]
		elpos := findInLookup(el)
		if elpos < findInLookup(board.tiles[idx%size][idx/size]) {
			count++
		} else {
			count--
		}
		idx--
		if el < findInLookup(board.tiles[idx%size][idx/size]) {
			count++
		} else {
			count--
		}
		idx--
		if elpos < findInLookup(board.tiles[idx%size][idx/size]) {
			count++
		} else {
			count--
		}
	}
	if count < 0 {
		heuristic += abs(count)/3 + 1
	} else {
		heuristic -= abs(count) / 3
	}
	return heuristic
}

const int64Max = int64(9223372036854775807)

// Hash implements a hash function for Board
func (board *Board) Hash() int64 {
	hash := int64(0)
	for i := 0; i < int(board.size); i++ {
		for j := 0; j < int(board.size); j++ {
			hash += (16*(hash%int64Max) + int64(board.tiles[i][j])) % int64Max
			if hash < 0 {
				hash += int64Max
			}
		}
	}
	return hash
}
