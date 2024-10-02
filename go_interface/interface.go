package main

import "fmt"

type payment struct{}

func (p payment) makePayment(amount float32) {
	// creating the razorpay instance
	razorpayPaymentGw := razorpay{}
	razorpayPaymentGw.pay(amount)
}

type razorpay struct{}

func (p razorpay) pay(amount float32) {

	fmt.Println("Making payment using razorpay", amount)

}

func main() {
	// Creating the insance of payment struct
	newPayment := payment{}
	newPayment.makePayment(10)
}
