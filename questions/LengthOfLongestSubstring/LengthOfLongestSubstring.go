package leetcode

import "strings"

func lengthOfLongestSubstring(str string) int {
	if (len(str) == 0 ) {
		return 0
	}

	startIndex := 0
	// max will always be >= 1 if # of chars is >= 1
	max := 1

	// loop through length of input string
	for endIndex := 1 ; endIndex < len(str); endIndex++ {
		j := strings.IndexByte(str[startIndex:endIndex], str[endIndex])
		// if j exists.
		if  j != -1 {
			// set startIndex = the previous startIndex + j + 1
			startIndex = startIndex + j + 1
		}
		// if this is a new longest length, set it as the new longest length
		if (max < (endIndex - startIndex + 1)) {
			max = endIndex - startIndex + 1
		}
	}
	return max
}