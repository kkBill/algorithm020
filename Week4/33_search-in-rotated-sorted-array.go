//给你一个整数数组 nums ，和一个整数 target 。 
//
// 该整数数组原本是按升序排列，但输入时在预先未知的某个点上进行了旋转。（例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2]
// ）。 
//
// 请你在数组中搜索 target ，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。 
// 
//
// 示例 1： 
//
// 
//输入：nums = [4,5,6,7,0,1,2], target = 0
//输出：4
// 
//
// 示例 2： 
//
// 
//输入：nums = [4,5,6,7,0,1,2], target = 3
//输出：-1 
//
// 示例 3： 
//
// 
//输入：nums = [1], target = 0
//输出：-1
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 5000 
// -10^4 <= nums[i] <= 10^4 
// nums 中的每个值都 独一无二 
// nums 肯定会在某个点上旋转 
// -10^4 <= target <= 10^4 
// 
// Related Topics 数组 二分查找 
// 👍 1097 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	left, right, mid := 0, len(nums)-1, 0
	for left <= right {
		mid = (right-left)/2 + left
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < nums[right] { // [mid, right]有序
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else { // [left, mid]有序
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Printf("%d\n", search([]int{4, 5, 6, 7, 0, 1, 2}, 0)) // 4
	fmt.Printf("%d\n", search([]int{4, 5, 6, 7, 0, 1, 2}, 4)) // 0
	fmt.Printf("%d\n", search([]int{4, 5, 6, 7, 0, 1, 2}, 1)) // 5
	fmt.Printf("%d\n", search([]int{2, 1}, 1)) // 1
}
