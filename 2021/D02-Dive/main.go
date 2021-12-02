package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	forward string = "forward"
	up             = "up"
	down           = "down"
)

type Command struct {
	Direction string
	Distance  int
}

func main() {
	// read file to buffer
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}

	// convert content to a string
	str := string(b)

	// split into a slice of commands
	commands, err := parseInputToCommands(str)
	if err != nil {
		panic(err)
	}

	simpleDistance, err := getSimpleDistance(commands)
	if err != nil {
		panic(err)
	}

	complexDistance, err := getComplexDistance(commands)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("Part I: %d", simpleDistance))
	fmt.Println(fmt.Sprintf("Part II: %d", complexDistance))

}

// https://adventofcode.com/2021/day/2#part1
func getSimpleDistance(commands []Command) (int, error) {
	horizontal, depth := 0, 0
	for _, c := range commands {
		switch c.Direction {
		case forward:
			horizontal += c.Distance
		case up:
			depth -= c.Distance
		case down:
			depth += c.Distance
		default:
			return 0, errors.New(fmt.Sprintf("Invalid direction: %s", c.Direction))
		}
	}
	return horizontal * depth, nil
}

// https://adventofcode.com/2021/day/2#part2
func getComplexDistance(commands []Command) (int, error) {
	horizontal, depth, aim := 0, 0, 0
	for _, c := range commands {
		switch c.Direction {
		case forward:
			horizontal += c.Distance
			depth += c.Distance * aim
		case up:
			aim -= c.Distance
		case down:
			aim += c.Distance
		default:
			return 0, errors.New(fmt.Sprintf("Invalid direction: %s", c.Direction))
		}
	}
	return horizontal * depth, nil
}

// parseInputToCommands parses each input line into a Command
func parseInputToCommands(input string) ([]Command, error) {
	lines := strings.Split(input, "\n")
	arr := make([]Command, 0, len(lines))

	for _, line := range lines {
		values := strings.Split(line, " ")

		if len(values) != 2 {
			return nil, errors.New(fmt.Sprintf("Could not parse line: %s", line))
		}

		dir := values[0]
		switch dir {
		case forward, up, down:
			// does nothing, valid direction!
		default:
			return nil, errors.New(fmt.Sprintf("Could not parse direction: %s", dir))
		}

		distance, err := strconv.Atoi(values[1])
		if err != nil {
			return nil, err
		}

		arr = append(arr, Command{
			Direction: dir,
			Distance:  distance,
		})
	}
	return arr, nil
}
