package main

import "fmt"

func main() {
	i := 1
	sum := 0

	for ; i < 10; i++  {
		sum += i
	}

	fmt.Println(sum)
}
