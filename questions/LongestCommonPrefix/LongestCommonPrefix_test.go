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
	array []string
}

type ans struct {
	one string
}

func Test_Problem14(t *testing.T) {
	qs := []question{

		question {
			para{
				[]string{"flower","flow","flight"},
			},
			ans{"fl"},
		},

		question{
			para{
				[]string{"dog","racecar","car"},
			},
			ans{""},
		},

		question{
			para{
				[]string{"","b"},
			},
			ans{""},
		},

		question{
			para{
				[]string{"bb","","b"},
			},
			ans{""},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 14------------------------\n")

	for _, q := range qs {
		_, p := q.ans, q.para
		fmt.Printf("【input】:%v    【output】:%v\n", p.array, longestCommonPrefix(p.array))
	}
	fmt.Printf("\n\n\n")
}
