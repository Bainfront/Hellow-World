package main

import (
	"fmt"
	"math"
)

func main() {
	var m, h float64
	fmt.Println("Type mass in kg: ")
	fmt.Scanf("%f", &m)

	fmt.Println("Type height in meters: ")
	fmt.Scanf("%f", &h)
	B := m / math.Pow(h, 2)
	fmt.Println("Your BMI is:", B)
}
