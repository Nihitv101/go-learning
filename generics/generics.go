package main

import "fmt"

func printNameSlice(items []string) {
	for _, val := range items {
		fmt.Println(val)
	}
}

func printNumSlice(nums []int) {
	for _, val := range nums {
		fmt.Println(val)
	}
}

// using generics

func printData[T any](data []T) {
	for _, val := range data {
		fmt.Println(val)
	}
}

func main() {

	names := []string{"golang", "javascript", "typescript"}
	nums := []int{1, 2, 3, 4}
	vals := []bool{true, false, true, false}

	printData(names)
	printData(nums)
	printData(vals)

	// fmt.Println(names)

	// printNameSlice(names)
	// printNumSlice(nums)

}
