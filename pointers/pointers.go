package main

import "fmt"

// pointers - used to store the memory location of the variable

// passing by value
func changenum(num *int) {
	*num = 9
	fmt.Println("In changenum", *num)
}

func main() {

	num := 12
	changenum(&num)

	fmt.Println("After change num", num)

}
