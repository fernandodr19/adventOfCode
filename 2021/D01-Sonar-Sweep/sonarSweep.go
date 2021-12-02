package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read file to buffer
	b, err := ioutil.ReadFile("sonarSweep-data.txt")
	if err != nil {
		fmt.Print(err)
	}

	// convert content to a string
	str := string(b)

	// split into an slice
	depths, err := parseInputToInts(str)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, "Part I: %d\n", depthIncreaseCounter(depths))
	fmt.Fprintf(os.Stdout, "Part II: %d\n", depthWindowIncreaseCounter(depths))
}

// https://adventofcode.com/2021/day/1#part1
func depthIncreaseCounter(arr []int) int {
	count := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			count++
		}
	}

	return count
}

// https://adventofcode.com/2021/day/1#part2
func depthWindowIncreaseCounter(arr []int) int {
	count := 0
	for i := 2; i < len(arr)-1; i++ {
		a := sliceSum(arr[i-2 : i+1])
		b := sliceSum(arr[i-1 : i+2])
		if b > a {
			count++
		}
	}

	return count
}

// sliceSum returns the sum of the inputed slice
func sliceSum(arr []int) int {
	sum := 0
	for _, e := range arr {
		sum += e
	}
	return sum
}

// parseInputToInts returns a slice of numbers for a given input string
func parseInputToInts(input string) ([]int, error) {
	values := strings.Split(input, "\n")
	arr := make([]int, 0, len(values))

	for _, l := range values {
		if len(l) == 0 {
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		arr = append(arr, n)
	}
	return arr, nil
}
