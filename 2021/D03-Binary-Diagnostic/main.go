package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	// read file to buffer
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}

	// convert content to a string
	str := string(b)

	lines := strings.Split(str, "\n")

	gammaRate, epsilonRate, err := findGammaEpsilon(lines)
	if err != nil {
		panic(err)
	}

	fmt.Println(gammaRate * epsilonRate) // 3687446

	o2Rate, err := findGasRate(0, lines, true)
	if err != nil {
		panic(err)
	}

	co2Rate, err := findGasRate(0, lines, false)
	if err != nil {
		panic(err)
	}

	fmt.Println(o2Rate * co2Rate) // 4406844
}

func findGammaEpsilon(lines []string) (int64, int64, error) {
	// intialize maps
	bitMaps := make([]map[string]int, 12)
	for i := range bitMaps {
		bitMaps[i] = make(map[string]int)
	}

	for _, line := range lines {
		// break line into bits
		bits := strings.Split(line, "")

		for i, bit := range bits {
			fmt.Println(i)
			fmt.Println(len(bitMaps))
			bitMaps[i][bit]++
		}
	}

	binaryGammaRate, binaryEpsilonRate := "", ""
	for _, bitMap := range bitMaps {
		if bitMap["0"] > bitMap["1"] {
			binaryGammaRate += "0"
			binaryEpsilonRate += "1"
		} else if bitMap["0"] < bitMap["1"] {
			binaryGammaRate += "1"
			binaryEpsilonRate += "0"
		} else {
			return 0, 0, errors.New("0 == 1")
		}
	}

	fmt.Println(fmt.Sprintf("Binary gamma rate: %s", binaryGammaRate))
	fmt.Println(fmt.Sprintf("Binary epsilon rate: %s", binaryEpsilonRate))

	gammaRate, err := strconv.ParseInt(binaryGammaRate, 2, 64)
	if err != nil {
		return 0, 0, err
	}

	epsilonRate, err := strconv.ParseInt(binaryEpsilonRate, 2, 64)
	if err != nil {
		return 0, 0, err
	}

	fmt.Println(fmt.Sprintf("Gamma rate: %d", gammaRate))
	fmt.Println(fmt.Sprintf("Epsilon rate: %d", epsilonRate))
	return gammaRate, epsilonRate, nil
}

func findGasRate(index int, lines []string, o2 bool) (int64, error) {
	// intialize maps
	bitMap := make(map[string]int)

	for _, line := range lines {
		// break line into bits
		bits := strings.Split(line, "")

		bitMap[bits[index]]++
	}

	if bitMap["0"] > bitMap["1"] {
		if o2 {
			lines = removeLines(index, "1", lines)
		} else {
			lines = removeLines(index, "0", lines)
		}
	} else {
		if o2 {
			lines = removeLines(index, "0", lines)
		} else {
			lines = removeLines(index, "1", lines)
		}
	}

	if len(lines) != 1 {
		return findGasRate(index+1, lines, o2)
	}

	// only one line left

	fmt.Println(fmt.Sprintf("Binary gas rate: %s", lines[0]))

	oxygenRate, err := strconv.ParseInt(lines[0], 2, 64)
	if err != nil {
		return 0, err
	}

	fmt.Println(fmt.Sprintf("Gas rate: %d", oxygenRate))
	return oxygenRate, nil
}

// removeLines remove all lines where line[index] == bit
func removeLines(index int, bit string, lines []string) []string {
	remainingLines := make([]string, 0)
	for _, line := range lines {
		// break line into bits
		bits := strings.Split(line, "")
		if bits[index] != bit {
			remainingLines = append(remainingLines, line)
		}
	}

	return remainingLines
}
