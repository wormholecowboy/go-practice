package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readInput(input io.Reader) (words, lines, bytes int) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		text := scanner.Text()
		words += len(strings.Fields(text))
		bytes += len(text) + 1
		lines++
	}
	return
}

func main() {
	var f *os.File
	var filename string

	if len(os.Args) < 2 {
		f = os.Stdin
	} else {
		filename = os.Args[1]
		var err error
		f, err = os.Open(filename)
		if err != nil {
			log.Fatal("Not a file")
		}
		defer f.Close()

	}

	wordCount, lineCount, byteCount := readInput(f)
	fmt.Printf("%v %v %v %v \n", lineCount, wordCount, byteCount, filename)
}
