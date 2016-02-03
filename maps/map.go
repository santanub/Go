package main

import "fmt"

type Vertex struct {
	X, Y int
}

//var m map[string]Vertex

func main() {
	m := make(map[string]Vertex)
	m["Santanu"] = Vertex{ 5, 6 }
	fmt.Println(m)
	fmt.Println(m["Santanu"])

	var n = map[string]Vertex {
		"Santanu": Vertex{ 1, 0 },
		"Prosen": Vertex{ 2, 3 },
	}

	n["Rajib"] = Vertex { 5, 6 }
	fmt.Println(n)

	// Retrieve an element
	fmt.Println(n["Santanu"])

	//Delete an element
	delete(n, "Prosen")
	fmt.Println(n["Prosen"])
	fmt.Println(n)

	elem, ok := n["Prosen"]
	fmt.Println(ok)
	fmt.Println(elem)

	elem, ok = n["Rajib"]
	fmt.Println(ok)
	fmt.Println(elem)
}
