package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	printFizzy := func(num int) {
		if num % 5 == 0 && num % 3 == 0 {
			fmt.Println("Fizzbuzz")
			return
		}
		if num % 3 == 0 {
			fmt.Println("Fizz")
			return
		}
		if num % 5 == 0 {
			fmt.Println("Buzz")
			return
		}
		fmt.Println(num)
	}

	if len(os.Args) < 2 {
		fmt.Println("No args buddy...")
		os.Exit(2)
	}
	input := os.Args[1]
	inputInt, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Houston, we have a problem.")
		os.Exit(1)
	}

	for i := 1; i <= inputInt; i++ {
		printFizzy(i)

	}

}
