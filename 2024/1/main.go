package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func OpenFile(identifier string) ([]int, error) {
	// typically run from the year directory
	fileName := fmt.Sprintf("%s.txt", identifier)
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Initialize a slice to store integers
	var integers []int

	// Use bufio.Scanner for line-by-line reading
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Trim spaces and split if numbers are space-separated
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(line) // Splits the line into parts by spaces

		// Convert each part to an integer
		for _, part := range parts {
			number, err := strconv.Atoi(part)
			if err != nil {
				fmt.Printf("Skipping invalid number: %s\n", part)
				continue
			}
			integers = append(integers, number)
		}
	}

	// Handle potential scanner errors
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	return integers, nil
}

func GetList(identifier string) ([]int, error) {
	switch identifier {
	case "left":
		contents, err := OpenFile("left")
		if err != nil {
			return nil, err
		}

		return contents, nil
	case "right":
		contents, err := OpenFile("right")
		if err != nil {
			return nil, err
		}

		return contents, nil
	default:
		return nil, fmt.Errorf("invalid identifier")
	}
}

func SortListsNumericallyAscending(slice *[]int) {
	for i := 0; i < len(*slice); i++ {
		for j := i + 1; j < len(*slice); j++ {
			if (*slice)[i] > (*slice)[j] {
				(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
			}
		}
	}
}

func GetTotalDistance(left, right *[]int) int {
	totalDistance := 0
	for i := 0; i < len(*left); i++ {

		distance := (*left)[i] - (*right)[i]
		absoluteDistance := math.Abs(float64(distance))
		totalDistance += int(absoluteDistance)
	}

	return totalDistance
}

func main() {
	left, err := GetList("left")
	if err != nil {
		fmt.Println(err)
		return
	}
	SortListsNumericallyAscending(&left)
	right, err := GetList("right")

	if err != nil {
		fmt.Println(err)
		return
	}
	SortListsNumericallyAscending(&right)

	totalDistance := GetTotalDistance(&left, &right)
	fmt.Println(totalDistance)
}
