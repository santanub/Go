package main

import "fmt"

func fibonac() func(int, int) (int, int) {
	return func(x, y int) (int, int) {
		return y, (x + y)
	}
}

func main() {
	fib := fibonac()
	x := 0
	y := 1

	fmt.Println(x)
	fmt.Println(y)

	for i :=0; i < 10; i++ {
		x, y = fib(x , y)
		fmt.Println(y)
	}
}
