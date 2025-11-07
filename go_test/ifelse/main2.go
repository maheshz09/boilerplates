package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Starting...")

	err := doSomething()
	if err != nil {
		fmt.Println("Non-OK:", err)
	} else {
		fmt.Println("OK: operation successful")
	}
}

func doSomething() error {
	loginCount := 100

	if loginCount > 50 {
		return errors.New("too many logins, suspicious activity")
	}

	return nil
}
