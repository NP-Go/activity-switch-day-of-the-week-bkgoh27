package main

import (
	"fmt"
	"time"
)

func main() {
	// switch ch := "b"; ch {
	// case "a":
	// 	fmt.Println("a")
	// case "b", "c":
	// 	fmt.Println("b or c")
	// default:
	// 	fmt.Println("No matching character")
	// }
	switch day := time.Now().Weekday(); day {
	case time.Monday, time.Wednesday, time.Friday, time.Sunday:
		fmt.Println(day, "- odd day")
	case time.Tuesday, time.Thursday, time.Saturday:
		fmt.Println(day, "- even day")
	default:
		fmt.Println("what is this day? ", day)
	}
}
