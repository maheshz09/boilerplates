package main

import "fmt"

func main() {
	fmt.Println("struct in golang")
	// there no inheritance in golang; no super or parent no child
	mahesh := User{"Mahesh", "mahesh.oncloudev.com", true, 25}
	fmt.Println(mahesh)
	fmt.Printf("mahesh details are: %+v\n", mahesh)
	fmt.Printf("Name is %v and email is %v\n", mahesh.Name, mahesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
