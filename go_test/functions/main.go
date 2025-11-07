package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to functions")
	greeter()

	// function inside function is not allowed
	// func greeterTwo()  {
	//		fmt.Println("another method")
	//	}
	greeterTwo()

	result := adder(3, 5)
	fmt.Println("Result of addition is: ", result)
	proRes, myMessage := proAdder(2, 5, 8, 7)
	fmt.Println("pro result is ", proRes)
	fmt.Println("pro message is ", myMessage)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(vales ...int) (int, string) {
	total := 0

	for _, val := range vales {
		total += val
	}
	return total, "Hi Pro result funtion"
}

func greeter() {
	fmt.Println("hello from golang")
}

func greeterTwo() {
	fmt.Println("another method")
}
