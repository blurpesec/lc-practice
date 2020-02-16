package leetcode

import (
	"fmt"
	"testing"
	"github.com/blurpesec/lc-practice/utils"
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
			para{[]int{2,4,3}, []int{5,6,4}},
			ans{[]int{7,0,8}},
		},

	}

	fmt.Printf("------------------------Leetcode Problem 2------------------------\n")

	for _, q := range qs {
		a, p := q.ans, q.para
		test1 := utils.S2l(p.one)
		test2 := utils.S2l(p.another)
		test3 := AddTwoNumbers(test1, test2)
		fmt.Printf("【input】:%v		【output】:%v		【expectedOutput】:%v\n", p.one, utils.L2s(test3), a.one)
	}
	fmt.Printf("\n\n\n")
}
