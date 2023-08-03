package main

import (
	"fmt"
	"strings"
)

func main() {
factResult := fact(5)
fmt.Println(factResult)
}

func fact (n int) int {
if n ==0 (return 1)
return n*fact(n -1)
}