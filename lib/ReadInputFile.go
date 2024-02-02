package lib

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInputFile(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("ERROR")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	Tetromino := [][]string{}
	tetrominoLines := []string{}
	for scanner.Scan() {
		currentLine := scanner.Text()
		if len(currentLine) == 4 {
			tetrominoLines = append(tetrominoLines, currentLine)
		} else if currentLine != "" && len(currentLine) != 4 {
			return nil, fmt.Errorf("ERROR")
		}

		if currentLine == "" {
			Tetromino = append(Tetromino, tetrominoLines)
			tetrominoLines = []string{}
		} else if currentLine != "" && len(tetrominoLines) > 4 {
			return nil, fmt.Errorf("ERROR")
		}
	}
	if len(tetrominoLines) > 0 {
		Tetromino = append(Tetromino, tetrominoLines)
	}

	return Tetromino, nil
}
