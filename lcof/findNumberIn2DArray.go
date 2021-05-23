package lcof

// https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/

func findNumberIn2DArray(matrix [][]int, target int) bool {
	rows := len(matrix)
	if rows == 0 {
		return false
	}
	cols := len(matrix[0])
	if cols == 0 {
		return false
	}

	validIdx := func(r, c int) bool {
		return r < rows && r >= 0 &&
			c < cols && c >= 0
	}

	// start from top right
	rIdx := 0
	cIdx := cols - 1

	for {
		if !validIdx(rIdx, cIdx) {
			return false
		}

		val := matrix[rIdx][cIdx]
		if val == target {
			return true
		}

		if val > target {
			cIdx -= 1
		} else {
			rIdx += 1
		}
	}
}
