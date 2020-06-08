package main

import (
	"RedDocMD/fifteen_puzzle/puzzle"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Fifteen puzzle solver")

	const size = 4
	tiles := make([][]int, size)
	for i := range tiles {
		tiles[i] = make([]int, size)
	}
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <filename>", os.Args[0])
		os.Exit(1)
	}
	filename := os.Args[1]
	readIntoFile(tiles, size, filename)
	board := puzzle.NewBoard(tiles, size)

	fmt.Println("Start position")
	fmt.Println(*board)

	solved := puzzle.IterativeDeepeningSearch(board)
	if solved != nil {
		fmt.Println("Solved board")
		solved.PrintPath()
	} else {
		fmt.Println("Could not solve in given limit")
	}
}

func readIntoFile(tiles [][]int, size int, filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	dataString := string(data)
	lines := strings.Split(dataString, "\n")
	if len(lines) != size {
		panic(errors.New("Invalid input file: incompatible sizes"))
	}
	for i := range lines {
		numbers := strings.Split(lines[i], " ")
		if len(numbers) != size {
			panic(errors.New("Invalid input file: incompatible sizes"))
		}
		for j := range numbers {
			number, err := strconv.Atoi(numbers[j])
			if err != nil {
				panic(err)
			}
			tiles[i][j] = number
		}
	}
}
