package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to loops in golang")
	days := []string{"Sunday", "Tuesday", "Wedensday", "Friday", "Saturday"}
	fmt.Println(days)

	//	for d := 0; d < len(days); d++ {
	//		fmt.Println(days[d])
	//	}

	// for i := range days {
	//	 	fmt.Println(days[i])
	//}

	// for index, day := range days {
	//	fmt.Println("index is %V and value is %v\n", index, day)
	// }

	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
			rougueValue++
			continue
		}

		if rougueValue == 5 {
			rougueValue++
			fmt.Println("continue\n")
			continue
		}

		fmt.Println("Value is:", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("learn code online")
}
