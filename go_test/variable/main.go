package main

import "fmt"

// logic of go
// the first keyword that we have mentioned its capital Login the L is capital it means it public function its equalant to add public in front of variable/function.
// so this Login token is able to accessible by any other file into this folder or actually in this program to use it anywhare.
const LoginToken string = "fuck" // public statement

// it will not work heere
//jwtToken  := 300000

func main() {
	var username string = "mahesh"
	fmt.Println(username)
	fmt.Printf("varibe is type %T \n", username)

	var isLoggedIn bool = true //also false
	fmt.Println(isLoggedIn)
	fmt.Printf("varibe is type %T \n", isLoggedIn)

	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("varibe is type %T \n", smallVal)

	var smallFloat float64 = 255.45544545454
	fmt.Println(smallFloat)
	fmt.Printf("varibe is type %T \n", smallFloat)

	var anotherVariable int
	fmt.Println("varibe is type %T \n", anotherVariable)

	//implicit type without defining the datatype
	//lexer is doing the nessary things behind the scene
	var website = "oncloudev.com"
	fmt.Println(website)

	// no var type
	// it calls as vallurus operator, but inside method only, if i declare same thing on outside that would be a problem it will trow an error.
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	// calling that public thing
	fmt.Println(LoginToken) // it will work beacuse of global aswell as public
	fmt.Println("varibe is type %T \n", LoginToken)
}
