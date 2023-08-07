package stack_and_queue

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	a := []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}
	fmt.Println(topKFrequent(a, 10))
}
