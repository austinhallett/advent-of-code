package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func getWordSearch() ([]string, error) {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file's contents
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(content), "\n"), nil
}

func main() {
	wordSearch, err := getWordSearch()
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	occurances := 0
	for li, line := range wordSearch {
		for ci, _ := range line {

			// horizontal occurances
			if ci+3 < len(line) && (line[ci:ci+4] == "XMAS" || line[ci:ci+4] == "SAMX") {
				occurances++
			}

			// vertical occurances
			if li+3 < len(wordSearch) && (wordSearch[li][ci] == 'X' && wordSearch[li+1][ci] == 'M' && wordSearch[li+2][ci] == 'A' && wordSearch[li+3][ci] == 'S') {
				occurances++
			}
			if li+3 < len(wordSearch) && (wordSearch[li][ci] == 'S' && wordSearch[li+1][ci] == 'A' && wordSearch[li+2][ci] == 'M' && wordSearch[li+3][ci] == 'X') {
				occurances++
			}

			// diagonal occurances
			if li+3 < len(wordSearch) && ci+3 < len(line) && (wordSearch[li][ci] == 'X' && wordSearch[li+1][ci+1] == 'M' && wordSearch[li+2][ci+2] == 'A' && wordSearch[li+3][ci+3] == 'S') {
				occurances++
			}
			if li+3 < len(wordSearch) && ci+3 < len(line) && (wordSearch[li][ci] == 'S' && wordSearch[li+1][ci+1] == 'A' && wordSearch[li+2][ci+2] == 'M' && wordSearch[li+3][ci+3] == 'X') {
				occurances++
			}
			if li-3 >= 0 && ci+3 < len(line) && (wordSearch[li][ci] == 'X' && wordSearch[li-1][ci+1] == 'M' && wordSearch[li-2][ci+2] == 'A' && wordSearch[li-3][ci+3] == 'S') {
				occurances++
			}
			if li-3 >= 0 && ci+3 < len(line) && (wordSearch[li][ci] == 'S' && wordSearch[li-1][ci+1] == 'A' && wordSearch[li-2][ci+2] == 'M' && wordSearch[li-3][ci+3] == 'X') {
				occurances++
			}
		}
	}
	fmt.Printf("The word 'XMAS' appears %d times in the word search.\n", occurances)
}
