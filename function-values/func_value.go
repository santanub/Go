package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func computeStr(fn func(string) string) string {
	return fn("bad")
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	print := func(x string) string {
		return "Santanu is " + x
	}

	fmt.Println(hypot(3, 4));

	fmt.Println(compute(hypot));
	fmt.Println(print("good"));
	fmt.Println(compute(math.Pow));
	fmt.Println(computeStr(print));
}
