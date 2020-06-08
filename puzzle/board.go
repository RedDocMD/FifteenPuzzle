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
	f := Board{tiles, size, transitions, 0}
	return &f
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
		fmt.Println("There are", board.depth, "steps")
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
