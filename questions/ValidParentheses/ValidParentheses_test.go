package leetcode

import (
	"fmt"
	"testing"
)

type question struct {
	para
	ans
}

type para struct {
	one string
}

type ans struct {
	one bool
}

func Test_Problem(t *testing.T) {

	qs := []question{

		question{
			para{"()"},
			ans{true},
		},

		question{
			para{"()[]{}"},
			ans{true},
		},

		question{
			para{"(]"},
			ans{false},
		},

		question{
			para{"([)]"},
			ans{false},
		},

		question{
			para{"{[]}"},
			ans{true},
		},

		question{
			para{""},
			ans{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 20------------------------\n")

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("【input】:%v		【output】:%v		【expectedOutput】:%v\n", p.one, isValidParentheses(p.one), a.one)
	}
	fmt.Printf("\n\n\n")
}
