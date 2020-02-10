package leetcode

func isAllPrefixPresent(array []string, stringItem string, start int, end int) bool {
	for k := 0; k < len(array); k++ {
		for l := start; l < end; l++ {
			if array[k] == "" {
				return false
			}
			if (array[k][l] != stringItem[l]) {
				return false
			}
		}
	}
	return true
}

func longestCommonPrefix(inputStringArray []string) string {
	// Find smallest string in the array.
	longestCommonPrefix := ""
	shortestString := ""
	stringToUse := ""

	if len(inputStringArray) > 0 {
		stringToUse = inputStringArray[0]
	}
	if stringToUse == "" {
		return ""
	}
	for i := 0; i < len(inputStringArray); i++ {
		if (len(shortestString) >= len(inputStringArray[i])) || (shortestString == "") {
			shortestString = inputStringArray[i]
		}
	}
	
	low := 0
	high := len(shortestString)

	for low < high {
		mid := low + (high-low)
		if isAllPrefixPresent(inputStringArray, stringToUse, low, mid) {
			longestCommonPrefix += stringToUse[low:mid-low]
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return longestCommonPrefix
}
