package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"richardscull/tetris-optimizer/lib"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		return
	}

	arg := os.Args[1]
	Tetrominos, err := lib.ReadInputFile(arg)
	if err != nil {
		errMsg := err.Error()
		fmt.Println(errMsg)
		return
	}

	if !lib.IsTetrominoValid(Tetrominos) || len(Tetrominos) == 0 {
		fmt.Println("ERROR")
		return
	}

	// Calculate how big board should be using this formula: ceil(sqrt(4 * number_of_tetrominos))
	boardSize := int(math.Ceil(math.Sqrt(float64(len(Tetrominos) * 4))))

	board := make([][]string, boardSize)
	for i := range board {
		board[i] = make([]string, boardSize)
		for j := range board[i] {
			board[i][j] = "."
		}
	}

	compressedTetrominos := lib.CutUnusedLines(Tetrominos)

	fmt.Println("\n🔍 Started solving...")
	timeStarted := time.Now()

	resolvedBoard := lib.Resolve(compressedTetrominos, board)
	fmt.Println("\n⌛ Time took to solve: ", time.Since(timeStarted))

	fmt.Println("\n📝 Result:")
	fmt.Println()

	for _, row := range resolvedBoard {
		for _, char := range row {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}
