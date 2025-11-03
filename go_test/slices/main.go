package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to video on slices")
	var fruitList = []string{"Apple", "peach", "tomato"}
	fmt.Printf("type of that fruitList %T\n", fruitList)

	fruitList = append(fruitList, "mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867
	// highScore[4] = 777
	highScore = append(highScore, 555, 666, 321)
	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println("highScore after sorting", highScore)

	fmt.Println("ints are sorted ;", sort.IntsAreSorted(highScore))

	// how to remove a value from slice based on index
	var course = []string{"reactjs", "javascript", "switft", "python", "ruby"}
	fmt.Println(course)
	var index int = 2
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course)

	var index2 int = 3
	course = append(course[:index2], course[index2+1:]...)
	fmt.Println(course)
}
