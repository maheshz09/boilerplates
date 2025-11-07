package main

import "fmt"

func main() {
	fmt.Println("struct in golang")
	// there no inheritance in golang; no super or parent no child
	mahesh := User{"Mahesh", "mahesh.oncloudev.com", true, 25}
	fmt.Println(mahesh)
	fmt.Printf("mahesh details are: %+v\n", mahesh)
	fmt.Printf("Name is %v and email is %v.\n", mahesh.Name, mahesh.Email)
	mahesh.GetStatus()
	mahesh.NewMail()
	// actually vlaue go changed???
	fmt.Printf("Name is %v and email is %v.\n", mahesh.Name, mahesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("New email of this user is:", u.Email)
}
