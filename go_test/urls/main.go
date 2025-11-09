package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://learncodeonlineclone.netlify.app:3000/learn?coursename=reactjs&paymentid=gjojiofds"

func main() {
	fmt.Println("welcome to handling urls in golang")
	fmt.Println(myurl)
	// parsing url
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)
	// key values
	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println("Params is:", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev:3000",
		Path:    "/learn",
		RawPath: "course=reactjs",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
