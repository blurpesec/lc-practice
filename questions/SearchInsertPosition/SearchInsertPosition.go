package leetcode

// ToDo: Given a sorted array and a target value, return the index if the target is found. 
// If not, return the index where it would be if it were inserted in order.
/*
	How to:
	1) Step through the array. if target = nums[index] || target < nums[index], return index
*/
func searchInsert(nums []int, target int) int {
	for idx := 0; idx < len(nums); idx++ {
		if target <= nums[idx]  {
			return idx
		}
	}
	return len(nums)
}