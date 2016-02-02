// Exercise to determine the square root by newton rules
package main

import (
	"fmt"
)

func main() {
	var number, x float64 = 2.0, 10

	for i := 1; i < 10; i++ {
		x = newton_square(x, number)
	}

	fmt.Println(x)
}

func newton_square(x, number float64) float64 {
	v := (x * x - number) / (x * 2)
	return x - v
}
