package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}

func main() {

	// fmt.Println(11, 2, 3, 4, "ehe")
	// fmt.Println(sum(1, 2, 3, 4))

	result := []int{1, 2, 3, 4}

	fmt.Println(sum(result...))

}
