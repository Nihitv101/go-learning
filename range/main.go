package main

import "fmt"

// iterating over data structure
func main() {

	// nums := []int{4, 5, 6}

	// index and character in string

	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}

	// using maps

	// m := map[string]string{"fname": "john", "lname": "doe"}

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// sum := 0

	// for i, num := range nums {
	// 	fmt.Println(num, i)
	// 	sum += num
	// }

	// fmt.Println(sum)

}
