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

// --- Part Two ---
// The engineers are surprised by the low number of safe reports until they realize they forgot to tell you about the Problem Dampener.

// The Problem Dampener is a reactor-mounted module that lets the reactor safety systems tolerate a single bad level in what would otherwise be a safe report. It's like the bad level never happened!

// Now, the same rules apply as before, except if removing a single level from an unsafe report would make it safe, the report instead counts as safe.

// More of the above example's reports are now safe:

// 7 6 4 2 1: Safe without removing any level.
// 1 2 7 8 9: Unsafe regardless of which level is removed.
// 9 7 6 2 1: Unsafe regardless of which level is removed.
// 1 3 2 4 5: Safe by removing the second level, 3.
// 8 6 4 4 1: Safe by removing the third level, 4.
// 1 3 6 7 9: Safe without removing any level.
// Thanks to the Problem Dampener, 4 reports are actually safe!

// Update your analysis by handling situations where the Problem Dampener can remove a single level from unsafe reports. How many reports are now safe?

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/austinhallett/advent-of-code/2024/utils"
)

type Report struct {
	Levels []int
}

func (r *Report) IsSafe(dampener bool) bool {

	if !dampener && unsafeIdx(r.Levels) == -1 {
		return true
	} else if dampener {
		for i := 0; i < len(r.Levels); i++ {
			if unsafeIdx(deleteLevelAt(i, r.Levels)) == -1 {
				return true
			}
		}
	}
	return false
}

func deleteLevelAt(idx int, levels []int) []int {
	deleted := make([]int, len(levels)-1)
	copy(deleted[:idx], levels[:idx])
	copy(deleted[idx:], levels[idx+1:])
	return deleted
}

func unsafeIdx(levels []int) int {
	if len(levels) <= 1 {
		return -1
	}
	increasing := levels[1] > levels[0]
	for i := 1; i < len(levels); i++ {
		diff := utils.Abs(levels[i] - levels[i-1])
		isSequential := (increasing && levels[i] > levels[i-1]) || (!increasing && levels[i] < levels[i-1])
		validDiff := 1 <= diff && diff <= 3
		if !isSequential || !validDiff {
			return i
		}
	}
	return -1
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

func part1(reports *[]Report) int {
	safeReports := 0
	useDampener := false
	for _, report := range *reports {
		if report.IsSafe(useDampener) {
			safeReports++
		}
	}
	return safeReports
}

func part2(reports *[]Report) int {
	useDampener := true
	safeReports := 0

	for _, report := range *reports {
		if report.IsSafe(useDampener) {
			safeReports++
		}
	}
	return safeReports
}

func main() {
	reports, err := GetInput()
	if err != nil {
		log.Fatalf("GetInput: %v", err)
	}

	fmt.Println("Part 1")
	safeReports := part1(&reports)
	fmt.Printf("Safe reports: %d\n", safeReports)

	fmt.Println("Part 2")
	safeReports = part2(&reports)
	fmt.Printf("Safe reports: %d\n", safeReports)
}
