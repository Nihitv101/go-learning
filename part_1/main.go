package main

import "fmt"

func main() {

	// new thing in go

	for i := range 4 {
		fmt.Println(i)
	}

	// // classic for loop
	// for i := 0; i < 3; i++ {
	// 	if i == 1 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// 	// break

	// }

	// there is not while loop in go , so use only for to make while
	// fmt.Println("hello")
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

}
