package leetcode

import (
	"fmt"
	"testing"
)

func TestLeetCode(t *testing.T) {
	////points := [][]int{{-17, 5}, {-10, -8}, {-5, -13}, {-2, 7}, {8, -14}}
	////points := [][]int{{1, 3}, {2, 0}, {5, 10}, {6, -10}}
	//points := [][]int{{1, 100}, {2, 99}, {3, 101}, {500, 500}}
	////points := [][]int{{-19, -12}, {-13, -18}, {-12, 18}, {-11, -8}, {-8, 2}, {-7, 12}, {-5, 16}, {-3, 9}, {1, -7}, {5, -4}, {6, -20}, {10, 4}, {16, 4}, {19, -9}, {20, 19}}
	////points := [][]int{{-15, -1}, {-14, -5}, {-11, 1}, {-9, 7}, {-8, 18}, {-7, -5}, {-3, 3}, {4, 14}, {12, -4}, {13, 15}, {14, -19}, {19, -1}}
	//k := 498
	//fmt.Println(findMaxValueOfEquation4(points, k))

	a := []int{1, 2, 3, 4}
	m := map[int]int{}
	for i := 0; i < len(a); i++ {
		m[a[i]]++
	}
	if _, ok := m[2]; ok {
		fmt.Println("1")
	}
	m[2]--
	fmt.Println(m)
	if _, ok := m[5]; ok {
		fmt.Println("1")
	}
}
