package main

import "fmt"

func main() {
	// we can apply the switch case

	whoamI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Integer type")
		case bool:
			fmt.Println("Bool type")
		case string:
			fmt.Println("String type")
		default:
			fmt.Println("other", t)
		}
	}

	whoamI(true)

	// SWITCH STATEMENT IN GO

	// age := 28

	// switch age {
	// case 18:
	// 	fmt.Println("Person is 18")
	// case 28:
	// 	fmt.Println("Person is 28")

	// case 38:
	// 	fmt.Println("Person is 38")
	// default:
	// 	fmt.Println("other")
	// }

	// if else
	// age := 88

	// if age >= 0 && age < 18 {
	// 	fmt.Println("below 18")
	// } else if age >= 18 && age < 40 {
	// 	fmt.Println("person is not adult")
	// } else {
	// 	fmt.Println("Personn is old")
	// }

}
