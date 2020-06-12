package puzzle

import (
	"errors"
	"fmt"
)

// Board represents the n puzzle board at any given time
type Board struct {
	tiles       [][]int8
	size        int8
	transitions []int8
	depth       int8
	heuristic   int
}

// Defines possible actions on this state to generate the next
const (
	ShiftLeft  = iota
	ShiftRight = iota
	ShiftDown  = iota
	ShiftUp    = iota
)

// NewBoard returns a Board pointer for the given tile config
// Returns nil if tiles is an invalid board config
func NewBoard(tiles [][]int8, size int8) *Board {
	if !checkTiles(tiles, size) {
		return nil
	}
	transitions := make([]int8, 0)
	f := Board{tiles, size, transitions, 0, -1}
	return &f
}

// Depth returns the depth of the node (0 for the start node)
func (board *Board) Depth() int {
	return int(board.depth)
}

func checkTiles(tiles [][]int8, size int8) bool {
	cnt := make(map[int8]int8)
	var i, j int8
	for i = 0; i < size; i++ {
		for j = 0; j < size; j++ {
			cnt[tiles[i][j]]++
		}
	}
	for i = 0; i < size*size; i++ {
		if cnt[i] != 1 {
			return false
		}
	}
	return true
}

// NextBoard generates the next Board corresponding to the given action or nil of not possible
func (board *Board) NextBoard(action int8) *Board {
	newTiles := make([][]int8, board.size)
	for i := range newTiles {
		newTiles[i] = make([]int8, board.size)
		copy(newTiles[i], board.tiles[i])
	}
	newBoard := NewBoard(newTiles, board.size)
	newBoard.transitions = make([]int8, len(board.transitions))
	copy(newBoard.transitions, board.transitions)
	newBoard.transitions = append(newBoard.transitions, action)
	newBoard.depth = board.depth + 1

	i, j, _ := board.findZero()
	switch action {
	case ShiftLeft:
		if i == 0 {
			newBoard = nil
		} else {
			newBoard.tiles[i][j] = newBoard.tiles[i-1][j]
			newBoard.tiles[i-1][j] = 0
		}
	case ShiftRight:
		if i == board.size-1 {
			newBoard = nil
		} else {
			newBoard.tiles[i][j] = newBoard.tiles[i+1][j]
			newBoard.tiles[i+1][j] = 0
		}
	case ShiftUp:
		if j == 0 {
			newBoard = nil
		} else {
			newBoard.tiles[i][j] = newBoard.tiles[i][j-1]
			newBoard.tiles[i][j-1] = 0
		}
	case ShiftDown:
		if j == board.size-1 {
			newBoard = nil
		} else {
			newBoard.tiles[i][j] = newBoard.tiles[i][j+1]
			newBoard.tiles[i][j+1] = 0
		}
	}
	return newBoard
}

func (board *Board) reversedActionBoard(action int8) *Board {
	switch action {
	case ShiftDown:
		return board.NextBoard(ShiftUp)
	case ShiftUp:
		return board.NextBoard(ShiftDown)
	case ShiftRight:
		return board.NextBoard(ShiftLeft)
	case ShiftLeft:
		return board.NextBoard(ShiftRight)
	}
	return nil
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

//Actions returns the possible of the Game
func actions() []int8 {
	actions := []int8{ShiftUp, ShiftDown, ShiftLeft, ShiftRight}
	return actions
}

// PrintPath prints the path from the first node to this board
func (board *Board) PrintPath() {
	if board == nil {
		fmt.Println("Cannot print path from nil board")
	} else {
		fmt.Println("There are", board.depth, "transitions")
		boards := make([]*Board, 0)
		pt := board
		boards = append(boards, pt)
		for i := len(board.transitions) - 1; i >= 0; i-- {
			pt = pt.reversedActionBoard(board.transitions[i])
			if pt != nil {
				boards = append(boards, pt)
			} else {
				break
			}
		}
		for i := len(boards) - 1; i >= 0; i-- {
			fmt.Println(boards[i])
			fmt.Println()
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
		val := board.summedManhattan()
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
