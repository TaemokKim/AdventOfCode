package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	part1()
	part2()
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	caveMap := make(map[string][]string)
	caveCountMap := make(map[string]int)

	for scanner.Scan() {
		slice := strings.Split(scanner.Text(), "-")
		updateCaveMap(caveMap, slice[0], slice[1])
		updateCaveMap(caveMap, slice[1], slice[0])

		updateCaveCountMap(caveCountMap, slice[0], 1)
		updateCaveCountMap(caveCountMap, slice[1], 1)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	pathCount := findPath("start", caveMap, caveCountMap)

	fmt.Println("part1: ", pathCount)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	caveMap := make(map[string][]string)
	caveCountMap := make(map[string]int)

	for scanner.Scan() {
		slice := strings.Split(scanner.Text(), "-")
		updateCaveMap(caveMap, slice[0], slice[1])
		updateCaveMap(caveMap, slice[1], slice[0])

		updateCaveCountMap(caveCountMap, slice[0], 2)
		updateCaveCountMap(caveCountMap, slice[1], 2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	pathCount := findPath("start", caveMap, caveCountMap)

	fmt.Println("part2: ", pathCount)
}

func updateCaveMap(caveMap map[string][]string, value1, value2 string) {
	if value1 != "end" && value2 != "start" {
		if list, ok := caveMap[value1]; ok {
			list = append(list, value2)
			caveMap[value1] = list
		} else {
			caveMap[value1] = []string{value2}
		}
	}

}

func updateCaveCountMap(caveCountMap map[string]int, value string, pivot int) {
	if isLower(value) {
		caveCountMap[value] = pivot
	} else {
		caveCountMap[value] = -1
	}

}

// https://stackoverflow.com/questions/59293525/how-to-check-if-a-string-is-all-upper-or-lower-case-in-go
func isLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func findPath(s string, caveMap map[string][]string, caveCountMap map[string]int) int {
	if caveCountMap[s] == 0 {
		return 0
	} else if caveCountMap[s] > 0 {
		caveCountMap[s] -= 1
	}

	if s == "end" {
		caveCountMap[s] += 1
		return 1
	}

	newCountMap := map[string]int{}
	if caveCountMap[s] == 0 && caveCountMap["end"] == 2 {
		newCountMap = make(map[string]int)
		for k, v := range caveCountMap {
			if v > 0 {
				newCountMap[k] = v - 1
			} else {
				newCountMap[k] = v
			}
		}
	} else {
		newCountMap = caveCountMap
	}

	pathCount := 0

	for _, v := range caveMap[s] {
		pathCount += findPath(v, caveMap, newCountMap)
	}

	if caveCountMap[s] != -1 {
		caveCountMap[s] += 1
	}

	return pathCount
}
