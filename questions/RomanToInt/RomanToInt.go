package leetcode

import "strings"

func fetchRomanInteger(romanNum string) int {
	switch romanNum {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0
	}
}


func romanToInt(romanNum string) int {
	// split strings
	stringArray := strings.Split(romanNum, "")
	
	// firstValue is the return value
	value := 0
	// secondValue is the negative value on backwards stuff
	secondValue := 0

	for i := 0; i < len(stringArray); i++ {
		data := fetchRomanInteger(stringArray[i])
		rightData := 0
		if !(i+1 == len(stringArray)) {
			rightData = fetchRomanInteger(stringArray[i+1])
		}
		if data <= rightData {
			secondValue += data
		} else {
			if (secondValue != 0 && secondValue < data) {
				value += data - secondValue
			} else {
				value += data + secondValue
			}
			secondValue = 0
		}
	}
	return value

}
