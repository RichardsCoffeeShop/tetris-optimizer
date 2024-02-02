package main_test

import (
	"math"
	"reflect"
	"richardscull/tetris-optimizer/lib"
	"testing"
)

func TestIsTetrominoValid(t *testing.T) {
	t.Run("Should return false when tetromino is not valid", func(t *testing.T) {
		tetromino := []string{
			"....",
			"....",
			"....",
			"....",
		}
		tetrominos := [][]string{tetromino}

		result := lib.IsTetrominoValid(tetrominos)

		if result != false {
			t.Errorf("Expected false, got %v", result)
		}
	})

	t.Run("Should return true when tetromino is valid", func(t *testing.T) {
		tetromino := []string{
			"#...",
			"##..",
			"#...",
			"....",
		}
		tetrominos := [][]string{tetromino}

		result := lib.IsTetrominoValid(tetrominos)

		if result != true {
			t.Errorf("Expected true, got %v", result)
		}
	})
}

func TestCutUnusedLines(t *testing.T) {
	t.Run("Should return tetromino without unused lines", func(t *testing.T) {
		tetromino := []string{
			"#...",
			"##..",
			"#...",
			"....",
		}
		tetrominos := [][]string{tetromino}

		result := lib.CutUnusedLines(tetrominos)

		if len(result[0]) != 3 {
			t.Errorf("Expected 3, got %v", len(result[0]))
		}
	})
}

func TestResolve(t *testing.T) {
	t.Run("Should return solved board", func(t *testing.T) {
		tetromino := []string{
			"#...",
			"##..",
			"#...",
			"....",
		}
		tetrominos := [][]string{tetromino}
		boardSize := int(math.Ceil(math.Sqrt(float64(len(tetrominos) * 4))))

		board := make([][]string, boardSize)
		for i := range board {
			board[i] = make([]string, boardSize)
			for j := range board[i] {
				board[i][j] = "."
			}
		}

		result := lib.Resolve(tetrominos, board)

		expectedResult := [][]string{
			{"A", ".", "."},
			{"A", "A", "."},
			{"A", ".", "."},
		}

		if reflect.DeepEqual(result, expectedResult) != true {
			t.Errorf("Expected %v, got %v", expectedResult, result)
		}
	})
}
