package valid_sudoku

const ROWS = 9
const COLUMNS = 9

func isValidSudoku(board [][]byte) bool {

	// Check rows
	for rowIndex := 0; rowIndex < ROWS; rowIndex += 1 {
		lookup := make(map[byte]struct{})
		for colIndex := 0; colIndex < COLUMNS; colIndex += 1 {
			cell := board[rowIndex][colIndex]
			if cell == '.' {
				continue
			}

			_, exists := lookup[cell]
			if exists {
				return false
			}

			lookup[cell] = struct{}{}
		}
	}

	// Check cols
	for colIndex := 0; colIndex < COLUMNS; colIndex += 1 {
		lookup := make(map[byte]struct{})
		for rowIndex := 0; rowIndex < ROWS; rowIndex += 1 {
			cell := board[rowIndex][colIndex]
			if cell == '.' {
				continue
			}

			_, exists := lookup[cell]
			if exists {
				return false
			}

			lookup[cell] = struct{}{}
		}
	}

	// Checks Squares
	for rowCornerIndex := 0; rowCornerIndex < ROWS; rowCornerIndex += 3 {
		for colCornerIndex := 0; colCornerIndex < COLUMNS; colCornerIndex += 3 {
			lookup := make(map[byte]struct{})
			for rowVenture := 0; rowVenture < 3; rowVenture += 1 {
				for colVenture := 0; colVenture < 3; colVenture += 1 {
					cell := board[rowCornerIndex+rowVenture][colCornerIndex+colVenture]
					if cell == '.' {
						continue
					}

					_, exists := lookup[cell]
					if exists {
						return false
					}

					lookup[cell] = struct{}{}
				}
			}
		}
	}

	return true
}
