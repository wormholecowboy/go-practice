package main

import (
	"fmt"
	"os"
	"strconv"
)

var count int

func loop(num int) int {
	if num <= 1 {
		return num
	}

	fmt.Println("Starting loop version")
	for i := num; i >= 1; i-- {

	}

	fmt.Println("Starting recusive version")
	recurResult := recur(num)
	fmt.Println("%d iterations", recurResult)

	return -1
}

func recur(num int) int {
	if num <= 1 {
		return num
	}
	return recur(num-1) + recur(num-2)
}

func main() {
	num, err:= strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Print("Nope")
		os.Exit(1)
	}
	fmt.Println("Fibbing %d", num)
	fib := recur(55)
	fmt.Printf("fib: %d", fib)


}
