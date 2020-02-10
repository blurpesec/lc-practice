package leetcode

import (
	"fmt"
	"testing"
)

type question7 struct {
	para7
	ans7
}

// para 是参数
// one 代表第一个参数
type para7 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans7 struct {
	one int
}

func Test_Problem7(t *testing.T) {

	qs := []question7{

		question7{
			para7{"III"},
			ans7{3},
		},

		question7{
			para7{"IV"},
			ans7{4},
		},

		question7{
			para7{"IX"},
			ans7{9},
		},

		question7{
			para7{"LVIII"},
			ans7{58},
		},

		question7{
			para7{"MCMXCIV"},
			ans7{1994},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 12------------------------\n")

	for _, q := range qs {
		_, p := q.ans7, q.para7
		fmt.Printf("【input】:%v    【output】:%v\n", p.one, romanToInt(p.one))
	}
	fmt.Printf("\n\n\n")
}
