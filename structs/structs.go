package main

import (
	"fmt"
	"time"
)

// ecommerce order struct

type customer struct {
	name  string
	phone string
}

type orders struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosec precision
	customer
}

// we dont have directly constructors in the golang , but we try to use the function for it

// func newOrders(id string, amount float32, status string) *orders {
// 	myOrder := orders{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}

// 	return &myOrder
// }

// // receiver type
// func (o *orders) changeStatus(status string) {
// 	o.status = status
// }

// func (o *orders) getamount() float32 {
// 	return o.amount
// }

func main() {

	customer1 := customer{
		name:  "nihit",
		phone: "8921",
	}

	order1 := orders{
		id:       "1",
		amount:   89,
		status:   "returned",
		customer: customer1,
	}

	fmt.Println(order1)

	// Struct embedding

	// creatint the instance using constructor which will return the pointer to the object created

	// order1 := newOrders("1", 89.22, "received")
	// fmt.Println(order1)

	// language := struct {
	// 	name   string
	// 	isgood bool
	// }{"golang", true}

	// fmt.Println(language.name, language.isgood)

	// // creating the instance of the struct

	// order1 := orders{
	// 	id:     "1",
	// 	amount: 49.2,
	// 	status: "received",
	// }

	// fmt.Println(order1)

	// order1.changeStatus("returned")

	// fmt.Println(order1.getamount())

	// fmt.Println(order1)

	// order1.createdAt = time.Now()

	// fmt.Println(order1)

	// // getting the field

	// fmt.Println(order1.amount)

}
