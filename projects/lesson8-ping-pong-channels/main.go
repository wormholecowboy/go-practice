package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	ball := make(chan int)
	var wg sync.WaitGroup
	const defaultNum = 10
	num := defaultNum

	if len(os.Args) > 1 {
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(1)
		}
		num = n
	}

	wg.Add(2)

	player := func(playerName string) {
		defer wg.Done()
		for {
			count, ok := <- ball
			if !ok {
				return
			}
			fmt.Println(playerName, count)
			if count >= num {
				close(ball)
				return
			}
			ball <- count + 1
		}
	}

	go player("Ping")
	go player("Pong")

	ball <- 1
	wg.Wait()

}
