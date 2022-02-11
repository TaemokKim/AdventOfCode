package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	part1(100)
	part2()
}

func part1(steps int) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	energyMap := [][]int{}

	for scanner.Scan() {
		energyMap = append(energyMap, string2Int(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	paddingLine := make([]int, len(energyMap[0]))
	tempMap := [][]int{}
	tempMap = append(tempMap, paddingLine)
	tempMap = append(tempMap, energyMap...)
	tempMap = append(tempMap, paddingLine)
	energyMap = append(tempMap, paddingLine)

	flashCount := 0

	for i := 0; i < steps; i++ {
		flashCount += step(energyMap)
	}

	fmt.Println("part1: ", flashCount)
}

func string2Int(stringLine string) []int {
	intValue := make([]int, len(stringLine)+3)

	for i, v := range stringLine {
		intValue[i+1], _ = strconv.Atoi(string(v))
	}

	return intValue
}

func step(energyMap [][]int) int {
	flashCount := 0

	rows := len(energyMap) - 3
	cols := len(energyMap[0]) - 3

	for row := 1; row < rows+1; row++ {
		for col := 1; col < cols+1; col++ {
			energyMap[row][col] += 1
		}
	}

	for {
		flashes := false
		for row := 1; row < rows+1; row++ {
			for col := 1; col < cols+1; col++ {
				if energyMap[row][col] == 10 {
					flash(energyMap, row, col)
					energyMap[row][col] = 11
					flashes = true
				}
			}
		}

		if !flashes {
			break
		}
	}

	for row := 1; row < rows+1; row++ {
		for col := 1; col < cols+1; col++ {
			if energyMap[row][col] > 9 {
				energyMap[row][col] = 0
				flashCount += 1
			}
		}
	}

	return flashCount
}

func flash(energyMap [][]int, row, col int) {
	flashMap := energyMap[row-1 : row+2]

	for r := 0; r < 3; r++ {
		scanLine := flashMap[r][col-1 : col+2]

		for c := 0; c < 3; c++ {
			if scanLine[c] < 10 {
				scanLine[c] += 1
			}
		}
	}
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	energyMap := [][]int{}

	for scanner.Scan() {
		energyMap = append(energyMap, string2Int(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	paddingLine := make([]int, len(energyMap[0]))
	tempMap := [][]int{}
	tempMap = append(tempMap, paddingLine)
	tempMap = append(tempMap, energyMap...)
	tempMap = append(tempMap, paddingLine)
	energyMap = append(tempMap, paddingLine)

	stepCount := 0
	for {
		flashCount := step(energyMap)
		stepCount += 1

		if flashCount == (len(energyMap)-3)*(len(energyMap[0])-3) {
			break
		}
	}

	fmt.Println("part2: ", stepCount)
}
