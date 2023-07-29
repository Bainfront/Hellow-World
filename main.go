package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	grade := random.Intn(5) + 10

	// if grade == 1 {
	// 	fmt.Println("Fall!")
	// } else if grade == 2 {
	// 	fmt.Println("Bad!")
	// } else if grade == 3 {
	// 	fmt.Println("Good!")
	// } else if grade == 4 {
	// 	fmt.Println("Nice!")
	// } else if grade == 5 {
	// 	fmt.Println("Excelent!")
	// }
	// switch
	switch grade {
	case 1:
		fmt.Println("Fall!")
	case 2:
		fmt.Println("Bad!")
	case 3:
		fmt.Println("Good!")
	case 4:
		fmt.Println("Nice!")
	case 5:
		fmt.Println("Excelent!")
	default:
		fmt.Println("NONCENCE")
	}

}
