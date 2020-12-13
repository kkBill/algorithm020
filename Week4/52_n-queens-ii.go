//n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。 
//
// 
//
// 上图为 8 皇后问题的一种解法。 
//
// 给定一个整数 n，返回 n 皇后不同的解决方案的数量。 
//
// 示例: 
//
// 输入: 4
//输出: 2
//解释: 4 皇后问题存在如下两个不同的解法。
//[
// [".Q..",  // 解法 1
//  "...Q",
//  "Q...",
//  "..Q."],
//
// ["..Q.",  // 解法 2
//  "Q...",
//  "...Q",
//  ".Q.."]
//]
// 
//
// 
//
// 提示： 
//
// 
// 皇后，是国际象棋中的棋子，意味着国王的妻子。皇后只做一件事，那就是“吃子”。当她遇见可以吃的棋子时，就迅速冲上去吃掉棋子。当然，她横、竖、斜都可走一或 N
//-1 步，可进可退。（引用自 百度百科 - 皇后 ） 
// 
// Related Topics 回溯算法 
// 👍 215 👎 0


package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func totalNQueens(n int) int {
	colUsed := make([]bool, n) // colUsed[i]==true 表示第i列已经有皇后了，下同
	mainDiagUsed := make([]bool, 2*n-1)
	subDiagUsed := make([]bool, 2*n-1)
	count := 0
	var dfs func(row int)
	dfs = func(row int) {
		// 边界条件，遍历到最后一行
		if row == n {
			count++
			return
		}
		for col := 0; col < n; col++ {
			if !colUsed[col] && !mainDiagUsed[row-col+n-1] && !subDiagUsed[row+col] {
				// 把"Q"放在(row,col)处
				colUsed[col] = true
				mainDiagUsed[row-col+n-1] = true
				subDiagUsed[row+col] = true
				// 进入下一层
				dfs(row + 1)
				// 回溯
				colUsed[col] = false
				mainDiagUsed[row-col+n-1] = false
				subDiagUsed[row+col] = false
			}
		}
	}
	dfs(0)
	return count
}
//leetcode submit region end(Prohibit modification and deletion)


func main() {
	fmt.Printf("%d\n",totalNQueens(4))
	fmt.Printf("%d\n",totalNQueens(8))
}