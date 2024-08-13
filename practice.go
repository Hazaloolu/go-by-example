package main

import (
	"fmt"
)

// func sum(nums ...int) int {
// 	total := 0

// 	for _, num := range nums {
// 		total += num

// 	}
// 	return total
// }

// closures

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {

	// swicth statement
	// i := 3
	// fmt.Printf("Write %d  as ", i)
	// switch i {
	// case 1:
	// 	fmt.Print("one")
	// case 2:
	// 	fmt.Print("Two")
	// case 3:
	// 	fmt.Print("Three")

	// }

	// println()
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It is the weekdend")
	// default:
	// 	fmt.Println(("It is a weekday"))
	// }

	// VARIADIC FUNCTION

	// result1 := sum(1, 2, 3)
	// result2 := sum(1, 2, 3, 4, 5)

	// fmt.Println(result1)
	// fmt.Println(result2)

	// CLOSURES

	increment := counter()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

}
