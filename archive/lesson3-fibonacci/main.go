package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func loop(num int) int {
	first := 0
	second := 1
	for i := 0; i < num; i++ {
		first, second = second, second+first
	}
	return first
}

func recur(num int) int {
	if num <= 1 {
		return num
	}
	return recur(num-1) + recur(num-2)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Need arg")
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Print("Nope")
		os.Exit(1)
	}

	fibStart := time.Now()
	fib := recur(num)
	fibEnd := time.Now()
	fmt.Printf("fib: %d, time: %v \n", fib, fibEnd.Sub(fibStart))
	loopStart := time.Now()
	lp := loop(num)
	loopEnd := time.Now()
	fmt.Printf("loop: %d, time: %v \n", lp, loopEnd.Sub(loopStart))

}
