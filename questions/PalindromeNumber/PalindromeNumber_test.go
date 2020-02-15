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
	one int
}

type ans struct {
	one bool
}

func Test_Problem(t *testing.T) {

	qs := []question{

		question{
			para{121},
			ans{true},
		},

		question{
			para{-121},
			ans{false},
		},

		question{
			para{10},
			ans{false},
		},

		question{
			para{12321},
			ans{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 9------------------------\n")

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("【input】:%v		【output】:%v		【expectedOutput】:%v\n", p.one, isPalindromeNumber(p.one), a.one)
	}
	fmt.Printf("\n\n\n")
}
