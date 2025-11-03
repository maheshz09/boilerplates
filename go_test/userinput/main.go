package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	//we are getting output from user using bufio with os module of go
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for our pizza:")

	// whatever the reader reads i want to store that thing into variable
	// that whare the comma ok/ error ok sintax came in (in simple terms its a try/catch)
	// so it is something ither you will get input from user or error
	//we can able to define like that aswell input and error

	//input, error
	//OR _ for error
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating ", input)
	fmt.Printf("type of rating is %T", input) // its a string

	// notice one thing, when we are reading input, _  something from standurd input there chance to something gose wrong & for that wrong thing error might came up
	// for that we use store the error to another variable like this -- input, err := reader.ReadString('\n')
	// the _ is we can use it in both the places currently we dont much care about the error's so we are using _ . there is might the case we care more foucsing on the error might be on the loops so in that case we can also able to use
	// -- _, err := reader.ReadString('\n') so we can use it in multiple places

}
