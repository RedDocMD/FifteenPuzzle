package puzzle

import (
	"errors"
	"fmt"
)

// Board represents the n puzzle board at any given time
type Board struct {
	tiles  [][]int
	size   int
	parent *Board
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
func NewBoard(tiles [][]int, size int) *Board {
	if !checkTiles(tiles, size) {
		return nil
	}
	f := Board{tiles, size, nil}
	return &f
}

func checkTiles(tiles [][]int, size int) bool {
	cnt := make(map[int]int)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			cnt[tiles[i][j]]++
		}
	}
	for i := 0; i < size*size; i++ {
		if cnt[i] != 1 {
			return false
		}
	}
	return true
}

// NextBoard generates the next Board corresponding to the given action or nil of not possible
func (board *Board) NextBoard(action int) *Board {
	newBoard := new(Board)
	newBoard.tiles = board.tiles
	newBoard.parent = board
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

func (board *Board) findZero() (int, int, error) {
	for i := 0; i < board.size; i++ {
		for j := 0; j < board.size; j++ {
			if board.tiles[i][j] == 0 {
				return i, j, nil
			}
		}
	}
	return -1, -1, errors.New("Invalid board")
}

// Solved checks if the board is in the position which is the solved state
func (board *Board) Solved() bool {
	size := board.size
	for i := 0; i < size*size-1; i++ {
		if board.tiles[i/size][i%size] != i+1 {
			return false
		}
	}
	return board.tiles[size-1][size-1] == 0
}

// String returns string representation of Board
func (board Board) String() string {
	var str string = ""
	for i := 0; i < board.size; i++ {
		for j := 0; j < board.size; j++ {
			str += fmt.Sprintf("%3d", board.tiles[i][j])
		}
		str += "\n"
	}
	return str
}
