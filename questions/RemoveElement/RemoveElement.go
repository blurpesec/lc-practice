package leetcode

// ToDo: Given an array nums and a value val,
// remove all instances of that value in-place and return the new length.

/*
	How to:
	1) Shuffle through elements using 2 iterators, re-adding each element 1 step at a time, 
	excluding the elements containing val.
*/
func removeElement(elementArray []int, val int) int {
	if len(elementArray) == 0 {
		return 0
	}
	j := 0
	for i := 0; i < len(elementArray); i++ {
		if elementArray[i] != val {
			if i != j {
				elementArray[i], elementArray[j] = elementArray[j], elementArray[i]
			}
			j++
		}
	}
	return j
}