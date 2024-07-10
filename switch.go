package main

import (
	"fmt"
	"time"
)

func main() {
	i := 3
	fmt.Printf("Write %d  as ", i)
	switch i {
	case 1:
		fmt.Print("one")
	case 2:
		fmt.Print("Two")
	case 3:
		fmt.Print("Three")

	}

	println()
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is the weekdend")
	default:
		fmt.Println(("It is a weekday"))
	}

}
