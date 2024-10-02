package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(w *sync.WaitGroup) {

	// defer w.Done()

	defer func() {
		w.Done()
		p.mu.Unlock()
	}()
	p.mu.Lock()
	p.views += 1

}

func main() {
	var wg sync.WaitGroup
	post1 := post{views: 0}

	// this is running concurrenlty , parallely
	// so i want that this goroutine should run and compleelty before the main function ends
	// i will use waitgroup

	// But here is the problem that multiple goroutine are modifying the views resource , thats why a race condition is being created
	// so we use the mutex for that

	for i := 0; i < 200; i++ {
		wg.Add(1)
		go post1.inc(&wg)
	}
	wg.Wait()

	fmt.Println(post1.views)

}
