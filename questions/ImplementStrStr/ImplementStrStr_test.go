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
	one     string
	two			string
}

type ans struct {
	one int
}

func Test_Problem(t *testing.T) {

	qs := []question{

		question{
			para{"hello", "ll"},
			ans{2},
		},

		question{
			para{"aaaaa", "bba"},
			ans{-1},
		},

		question{
			para{"aaaaab", "aab"},
			ans{3},
		},

	}

	fmt.Printf("------------------------Leetcode Problem 28------------------------\n")

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("【input】:%v 【output】:%v 【expectedOutput】:%v\n", p.one, strStr(p.one, p.two), a.one)
	}
	fmt.Printf("\n\n\n")
}
