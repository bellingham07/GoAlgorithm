package interview

import (
	"fmt"
	"testing"
)

func TestInterview(t *testing.T) {

	total := 0
	for i := 0; i < 101; i++ {
		if i%2 == 0 {
			total += i
		}
	}
	fmt.Println(total)
}
