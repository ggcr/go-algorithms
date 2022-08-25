// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be
// validated according to the following rules:

// 1 - Each row must contain the digits 1-9 without repetition.

// 2 - Each column must contain the digits 1-9 without repetition.

// 3 - Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without
//     repetition.

package leetcode36

import "fmt"

// Time: O(m*n)
// We split the sudoku into the boxes
// . Box 1 = row(0..2) col(0..2)
func isValidSudoku(board [][]byte) bool {
	hashmap := make(map[byte]int, 9)
	arr_rows := make(map[string]int, 9)
	arr_cols := make(map[string]int, 9)
	for i := 0; i < 9/3; i += 1 {
		for j := 0; j < 9/3; j += 1 {
			for k := 0; k < 3; k += 1 {
				for l := 0; l < 3; l += 1 {
					value := board[i*3+k][j*3+l]
					if value != byte('.') {
						s_rows := fmt.Sprint("", i*3+k, " ", value)
						s_cols := fmt.Sprint("", j*3+l, " ", value)
						if hashmap[value] == 0 &&
							arr_rows[s_rows] == 0 &&
							arr_cols[s_cols] == 0 {
							hashmap[value] = 1
							arr_rows[s_rows] = 1
							arr_cols[s_cols] = 1
							continue
						}
						return false
					}
				}
			}
			hashmap = make(map[byte]int, 9)
		}
	}
	return true
}
