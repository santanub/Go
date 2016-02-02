package main

import(
	"fmt"
	"math"
	"math/rand"
)

func add(x , y int) int {
	return x + y
}

func split(sum int) (x, y int) {
	x = sum/4
	y = sum+4
	return
}

func main() {
	fmt.Println("Your lucky number is", rand.Intn(10))
	fmt.Println("Your lucky number is", rand.Seed)
	fmt.Println(math.Pi)
	fmt.Println(add(4,3))
	fmt.Println(split(50))
}

