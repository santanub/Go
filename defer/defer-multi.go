package main

import "fmt"
// defer will stop execution of a func and execute the surrounding functions, and it will store it in a stack and execute it LIFO order.
func main() {
	fmt.Println("Counting")

	for i:= 0; i< 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done..")
}
