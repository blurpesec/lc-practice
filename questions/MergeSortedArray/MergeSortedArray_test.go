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
	two 		int
	three		[]int
	four 		int
}

type ans struct {
	one []int
}

func Test_Problem(t *testing.T) {

	qs := []question{

		question{
			para{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
			ans{[]int{1, 2, 2, 3, 5, 6}},
		},

	}

	fmt.Printf("------------------------Leetcode Problem 88------------------------\n")

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("【input】:%v 【output】:%v 【expectedOutput】:%v\n", p.one, mergeSortedArray(p.one, p.two, p.three, p.four), a.one)
	}
	fmt.Printf("\n\n\n")
}
