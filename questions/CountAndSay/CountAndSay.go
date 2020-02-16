package leetcode

import (
	"strconv"
	"fmt"
)

/*
	How to:
	1) Convert int to string
	2) Create a function that's capable of generating the next step from the previous step.
	3) Recursively call that function until the step generated is equal to the input n needed
	
	To create generateNextStepFunction:
	1) Read number of same digits into a stack ?
	2) Go through stack and pop them out. count type and quantity of same element. { 1: 3 }
	3) When next element is different, add the nums to an array

*/

func countAndSay(n int) string {
	if (n == 1){
		return "1"
	}
	if (n == 2){
		return "11"
	} 
			
	s := "11" 
	for i := 3; i <= n; i++ {
		fmt.Printf("[Step]: %d - %s\n", i, s)
		s += "$"
		l := len(s) 

		cnt := 1
		tmp := ""
		
		for j := 1; j < l; j++ {
			fmt.Printf("[Step2nd]: %d - %s\n", j, tmp)
			if (s[j] != s[j-1]){
				tmp += strconv.Itoa(cnt+0)
				x := s[j-1]
				tmp += string(x)
				cnt = 1
			}	else {
				cnt += 1
			}										
		}
			
		s := tmp
		fmt.Printf("Count: %d\n",cnt)
		fmt.Printf("S: %s\n",s)
	}
	return s;
}