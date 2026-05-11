package main

import (
	"fmt"
)

func main() {

	printFizzy := func(num int) int {
	fmt.Println("Fizzbuzz")
		return num
	}

	var num int

	printFizzy(num)
}
