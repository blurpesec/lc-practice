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
	one     []int
	two			int
}

type ans struct {
	one int
}

func Test_Problem(t *testing.T) {

	qs := []question{

		question{
			para{[]int{1,3,5,6}, 5},
			ans{2},
		},

		question{
			para{[]int{1,3,5,6}, 2},
			ans{1},
		},

		question{
			para{[]int{1,3,5,6}, 7},
			ans{4},
		},

		question{
			para{[]int{1,3,5,6}, 0},
			ans{0},
		},

	}

	fmt.Printf("------------------------Leetcode Problem 35------------------------\n")

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("【input】:%v 【output】:%v 【expectedOutput】:%v\n", p.one, searchInsert(p.one, p.two), a.one)
	}
	fmt.Printf("\n\n\n")
}
