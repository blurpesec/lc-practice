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
			para{[]int{3,2,2,3}, 3},
			ans{2},
		},

		question{
			para{[]int{0,1,2,2,3,0,4,2}, 2},
			ans{5},
		},

	}

	fmt.Printf("------------------------Leetcode Problem 27------------------------\n")

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Println(p)
		fmt.Printf("【input】:%v 【output】:%v 【expectedOutput】:%v\n", p.one, removeElement(p.one, p.two), a.one)
	}
	fmt.Printf("\n\n\n")
}
