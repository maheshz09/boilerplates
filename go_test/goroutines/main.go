package main

import (
	"fmt"
	"net/http"
	"sync"
)

var singnals = []string{"test"}
var wg sync.WaitGroup //pointers
var mut sync.Mutex    //pointer

func main() {
	// go greeter("hello")
	// greeter("world")
	websitelist := []string{
		"https://google.com",
		"https://go.dev",
		"https://fb.com",
		"https://github.com",
		"https://db.dev",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(singnals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println("hello ", s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("something is wrong with endpoint", endpoint)
	} else {
		mut.Lock()
		singnals = append(singnals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for website %s\n", res.StatusCode, endpoint)
	}
}
