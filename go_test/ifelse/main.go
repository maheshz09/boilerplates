package main

import "fmt"

func main() {
	fmt.Println("this is if-else statement")
	loginCount := 100
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "watchout"
	} else {
		result = "if its exactly 10"
	}

	// even or odd
	if 19%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("The number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is NOT less than 10")
	}
	fmt.Println(result)
}
