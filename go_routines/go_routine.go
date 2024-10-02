package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {

	// defer will run this in the last of the function
	defer w.Done()
	fmt.Println("Doing task", id)

}

// multi threading using go
// we can achieve parellism and power of multi threadig using go routines
// this is something like asynchronous mode of javascript

// The problem with go routines is that it has seperate scheduler for it and it runs fast , before it our main program executes and its  removed from the call stack
// so to resolve this we use the waitgroups
// waitGroup has 3 methods - Done , Add, Wait
// Add - when we use the go routine add is used
// Done - when the go routine is added it needs to be decremented
// Wait - wait until go routine is completely finished

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {

		wg.Add(1)
		go task(i, &wg)
	}

	fmt.Println("Running")

	// time.Sleep(time.Second * 2)

	wg.Wait()

}
