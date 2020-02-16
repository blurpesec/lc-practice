package leetcode

func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) []int {
	if m == 0 {
		copy(nums1, nums2)
		return nums1
	}
	// index of arr 1, starting at the top
	i := m - 1
	// index of arr 2, starting at the top
	j := n - 1
	// index of output arr
	k := m + n - 1
	// Loop through array 1 and array 2 while k (total array length) is decrementing
	for ; i >= 0 && j >= 0; k-- {
		// if element in array 1 @ indexer i is less than element in array 2 @ indexer j, 
		// add that element to nums1 at index k
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			// else, element in array 2 at index j is greater ->
			// add that element to nums1 at index k. then decrement indexer j
			nums1[k] = nums2[j]
			j--
		}
	}
	// This cleans up any remaining elements in array 2
	for ; j >= 0; k-- {
		nums1[k] = nums2[j]
		j--
	}
	return nums1
}