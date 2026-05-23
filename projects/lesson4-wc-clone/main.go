package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func readInput(input io.Reader) (int, int, int) {
	var wordCount, byteCount, lineCount int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		text := scanner.Text()

		textStrings := strings.Fields(text)
		wordCount += len(textStrings)

		byteCount += len(text) + 1

		lineCount++

	}
	return wordCount, lineCount, byteCount
}

func main() {
	var reader io.Reader

	if len(os.Args) < 2 {
		reader = os.Stdin
	}

	var input string
	if os.Args[1] != "" {
		input = os.Args[1]
	}

	if _, err := os.Stat(input); err == nil {
		var openErr error
		reader, openErr = os.Open(input)
		defer reader.Close()
		if openErr != nil {
			fmt.Print("File does not exist or could not be opened. \n")
			os.Exit(1)
		}
	} else {
		reader = os.Args[1]
	}

	wordCount, lineCount, byteCount := readInput(reader)
	fmt.Printf("%v %v %v %v \n", lineCount, wordCount, byteCount, os.Args[1])
}
