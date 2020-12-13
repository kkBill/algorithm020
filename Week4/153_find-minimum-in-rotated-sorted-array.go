//假设按照升序排序的数组在预先未知的某个点上进行了旋转。例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] 。 
//
// 请找出其中最小的元素。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [3,4,5,1,2]
//输出：1
// 
//
// 示例 2： 
//
// 
//输入：nums = [4,5,6,7,0,1,2]
//输出：0
// 
//
// 示例 3： 
//
// 
//输入：nums = [1]
//输出：1
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 5000 
// -5000 <= nums[i] <= 5000 
// nums 中的所有整数都是 唯一 的 
// nums 原来是一个升序排序的数组，但在预先未知的某个点上进行了旋转 
// 
// Related Topics 数组 二分查找 
// 👍 310 👎 0


package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func findMin_153(nums []int) int {
	left, right, mid := 0, len(nums)-1, 0
	for left <= right {
		mid = (right - left)/2 + left
		if nums[mid] <= nums[right] { // [mid, right]为增序，则最小值一定在[left, mid]中
			right = mid
		}else { // [left, mid]为增序，且在[mid, right]中必定存在截断转折点，则故最小值一定在(mid, right]中
			left = mid+1
		}
	}
	return nums[right]
}
//leetcode submit region end(Prohibit modification and deletion)


func main() {
	fmt.Printf("%d\n", findMin_153([]int{4,5,6,7,0,1,2}))
	fmt.Printf("%d\n", findMin_153([]int{1}))
	fmt.Printf("%d\n", findMin_153([]int{2,1}))
}