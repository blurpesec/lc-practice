package leetcode

// Thought process:
/*
	a number isn't a palindrome if it's negative,
*/

func isPalindromeNumber(inputNumber int) bool {
	if (inputNumber < 0) {
		return false
	}
	// Creates a separate copy of inputNumber that makes modification not affect dupNum
	duplicateNum := inputNumber;

	return isPalindromeUtility(inputNumber, &duplicateNum)
}

func isSingleDigit(item int) bool {
	return (item >= 0 && item < 10)
}

func isPalindromeUtility(inputNumber int, duplicateNum *int) bool {
	
	// Recursively drills down to 1 digit in inputNumber
	if (isSingleDigit(inputNumber)) {
		return (inputNumber == ((*duplicateNum) % 10))
	}

	if (!(isPalindromeUtility(inputNumber/10, duplicateNum))) {
		// If false anywhere in the tree, drill back up with 'false'
		return false;
	}

	// Divide duplicate number to check digits against each other in reverse.
	*duplicateNum /= 10;

	// Drills back up the tree with each 'true'
	return (inputNumber % 10 == (*duplicateNum) % 10)
}