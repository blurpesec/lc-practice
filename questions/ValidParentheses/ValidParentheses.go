package leetcode

// Thought process:
/*
	Things to consider here:
	1) order of parentheses
	2) opening && closing bracket match

	maybe - stack?
	to implement: 
	1) for each open bracket, push it into a stack
	2) for each close bracket, check the top of stack. If it matches, pop the top element out of the stack; else -> return false;
	3) if there is leftover, in the stack at the end of the string, return false;

	to make a stack:
	1) create a struct with Element { next: Element, data: string, prev: Element };
	2) for each item in array, if item == open parentheses; data = item, element.next = new Element({  }) && element = element.next
*/

type Stack []rune

func (s Stack) stackPush (newString rune) Stack {
	return append(s,newString);
}

func (s Stack) stackPop () Stack {
	return s[:len(s)-1]
}

func isValidParentheses(inputString string) bool {
	if (len(inputString) == 0) {
		return true
	}
	stack := make(Stack, 0)
	for _, v := range inputString {
		
		if (v == '[') || (v == '(') || (v == '{') {
			stack = stack.stackPush(v)
		} else if ((v == ']') && len(stack) > 0 && stack[len(stack)-1] == '[') ||
			((v == ')') && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			((v == '}') && len(stack) > 0 && stack[len(stack)-1] == '{') {
			stack = stack.stackPop()
		} else {
			return false
		}
	}
	return len(stack) == 0
}