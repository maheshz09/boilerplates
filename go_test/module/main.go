package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello from golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Golang Series</h1>"))
}

// go mod commands
//  2008  go mod init github.com/maheshz09/boilerplates/go_test/
//  2009  go mod init github.com/maheshz09/boilerplates/go_test
//  2010  go get -u github.com/gorilla/mux
//  2011  go env
//  2014  go build
//  2016  go build .
//  2017  go run main.go
//  2019  go mod verify
//  2020  go list
//  2021  go list all
//  2022  go list all -m
//  2023  go list -m all
//  2024  go list -m -versions github.com/gorilla/mux
//  2025  go mod tidy
//  2026  go mod why
//  2027  go mod why github.com/gorilla/mux
//  2028  go mod graph
//  2032  go mod edit
//  2034  go mod vendor
