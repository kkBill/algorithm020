//给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。 
//
// 
//
// 示例 1: 
//
// 输入: [2,3,-2,4]
//输出: 6
//解释: 子数组 [2,3] 有最大乘积 6。
// 
//
// 示例 2: 
//
// 输入: [-2,0,-1]
//输出: 0
//解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。 
// Related Topics 数组 动态规划 
// 👍 871 👎 0


package main

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func maxProduct(nums []int) int {
	curMin, curMax, maxProd := 1, 1, math.MinInt64
	for _, num := range nums {
		if num <= 0 {
			curMaxOrigin := curMax
			curMax = max(num, curMin*num)
			curMin = min(num, curMaxOrigin*num)
		}else {
			curMax = max(num, curMax*num)
			curMin = min(num, curMin*num)
		}
		maxProd = max(maxProd, curMax)
	}
	return maxProd
}

//func max(x, y int) int {
//	if x >= y {
//		return x
//	}else {
//		return y
//	}
//}
//
//func min(x, y int) int {
//	if x < y {
//		return x
//	}else {
//		return y
//	}
//}
//leetcode submit region end(Prohibit modification and deletion)


func main() {

}