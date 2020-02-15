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
			para{"III"},
			ans{3},
		},

		question{
			para{"IV"},
			ans{4},
		},

		question{
			para{"IX"},
			ans{9},
		},

		question{
			para{"LVIII"},
			ans{58},
		},

		question{
			para{"MCMXCIV"},
			ans{1994},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 12------------------------\n")

	for _, q := range qs {
		_, p := q.ans, q.para
		fmt.Printf("【input】:%v    【output】:%v\n", p.one, romanToInt(p.one))
	}
	fmt.Printf("\n\n\n")
}
