package main

import "fmt"

// There is not a concept of builtin enum in go, we create our own enum using struct type

// type OrderStatus int

type OrderStatus string

const (
	// Received OrderStatus = iota
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to ", status)
}

func main() {
	changeOrderStatus(Delivered)
}
