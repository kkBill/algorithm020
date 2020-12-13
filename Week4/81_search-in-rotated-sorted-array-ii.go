//假设按照升序排序的数组在预先未知的某个点上进行了旋转。 
//
// ( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。 
//
// 编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。 
//
// 示例 1: 
//
// 输入: nums = [2,5,6,0,0,1,2], target = 0
//输出: true
// 
//
// 示例 2: 
//
// 输入: nums = [2,5,6,0,0,1,2], target = 3
//输出: false 
//
// 进阶: 
//
// 
// 这是 搜索旋转排序数组 的延伸题目，本题中的 nums 可能包含重复元素。 
// 这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？ 
// 
// Related Topics 数组 二分查找 
// 👍 257 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func search_81(nums []int, target int) bool {
	left, right, mid := 0, len(nums)-1, 0
	for left <= right {
		mid = (right-left)/2 + left
		if nums[mid] == target {
			return true
		}
		if nums[mid] == nums[right] {
			right--
		}else if nums[mid] < nums[right] { // [mid, right] 有序
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else { // [left, mid] 有序
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Printf("%v\n", search_81([]int{2, 5, 6, 0, 0, 1, 2}, 0)) // true
	fmt.Printf("%v\n", search_81([]int{2, 5, 6, 0, 0, 1, 2}, 2)) // true
	fmt.Printf("%v\n", search_81([]int{2, 5, 6, 0, 0, 1, 2}, 3)) // false
	fmt.Printf("%v\n", search_81([]int{1, 1, 3, 1}, 3))          // true
}
