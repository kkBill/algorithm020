//n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。 
//
// 
//
// 上图为 8 皇后问题的一种解法。 
//
// 给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。 
//
// 每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。 
//
// 
//
// 示例： 
//
// 输入：4
//输出：[
// [".Q..",  // 解法 1
//  "...Q",
//  "Q...",
//  "..Q."],
//
// ["..Q.",  // 解法 2
//  "Q...",
//  "...Q",
//  ".Q.."]
//]
//解释: 4 皇后问题存在两个不同的解法。
// 
//
// 
//
// 提示： 
//
// 
// 皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。 
// 
// Related Topics 回溯算法 
// 👍 687 👎 0

package main

import (
	"fmt"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	var res [][]string
	// 初始化棋盘，空位以"."填充
	matrix := make([][]string, n)
	for i := range matrix {
		for j := 0; j < n; j++ {
			matrix[i] = append(matrix[i], ".")
 		}
	}
	colUsed := make([]bool, n) // colUsed[i]==true 表示第i列已经有皇后了，下同
	mainDiagUsed := make([]bool, 2*n-1)
	subDiagUsed := make([]bool, 2*n-1)

	var dfs func(row int)
	dfs = func(row int) {
		// 边界条件，遍历到最后一行
		if row == n {
			var solution []string
			for i := range matrix {
				line := strings.Join(matrix[i],"")
				solution = append(solution, line)
			}
			res = append(res, solution)
			return
		}
		//
		for col := 0; col < n; col++ {
			if !colUsed[col] && !mainDiagUsed[row-col+n-1] && !subDiagUsed[row+col] {
				// 把"Q"放在(row,col)处
				colUsed[col] = true
				mainDiagUsed[row-col+n-1] = true
				subDiagUsed[row+col] = true
				matrix[row][col] = "Q"
				// 进入下一层
				dfs(row + 1)
				// 回溯
				colUsed[col] = false
				mainDiagUsed[row-col+n-1] = false
				subDiagUsed[row+col] = false
				matrix[row][col] = "."
			}
		}
	}

	dfs(0)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Printf("%v\n", solveNQueens(4))
}
