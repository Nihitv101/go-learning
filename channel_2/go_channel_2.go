package main

import (
	"fmt"
	"time"
)

// NORMAL CHANNELS - UNBUFFERED CHANNELS - RECEIVING / SENDING OF DATA IS BLOCKING
// IN Normal channels , its blocking so we have to wait until the data we ssend and receive

// BUFFERED CHANNEL - limited amount of data we can send without blocking

func emailSender(emailChan chan string, done chan bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("sending emails to ", email)
		time.Sleep(time.Second)
	}
}

func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)

	// inline goroutine

	go func() {
		chan1 <- 19
	}()

	go func() {
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1val := <-chan1:
			fmt.Println("received data from chan1", chan1val)
		case chan2val := <-chan2:
			fmt.Println("received data from chan2", chan2val)
		}

	}

	// we have two channels and we want to send the data to both the channel

	// the size we are giving  - that much data is non blocking , data more than size will be blocking
	// emailChan := make(chan string, 100) // non blocking

	// done := make(chan bool) // blocking

	// emailChan <- "nihit@gmail.com"
	// emailChan <- "verma@example.com"

	// for i := 0; i < 5; i++ {
	// 	emailChan <- fmt.Sprintf("nihit_%d@gmail.com", i)
	// }

	// fmt.Println("done sending")

	// go emailSender(emailChan, done)

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	// this is important
	// close(emailChan)

	// <-done // blocking the main goroutine until our process is completed

}
