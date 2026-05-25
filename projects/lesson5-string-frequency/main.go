package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No arg provided")
	}

	fn := os.Args[1]
	f, err := os.Open(fn)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer f.Close()

	m := make(map[string]int)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		for _, word := range words {
			formatted := strings.Trim(strings.ToLower(word), " .,")
			if len(formatted) == 0 {
				continue
			}
			m[formatted]++
		}
	}

	type pair struct {
		word  string
		count int
	}
	var seenWords []pair

	for k, v := range m {
		s := pair{k, v}
		seenWords = append(seenWords, s)
	}

	slices.SortFunc(seenWords, func(a, b pair) int {
		// return b.count - a.count
		return cmp.Compare(b.count, a.count)

	})

	topTen := seenWords[:min(10, len(seenWords))]
	for _, item := range topTen {
		fmt.Printf("%-15s %d\n", item.word, item.count)
	}

}
