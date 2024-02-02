package lib

type pos struct {
	X int
	Y int
}

type save struct {
	board [][]string
	pos   [][]pos
}

func getBoard(board [][]string) [][]string {
	newBoard := [][]string{}

	for _, row := range board {
		newRow := []string{}

		newRow = append(newRow, row...)

		newBoard = append(newBoard, newRow)
	}

	return newBoard
}

func Resolve(tetrominos [][]string, board [][]string) [][]string {
	boardSaves := []save{}

	for i := 0; i < len(tetrominos); i++ {
		beenPlaced := false        // Check if we have placed tetrimo
		isOnLastPos := true        // Check if we are on last possible position of tetrimo
		wasDuplicateFound := false // Check if we have duplicate position
		couldPutTetromino := false // Check if we could put tetrimo
		Positions := [][]pos{}     // Positions of tetrimo we have tried

		for indexHorizontal := 0; indexHorizontal < len(board[0]) && !beenPlaced; indexHorizontal++ {
			for indexVertical := 0; indexVertical < len(board) && !beenPlaced; indexVertical++ {
				if CanPutTetromino(tetrominos[i], board, indexHorizontal, indexVertical) && !beenPlaced {
					couldPutTetromino = true
					hasDuplicate := false

					if len(boardSaves) > 0 && len(boardSaves) == i+1 {
						for _, pos := range boardSaves[len(boardSaves)-1].pos[len(boardSaves[len(boardSaves)-1].pos)-1] {
							if pos.X == indexHorizontal && pos.Y == indexVertical {
								hasDuplicate = true
								wasDuplicateFound = true
							}
						}
					}

					if len(boardSaves) > 0 {
						Positions = boardSaves[len(boardSaves)-1].pos
					}

					if hasDuplicate {
						continue
					}

					if !wasDuplicateFound {
						Positions = append(Positions, []pos{})
					}

					beenPlaced = true

					localBoard := getBoard(board)
					Positions[len(Positions)-1] = append(Positions[len(Positions)-1], pos{indexHorizontal, indexVertical})

					if !wasDuplicateFound {
						isOnLastPos = false
						boardSaves = append(boardSaves, save{localBoard, Positions})
					}

					board = PutTetromino(tetrominos[i], board, indexHorizontal, indexVertical, i)
				}
			}
		}

		if !couldPutTetromino {
			isOnLastPos = false
		}

		if !beenPlaced {
			i-- // Go back in next cycle to current tetrimo

			// Make new board if we are out of space
			if len(boardSaves) <= 1 {
				i = -1 // Go back in next cycle to the first tetrimo

				boardSize := len(board) + 1
				board = make([][]string, boardSize)
				for i := range board {
					board[i] = make([]string, boardSize)
					for j := range board[i] {
						board[i][j] = "."
					}
				}

				boardSaves = []save{}
			}

			if len(boardSaves) > 1 {
				i-- // Go back in next cycle to the previous tetrimo

				if isOnLastPos {
					{
						boardSaves = boardSaves[:len(boardSaves)-1]
						board = getBoard(boardSaves[len(boardSaves)-1].board)
					}
				} else {
					board = getBoard(boardSaves[len(boardSaves)-1].board)
				}
			}
		}
	}

	return board
}

func CanPutTetromino(tetromino []string, board [][]string, indexVertical int, indexHorizontal int) bool {
	for indexVerticalTetromino, line := range tetromino {
		for indexHorizontalTetromino, char := range line {
			if char == '#' {
				if indexVertical+indexVerticalTetromino > len(board)-1 || indexHorizontal+indexHorizontalTetromino > len(board)-1 {
					return false
				}
				if board[indexVertical+indexVerticalTetromino][indexHorizontal+indexHorizontalTetromino] != "." {
					return false
				}
			}
		}
	}
	return true
}

func PutTetromino(tetromino []string, board [][]string, indexVertical int, indexHorizontal int, tetrominoIndex int) [][]string {
	for indexVerticalTetromino, line := range tetromino {
		for indexHorizontalTetromino, char := range line {
			if char == '#' {
				startOfAlphabet := 65 + tetrominoIndex
				if startOfAlphabet > 90 {
					board[indexVertical+indexVerticalTetromino][indexHorizontal+indexHorizontalTetromino] = "X"
				} else {
					board[indexVertical+indexVerticalTetromino][indexHorizontal+indexHorizontalTetromino] = string(rune(startOfAlphabet))
				}

			}
		}
	}
	return board
}
