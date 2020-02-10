package leetcode

func reverseInteger(x int) int {
	// Thought process:
	// Take in an integer.
	// Pop each integer out of the stack.
	// Push it into a new stack;
	
	// Set reversed integer at 0
	reversedInt := 0;

	// while x is greater than 0, keep reversing the digits
	for x != 0 {
		// pop out a number by taking the modulo when dividing by 10.
		pop := x%10
		// then divide the input number by 10 to remove the integer from the stack.
		x = x/10;
		// if the popped number has an overflow, return 0;
		if reversedInt > ((1<<31-1)/10) || reversedInt < -((1<<31-1)/10) {
			return 0
		}
		// else, multiply the reversed integer by 10 and add the popped modulo to it to handle the new integer added..
		reversedInt = reversedInt*10 + pop
	}
	return reversedInt
}