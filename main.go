package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(a)
	fmt.Println(min(1, 2))
	fmt.Println(max(11, 2))
	clear(a)
	fmt.Println(a)
}
