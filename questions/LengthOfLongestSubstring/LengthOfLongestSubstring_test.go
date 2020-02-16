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
	one int
}

func Test_Problem(t *testing.T) {

	qs := []question{

		question{
			para{"abcabcbb"},
			ans{3},
		},

		question{
			para{"bbbbb"},
			ans{1},
		},

		question{
			para{"pwwkew"},
			ans{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3------------------------\n")

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("【input】:%v		【output】:%v		【expectedOutput】:%v\n", p.one, lengthOfLongestSubstring(p.one), a.one)
	}
	fmt.Printf("\n\n\n")
}
