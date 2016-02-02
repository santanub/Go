package main

import "fmt"

func main() {
	s := []int{2,3,4,5,6,7,78,8,6,4}

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}
}
