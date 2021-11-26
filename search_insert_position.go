/**
  35. Search Insert Position
  https://leetcode.com/problems/search-insert-position/
 */

func searchInsert(nums []int, target int) int {
    left, right := 0, len(nums)-1
    
    for left <= right {
        middle := left + (right-left)/2
        if nums[middle] == target {
            return middle
        } else if (nums[middle] > target) {
            right = middle -1
        } else {
            left = middle + 1
        }
    }
    return left
}
