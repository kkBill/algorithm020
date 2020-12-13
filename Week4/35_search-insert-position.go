//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。 
//
// 你可以假设数组中无重复元素。 
//
// 示例 1: 
//
// 输入: [1,3,5,6], 5
//输出: 2
// 
//
// 示例 2: 
//
// 输入: [1,3,5,6], 2
//输出: 1
// 
//
// 示例 3: 
//
// 输入: [1,3,5,6], 7
//输出: 4
// 
//
// 示例 4: 
//
// 输入: [1,3,5,6], 0
//输出: 0
// 
// Related Topics 数组 二分查找 
// 👍 757 👎 0


package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func searchInsert(nums []int, target int) int {
	left, right, mid := 0, len(nums)-1, 0
	for left <= right {
		mid = (right - left) / 2 + left
		if nums[mid] == target {
			return mid
		}else if nums[mid] < target {
			left = mid + 1
		}else {
			right = mid - 1
		}
	}
	return left
}
//leetcode submit region end(Prohibit modification and deletion)


func main() {
	fmt.Printf("%d\n",searchInsert([]int{1,3,5,6}, 0)) // 0
	fmt.Printf("%d\n",searchInsert([]int{1,3,5,6}, 7)) // 4
	fmt.Printf("%d\n",searchInsert([]int{1,3,5,6}, 2)) // 1

}