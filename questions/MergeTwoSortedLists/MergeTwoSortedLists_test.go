package leetcode

import (
	"fmt"
	"testing"
	"github.com/blurpesec/lc-practice/questions/utils"
)

type question struct {
	para
	ans
}

type para struct {
	one     []int
	another []int
}

type ans struct {
	one []int
}

func Test_Problem(t *testing.T) {

	qs := []question{

		question{
			para{[]int{}, []int{}},
			ans{[]int{}},
		},

		question{
			para{[]int{1}, []int{1}},
			ans{[]int{1, 1}},
		},

		question{
			para{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			ans{[]int{1, 1, 2, 2, 3, 3, 4, 4}},
		},

		question{
			para{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
			ans{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}},
		},

		question{
			para{[]int{1}, []int{9, 9, 9, 9, 9}},
			ans{[]int{1, 9, 9, 9, 9, 9}},
		},

		question{
			para{[]int{9, 9, 9, 9, 9}, []int{1}},
			ans{[]int{1, 9, 9, 9, 9, 9}},
		},

		question{
			para{[]int{2, 3, 4}, []int{4, 5, 6}},
			ans{[]int{2, 3, 4, 4, 5, 6}},
		},

		question{
			para{[]int{1, 3, 8}, []int{1, 7}},
			ans{[]int{1, 1, 3, 7, 8}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 21------------------------\n")

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("【input】:%v		【output】:%v		【expectedOutput】:%v\n", p.one, utils.L2s(
			mergeTwoSortedLists(utils.S2l(p.one), utils.S2l(p.another))), a.one)
	}
	fmt.Printf("\n\n\n")
}
