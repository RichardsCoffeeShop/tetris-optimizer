package lib

func IsTetrominoValid(Tetrominos [][]string) bool {
	for _, Tetromino := range Tetrominos {
		TetrominoConnections := 0
		HashtagsCount := 0
		for indexVertical, line := range Tetromino {
			for indexHorizontal, char := range line {
				HashtagConnections := 0
				if char != '#' && char != '.' {
					return false
				} else if char == '#' {
					HashtagsCount++
					// ! https://miro.medium.com/v2/resize:fit:720/format:webp/1*0i1ZfKxMiFWzdTsQoAGxQw.png
					if indexVertical > 0 && Tetromino[indexVertical-1][indexHorizontal] == '#' {
						HashtagConnections++
					}
					if indexVertical < len(Tetromino)-1 && Tetromino[indexVertical+1][indexHorizontal] == '#' {
						HashtagConnections++
					}
					if indexHorizontal > 0 && Tetromino[indexVertical][indexHorizontal-1] == '#' {
						HashtagConnections++
					}
					if indexHorizontal < len(line)-1 && Tetromino[indexVertical][indexHorizontal+1] == '#' {
						HashtagConnections++
					}
					if HashtagConnections == 0 {
						return false
					} else {
						TetrominoConnections += HashtagConnections
					}
				}
			}
		}
		if TetrominoConnections < 6 || HashtagsCount > 4 {
			return false
		}
	}

	return true
}
