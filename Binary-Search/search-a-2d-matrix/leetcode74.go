package leetcode74

import "fmt"

func binarySearch(row []int, target int) bool {
	// base case
	half := len(row) / 2
	res := len(row) % 2
	idx := half + res
	fmt.Printf("row: %v, idx = %d, res = %d\n", row, idx, res)

	if len(row) == 1 {
		if row[0] == target {
			return true
		} else {
			return false
		}
	}

	if row[idx] == target {
		return true
	}

	if row[idx] < target {
		return binarySearch(row[idx-res:], target)
	} else if row[res+half] > target {
		return binarySearch(row[0:idx], target)
	}

	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if matrix[i][len(matrix[i])-1] >= target && matrix[i][0] <= target {
			fmt.Printf("entering:.... %v\n", matrix[i])
			if matrix[i][len(matrix[i])-1] == target || matrix[i][0] == target {
				return true
			}
			return binarySearch(matrix[i], target)
		}
	}
	return false
}
