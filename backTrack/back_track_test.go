package backTrack

import (
	"testing"
)

var (
	a1   []int
	res1 [][]int
)

func TestBackTrack(t *testing.T) {
	combine(7, 3)
}

func tt() {
	for i := 0; i < 5; i++ {
		a1 = append(a1, i)
	}
	res1 = append(res1, a1)
}
