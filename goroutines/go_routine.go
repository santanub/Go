package main

import "fmt"

func f(from string) {
	for i:=0; i < 10; i++ {
		fmt.Println(from, ": ", i)
	}
}

func main() {
	f("direct")

	go f("go routine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	go f("go ing routine")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
