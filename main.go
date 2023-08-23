package main

import "fmt"

func counter(out chan<- int) {
	println("111111111")
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	println("22222222")
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	println("33333333333333")
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
