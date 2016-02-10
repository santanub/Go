package main

import(
	"fmt"
	"strings"
)

func determine_count(array []string, elem string) int {
	kount := 0

	for i:=0; i < len(array); i++ {
		if elem == array[i] {
			kount += 1
		}
	}

	return kount
}

func main() {
	input_string := "Santanu Prosen Santanu Santanu Santanu Prosen Santanu Prosen Ankit Aniket Aniket Halka Rajib Santanu"
	m := strings.Fields(input_string)
	x := make(map[string]int)

	for i :=0; i < len(m); i++ {
		_, ok := x[m[i]]

		if ok == false {
			x[m[i]] = determine_count(m, m[i])
		}
	}

	fmt.Println(x)
}
