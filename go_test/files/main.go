package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to files in golang")
	content := "this needs to go in file - oncloudev.com"

	file, err := os.Create("./mylcofile.txt")
	// check and fail fast if create failed
	checkNilErr(err)
	// ensure file is closed when main returns
	defer file.Close()

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is:", length)

	readfile("./mylcofile.txt")

}

func readfile(filname string) {
	databyte, err := ioutil.ReadFile(filname)
	// insted of this
	// if err != nil {
	// 	panic(err)
	// }

	// new one
	checkNilErr(err)
	fmt.Println("text data inside the file is \n", string(databyte))
}

// insted of writing nill errors we can use a single function to do that
// for now everyting we are checking for nill sintax

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
