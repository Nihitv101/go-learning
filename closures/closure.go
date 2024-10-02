package main

import "fmt"

// counter function will return the function func()
// and func() will return the int

func counter() func() int {
	var count int = 0
	return func() int {
		count += 1
		return count
	}
}

func main() {

	ans := counter()
	fmt.Println(ans())

}
