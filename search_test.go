package main

import (
	"RedDocMD/fifteen_puzzle/puzzle"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func BenchmarkIDAStar(t *testing.B) {
	t.ReportAllocs()
	const size int8 = 4
	tiles := make([][]int8, size)
	for i := range tiles {
		tiles[i] = make([]int8, size)
	}
	filename := "default_input"
	readFromFile(t, tiles, size, filename)
	board := puzzle.NewBoard(tiles, size, puzzle.InversionDistance)
	solved := puzzle.IterativeDeepeningAStar(board, 12)
	if solved != nil {
		fmt.Println("Solved board")
		solved.PrintPath()
	} else {
		fmt.Println("Could not solve in given limit")
	}
}

func readFromFile(t testing.TB, tiles [][]int8, size int8, filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	dataString := string(data)
	lines := strings.Split(dataString, "\n")
	if int8(len(lines)) != size {
		t.Fatal(errors.New("Invalid input file: incompatible sizes"))
	}
	for i := range lines {
		numbers := strings.Split(lines[i], " ")
		if int8(len(numbers)) != size {
			t.Fatal(errors.New("Invalid input file: incompatible sizes"))
		}
		for j := range numbers {
			number, err := strconv.Atoi(numbers[j])
			if err != nil {
				t.Fatal(err)
			}
			tiles[i][j] = int8(number)
		}
	}
}
