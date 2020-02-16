package leetcode


// ToDo: Implement JS's .indexOf() which returns index of first occurrence of needle in haystack
/*
	How to:
	1) Find length to check (longest string length - shortest string length)
	2) Loop through haystack, looking for instances of first char of needle
	3) When found, loop through length of needle and check each index 
		against needle and substring to make sure it's valid.
	5) if it's not valid, break;
	6) if len(needle) == the length of the substring, return the index from the first loop
	
*/
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	lengthToCheck := len(haystack) - len(needle)
	for i := 0; i <= lengthToCheck ; i++ {
		if haystack[i] == needle[0] {
			for idx := 0; idx < len(needle); idx++ {
				if haystack[i+idx] != needle[idx] {
					break
				}
				if len(needle) == idx+1 {
					return i
				}		
			}
		}
	}
	return -1
}