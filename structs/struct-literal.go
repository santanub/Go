package main

import "fmt"

type Santanu struct {
	R, N int
}

var (
	s1 = Santanu { R: 1 }
	s2 = Santanu { N: 2 }
	s3 = Santanu { R: 1, N: 2 }
	s4 = Santanu { 4, 5 }
	p = &Santanu{3, 4}
)

var pow = []int{ 1, 2, 4, 8, 16, 32 }

func main() {
	fmt.Println(s1, s2, s3, s4, *p)

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for i := range pow {
		fmt.Printf("2**%d = ?\n", i)
	}

	for _, v := range pow {
		fmt.Printf("2**? = %d\n", v)
	}
}
