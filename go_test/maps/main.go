package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	fmt.Println("list of all languages:", languages)
	fmt.Println("JS stands for:", languages["JS"])

	delete(languages, "RB") // this delting method also works in slices
	fmt.Println("list of all languages:", languages)

	// loops in go
	for key, value := range languages {
		fmt.Println("For key %v, value is ")
	}

}
