package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("this is all about channels in golang")
	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}
	// fmt.Println(myCh)
	// myCh <- 5
	wg.Add(2)

	// receve only cahnnel because of we have added <-
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myCh

		fmt.Println("we pushed 0 value to channel, it is showing it's default 0 value or our 0 value which we have pushed (if true means our value):", isChannelOpen)
		fmt.Println(val)
		defer wg.Done()
		//fmt.Println(<-myCh)
		// fmt.Println(<-myCh)
	}(myCh, wg)

	// into the channel that means we are sending somedata into channels
	go func(ch chan<- int, w *sync.WaitGroup) {
		defer wg.Done()
		myCh <- 0
		close(myCh)

		// myCh <- 5
		// myCh <- 6
	}(myCh, wg)
	wg.Wait()
}
