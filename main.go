package main

import "fmt"

func main() {
	// Array
	//arr := [...]string{"Belka", "Up", "Down", "Break", "Absolut", "Drug"}
	// Full Format
	//for i := 0; i < 6; i++ {

	//	fmt.Println(arr[i])

	//}
	// Alternative Format
	//i := 0
	//for i < 6 {
	//		fmt.Println(arr[i])
	//	i++
	//}

	// Infinite Loop
	//for {
	//	fmt.Println("Hello World")
	//}

	// Sum of numbers
	//arr := [...]int{10, 15, 51, 95, 47, 84, 25, 63, 8, 74, 1, 952, 98}
	//sum := 0
	//for i := 0; i < len(arr); i++ {
	//	sum += arr[i]
	//}
	//fmt.Println(sum)

	//for i := 0; i < 10; i++ {
	//
	//	if i == 5 {

	//		continue

	//	}
	//	fmt.Println(i)
	//}

	// Range
	//arr := [...]int{10, 15, 51, 95, 47, 84, 25, 63, 8, 74, 1, 952, 98}
	//	for i, val := range arr {
	//
	//	fmt.Println(i, val)
	//
	//}
	// Slice
	contacts := []string{"add", "see", "more"}
	contacts2 := contacts[1:]

	fmt.Println(contacts2)

}
