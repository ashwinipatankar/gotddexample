package main

import (
	"fmt"
)


func Sum(x int, y int) int {
	fmt.Println("Adding two numbers")
	return x+y
}

func main() {
	Sum(10, 3)
}