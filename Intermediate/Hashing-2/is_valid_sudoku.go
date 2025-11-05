package hashing2

/*
Determine if a Sudoku is valid, according to: http://sudoku.com.au/TheRules.aspx
The Sudoku board could be partially filled, where empty cells are filled with the character '.'.
The input corresponding to the above configuration :
["53..7....", "6..195...", ".98....6.", "8...6...3", "4..8.3..1", "7...2...6", ".6....28.", "...419..5", "....8..79"]
A partially filled sudoku which is valid.

Note:

A valid Sudoku board (partially filled) is not necessarily solvable. Only the filled cells need to be validated.
*/

/**
 * @input A : array of strings
 *
 * @Output Integer
 */
func isValidSudoku(A []string) int {
	rowMaps := make(map[int]map[rune]struct{})
	colMaps := make(map[int]map[rune]struct{})
	cellMaps := make(map[int]map[int]map[rune]struct{})

	for i, rowString := range A {
		for j, currRune := range rowString {
			if currRune == '.' {
				continue
			}

			if _, exists := rowMaps[i]; !exists {
				rowMaps[i] = make(map[rune]struct{})
			}

			if _, exists := colMaps[j]; !exists {
				colMaps[j] = make(map[rune]struct{})
			}

			if _, exists := cellMaps[i/3]; !exists {
				cellMaps[i/3] = make(map[int]map[rune]struct{})
			}

			if _, exists := cellMaps[i/3][j/3]; !exists {
				cellMaps[i/3][j/3] = make(map[rune]struct{})
			}

			if _, exists := rowMaps[i][currRune]; exists {
				return 0
			}

			if _, exists := colMaps[j][currRune]; exists {
				return 0
			}

			if _, exists := cellMaps[i/3][j/3][currRune]; exists {
				return 0
			}

			rowMaps[i][currRune] = struct{}{}
			colMaps[j][currRune] = struct{}{}
			cellMaps[i/3][j/3][currRune] = struct{}{}
		}
	}

	return 1

	/* if !validateRow(A) || !validateColumn(A) || !validateCell(A){
	       return 0
	   }

	   return 1
	*/
}

/*
func validateRow(A []string) bool{
    for _, row := range A{
        rowMap := make(map[rune]struct{})
        for _, rowVal := range row{
              if rowVal == '.'{
                  continue
              }

              if _, exists := rowMap[rowVal]; exists{
                return false
              }

              rowMap[rowVal] = struct{}{}
        }
    }

    return true
}

func validateColumn(A []string) bool{
     column := [...]int{0,1,2,3,4,5,6,7,8}
     for _, colNumber := range column{
        colMap := make(map[rune]struct{})

        for _, rowString := range A{
              currVal := rune(rowString[colNumber])
              if currVal == '.'{
                  continue
              }
              if _, exists := colMap[rune(currVal)]; exists{
                return false
              }

              colMap[currVal] = struct{}{}
        }
    }

    return true
}

func validateCell(A []string) bool{
    segment := []int{0,1,2,3,4,5,6,7,8}

    for _, seg := range segment{
        rowStart := seg/3 * 3
        colStart := seg%3 * 3

        segMentMap := make(map[rune]struct{})
        for i:= rowStart; i<rowStart+3; i++{
            for j:= colStart; j<colStart+3; j++{
               dotRune := '.'
               if rune(A[i][j]) == dotRune{
                    continue
                }
               if _,exists := segMentMap[rune(A[i][j])]; exists{
                   return false
               }

               segMentMap[rune(A[i][j])] = struct{}{}
            }
        }
    }

    return true
}

*/
