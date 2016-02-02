package main

import (
	"fmt"
	"math"
)

func main() {
	i := 1
	sum := 0

	for ; i < 10; i++  {
		sum += i
	}

	fmt.Println(sum)
	fmt.Println(for_is_gos_while(100))

	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-4))

	fmt.Println(pow(3, 4, 100))
	fmt.Println(pow(4, 4, 100))
}

func for_is_gos_while(upper int) int {
	sum := 1

	for sum < upper  {
		sum += sum
	}

	return sum
}

func sqrt(x float64) string {
	if x < 0 {
		return fmt.Sprint(math.Sqrt(-x)) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		fmt.Printf("Since %g >= %g\n", v, limit)
	}

	return limit
}
