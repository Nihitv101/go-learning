package main

import (
	"fmt"
	"maps"
)

func main() {

	m1 := map[string]int{"price": 80, "products": 9}
	m2 := map[string]int{"price": 80, "products": 9}

	fmt.Println(maps.Equal(m1, m2))

	// fmt.Println(m["price"])

	// v, ok := m["price"]

	// fmt.Println(ok, v)

	// m := make(map[string]int)
	// m["age"] = 19
	// m["price"] = 1000

	// fmt.Println(m["age"])

	// fmt.Println(len(m))

	// delete(m, "price")
	// fmt.Println(len(m))

	// clear(m)

	// maps
	// m := make(map[string]string)

	// m["language"] = "go lang"
	// m["area"] = "backend"

	// fmt.Println(m["language"], m["area"])

	// fmt.Println("hello")
}
