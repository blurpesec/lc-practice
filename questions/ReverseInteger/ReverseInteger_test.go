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
	one int
}

func Test_Problem(t *testing.T) {

	qs := []question{

		question{
			para{321},
			ans{123},
		},

		question{
			para{-123},
			ans{-321},
		},

		question{
			para{120},
			ans{21},
		},

		question{
			para{1534236469},
			ans{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 7 ------------------------\n")

	for _, q := range qs {
		_, p := q.ans, q.para
		fmt.Printf("【input】:%v    【output】:%v\n", p.one, reverseInteger(p.one))
	}
	fmt.Printf("\n\n\n")
}
