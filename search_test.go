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
	idaStarTester(t, "default_input", puzzle.InversionDistance, "Inversion Distance")
}

func idaStarTester(t testing.TB, filename string, heuristicType int8, heuristicName string) {
	fmt.Println(heuristicName, filename)
	fmt.Println()
	const size int8 = 4
	tiles := make([][]int8, size)
	for i := range tiles {
		tiles[i] = make([]int8, size)
	}
	readFromFile(t, tiles, size, filename)
	board := puzzle.NewBoard(tiles, size, heuristicType)
	solved := puzzle.IterativeDeepeningAStar(board, -1)
	if solved != nil {
		fmt.Println("Solved board")
		solved.PrintPath()
	} else {
		fmt.Println("Could not solve in given limit")
	}
}

func TestIDAStarDefault(t *testing.T) {
	idaStarTester(t, "default_input", puzzle.InversionDistance, "Inversion Distance")
}

func TestIDAStarEasy(t *testing.T) {
	idaStarTester(t, "easy_input", puzzle.InversionDistance, "Inversion Distance")
}

func TestIDAStarThird(t *testing.T) {
	idaStarTester(t, "third_input", puzzle.InversionDistance, "Inversion Distance")
}

func TestIDAStarFourth(t *testing.T) {
	idaStarTester(t, "fourth_input", puzzle.InversionDistance, "Inversion Distance")
}

func TestIDAStarFifth(t *testing.T) {
	idaStarTester(t, "fifth_input", puzzle.InversionDistance, "Inversion Distance")
}

func TestIDAStarDefaultSummedManhattan(t *testing.T) {
	idaStarTester(t, "default_input", puzzle.SummedManhattan, "Summed Manhattan")
}

func TestIDAStarEasySummedManhattan(t *testing.T) {
	idaStarTester(t, "easy_input", puzzle.SummedManhattan, "Summed Manhattan")
}

func TestIDAStarEasyCombineddManhattan(t *testing.T) {
	idaStarTester(t, "easy_input", puzzle.CombinedManhattan, "Combined Manhattan")
}

func TestIDAStarThirdCombineddManhattan(t *testing.T) {
	idaStarTester(t, "third_input", puzzle.CombinedManhattan, "Combined Manhattan")
}

func TestIDAStarFourthCombineddManhattan(t *testing.T) {
	idaStarTester(t, "fourth_input", puzzle.CombinedManhattan, "Combined Manhattan")
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
