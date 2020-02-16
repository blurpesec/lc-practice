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
			para{1},
			ans{1},
		},

		question{
			para{2},
			ans{11},
		},

		question{
			para{3},
			ans{21},
		},

		question{
			para{4},
			ans{1211},
		},

		question{
			para{5},
			ans{111221},
		},

		question{
			para{6},
			ans{312211},
		},

		question{
			para{7},
			ans{13112221},
		},
		
	}

	fmt.Printf("------------------------Leetcode Problem 38 ------------------------\n")

	for _, q := range qs {
		_, p := q.ans, q.para
		fmt.Printf("【input】:%v    【output】:%v\n", p.one, countAndSay(p.one))
	}
	fmt.Printf("\n\n\n")
}
