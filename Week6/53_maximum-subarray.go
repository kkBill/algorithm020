//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。 
//
// 示例: 
//
// 输入: [-2,1,-3,4,-1,2,1,-5,4]
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
// 
//
// 进阶: 
//
// 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。 
// Related Topics 数组 分治算法 动态规划 
// 👍 2761 👎 0


package main

import (
	"fmt"
	"math"
)

//leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
	maxSum, curSum := math.MinInt64, 0
	for _, x := range nums {
		curSum = intMax(x, curSum+x)
		maxSum = intMax(maxSum, curSum)
	}
	return maxSum
}

func intMax(x, y int) int {
	if x > y {
		return x
	}else {
		return y
	}
}
//leetcode submit region end(Prohibit modification and deletion)


func main() {
	fmt.Printf("%d\n", maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
	fmt.Printf("%d\n", maxSubArray([]int{-2,-1,-3}))
}