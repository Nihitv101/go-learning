package main

import "fmt"

// Its like a pipe , we send the data from one end and receive in another end
// Channel - communication of data between the go routines

// Sending the data from one goroutine to another goroutine
// func processNum(numChan chan int) {

// 	for num := range numChan {
// 		fmt.Println("processing number", num)
// 	}
// 	// fmt.Println("processing number", <-numChan)
// }

// func sum(result chan int, num1 int, num2 int) {
// 	answer := num1 + num2
// 	// sending this result to another goroutine
// 	result <- answer

// }

// WE USED THE WAITGROUP FOR THE SYNCHRONISATON OF THE GOROUTINES
// It allows that all our goroutines are removed and then remainnig the main go routine is left at last
// we can also perform the same using the channel

func task(done chan bool) {

	// defer function - cleaning function - wehen the task function exexutes and completed , defer function will run
	defer func() { done <- true }()
	fmt.Println("processsing task ...")

}

func main() {

	// done := make(chan bool)

	// go task(done)

	// val := <-done // block

	// // what is blocking in channel -> sending and receiving is blocking in channel

	// fmt.Print("received  ", val)

	// resultChan := make(chan int)

	// go sum(resultChan, 3, 4)

	// res := <-resultChan

	// fmt.Println("result ", res)

	// var wg sync.WaitGroup

	// numChan := make(chan int)

	// // wg.Add(1)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(20)
	// 	// wg.Wait()
	// }

	// numChan <- 4

	// THIS IS GIVING ERROR BECAUSE THE CHANNEL IS BLOCKING
	// Creating channel
	// messageChan := make(chan string) // Blocking

	// messageChan <- "ping"

	// msg := <-messageChan

	// fmt.Println(msg)
}
