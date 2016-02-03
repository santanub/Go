package main

import "fmt"

func main() {
	s := []int{2,3,4,5,6,7,78,8,6,4}

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}

	a := make([]int, 4)
	fmt.Printf("%d\n", len(a))
	fmt.Printf("%d\n", cap(a))

	b := make([]int, 4, 5)
	fmt.Printf("%d\n", len(b))
	fmt.Printf("%d\n", cap(b))

	b = b[2:3]
	fmt.Printf("%d\n", len(b))
	fmt.Printf("%d\n", cap(b))
	fmt.Printf("%d\n", b)
}
