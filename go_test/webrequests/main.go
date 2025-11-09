package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://learncodeonlineclone.netlify.app/"

func main() {
	fmt.Println("webrequests")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close() // its caller's responsibility to close the connection
	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// StatusCode is a field (int) on http.Response — it does not return an error.
	// Use it directly.
	statusCode := response.StatusCode
	statusText := http.StatusText(statusCode)
	fmt.Println("the status code of the webpage is:", statusCode, statusText)

	content := string(databytes)
	fmt.Println(content)
}
