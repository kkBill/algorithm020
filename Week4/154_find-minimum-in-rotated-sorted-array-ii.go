//假设按照升序排序的数组在预先未知的某个点上进行了旋转。 
//
// ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。 
//
// 请找出其中最小的元素。 
//
// 注意数组中可能存在重复的元素。 
//
// 示例 1： 
//
// 输入: [1,3,5]
//输出: 1 
//
// 示例 2： 
//
// 输入: [2,2,2,0,1]
//输出: 0 
//
// 说明： 
//
// 
// 这道题是 寻找旋转排序数组中的最小值 的延伸题目。 
// 允许重复会影响算法的时间复杂度吗？会如何影响，为什么？ 
// 
// Related Topics 数组 二分查找 
// 👍 212 👎 0


package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func findMin_154(nums []int) int {
	left, right, mid := 0, len(nums)-1, 0
	for left <= right {
		mid = (right - left)/2 + left
		if nums[mid] == nums[right] {
			right--
		}else if nums[mid] < nums[right] { // [mid, right]为增序，则最小值一定在[left, mid]中
			right = mid
		}else { // [left, mid]为增序，且在[mid, right]中必定存在截断转折点，则故最小值一定在(mid, right]中
			left = mid+1
		}
	}
	return nums[left]
}
//leetcode submit region end(Prohibit modification and deletion)


func main() {
	fmt.Printf("%d\n", findMin_154([]int{2,2,2,0,1}))
	fmt.Printf("%d\n", findMin_154([]int{2,2,4,0,1}))
	fmt.Printf("%d\n", findMin_154([]int{1}))
}