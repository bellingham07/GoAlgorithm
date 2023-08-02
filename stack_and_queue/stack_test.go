package stack_and_queue

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := "abbaca"
	duplicates := removeDuplicates(s)
	fmt.Println(duplicates)
}
