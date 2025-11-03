package main

import "fmt"

func main() {
	fmt.Println("welcome to class on pointers")
	// var ptr *int
	// fmt.Println("value of ptr is:", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("creaiting ptr for refrencing some number/memory:", ptr)  // pointer is a refrance to actual memory location
	fmt.Println("creaiting ptr for refrencing some number/memory:", *ptr) // but the value inside this pointer is actually 23

	*ptr = *ptr + 2
	fmt.Println("what do you think it will update the exesting value witch is 23 ???", myNumber)

}
