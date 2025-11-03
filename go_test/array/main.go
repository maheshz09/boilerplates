package main

import "fmt"

func main() {
	fmt.Println("welcome to array in golang")
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is:", fruitList)
	fmt.Println("Fruit list is:", len(fruitList))

	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println("vegi list is:", vegList)
	fmt.Println("vegi list is:", len(vegList))
}
