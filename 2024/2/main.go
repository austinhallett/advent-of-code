// --- Day 2: Red-Nosed Reports ---
// Fortunately, the first location The Historians want to search isn't a long walk from the Chief Historian's office.

// While the Red-Nosed Reindeer nuclear fusion/fission plant appears to contain no sign of the Chief Historian, the engineers there run up to you as soon as they see you. Apparently, they still talk about the time Rudolph was saved through molecular synthesis from a single electron.

// They're quick to add that - since you're already here - they'd really appreciate your help analyzing some unusual data from the Red-Nosed reactor. You turn to check if The Historians are waiting for you, but they seem to have already divided into groups that are currently searching every corner of the facility. You offer to help with the unusual data.

// The unusual data (your puzzle input) consists of many reports, one report per line. Each report is a list of numbers called levels that are separated by spaces. For example:

// 7 6 4 2 1
// 1 2 7 8 9
// 9 7 6 2 1
// 1 3 2 4 5
// 8 6 4 4 1
// 1 3 6 7 9
// This example data contains six reports each containing five levels.

// The engineers are trying to figure out which reports are safe. The Red-Nosed reactor safety systems can only tolerate levels that are either gradually increasing or gradually decreasing. So, a report only counts as safe if both of the following are true:

// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.
// In the example above, the reports can be found safe or unsafe by checking those rules:

// 7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
// 1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
// 9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
// 1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
// 8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
// 1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.
// So, in this example, 2 reports are safe.

// Analyze the unusual data from the engineers. How many reports are safe?

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Report struct {
	Levels []int
}

type numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Abs[T numeric](val T) T {
	if val < 0 {
		return -val
	}
	return val
}

func (r *Report) IsSafe() bool {
	increasing := r.Levels[1] > r.Levels[0]
	for i, _ := range r.Levels {
		if i == 0 {
			continue
		}
		diff := Abs(r.Levels[i] - r.Levels[i-1])
		isSequential := (increasing && r.Levels[i] > r.Levels[i-1]) || (!increasing && r.Levels[i] < r.Levels[i-1])
		validDiff := 1 <= diff && diff <= 3
		if !isSequential || !validDiff {
			return false
		}
	}
	return true
}

func stringToIntSlice(input string) ([]int, error) {
	// Split the string into a slice of substrings
	parts := strings.Fields(input)

	// Create a slice to hold the integers
	intSlice := make([]int, len(parts))

	// Convert each substring to an integer
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, err // Return an error if conversion fails
		}
		intSlice[i] = num
	}
	return intSlice, nil
}

func GetInput() ([]Report, error) {
	// typically run from the year directory
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reports []Report

	// Use bufio.Scanner for line-by-line reading
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Trim spaces and split if numbers are space-separated
		line := strings.TrimSpace(scanner.Text())
		levels, err := stringToIntSlice(line) // Splits the line into parts by spaces
		if err != nil {
			fmt.Printf("Skipping invalid number: %s\n", line)
			continue
		}
		reports = append(reports, Report{Levels: levels})
	}

	// Handle potential scanner errors
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
		return nil, err
	}
	return reports, nil
}

func main() {
	reports, err := GetInput()
	if err != nil {
		log.Fatalf("GetInput: %v", err)
	}

	safeReports := 0
	for _, report := range reports {
		if report.IsSafe() {
			safeReports++
		}
	}

	fmt.Printf("Safe reports: %d\n", safeReports)
}
