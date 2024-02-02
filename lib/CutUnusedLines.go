package lib

func CutUnusedLines(tetrominos [][]string) [][]string {
	CuttedTetrominos := [][]string{}

	// Cut unused lines horizontally
	for _, tetromino := range tetrominos {
		var newTetromino []string
		for _, line := range tetromino {
			countOfHashtagsHorizontal := 0
			for _, char := range line {
				if char == '#' {
					countOfHashtagsHorizontal++
				}
			}
			if countOfHashtagsHorizontal > 0 {
				newTetromino = append(newTetromino, line)
			}
		}
		CuttedTetrominos = append(CuttedTetrominos, newTetromino)
	}

	// Cut unused lines verticaly
	for i := range CuttedTetrominos {
		for j := len(CuttedTetrominos[i][0]) - 1; j >= 0; j-- {
			countOfHashtagsVerticaly := 0
			for k := range CuttedTetrominos[i] {
				if CuttedTetrominos[i][k][j] == '#' {
					countOfHashtagsVerticaly++
				}
			}
			if countOfHashtagsVerticaly == 0 {
				for l := range CuttedTetrominos[i] {
					CuttedTetrominos[i][l] = removeCharAtIndex(CuttedTetrominos[i][l], j)
				}
			}
		}
	}

	return CuttedTetrominos
}

func removeCharAtIndex(input string, index int) string {
	if index < 0 || index >= len(input) {
		return input
	}

	runes := []rune(input)
	runes = append(runes[:index], runes[index+1:]...)
	return string(runes)
}
