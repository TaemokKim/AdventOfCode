package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	part1(10)
	part1(40)
}

func part1(stepCount int) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	template := scanner.Text()

	scanner.Scan()

	insertRule := make(map[string]string)
	for scanner.Scan() {
		inputLine := scanner.Text()

		rule := strings.Split(inputLine, " -> ")

		insertRule[rule[0]] = rule[1]
	}

	result := template

	for i := 0; i < stepCount; i++ {
		result = step(result, insertRule)
	}

	charMap := countChar(result)

	max := 0
	min := int(^uint(0) >> 1)
	for _, v := range charMap {
		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	fmt.Println("part1: ", max-min)
}

func step(template string, insertRule map[string]string) string {
	result := ""
	for i := 0; i < len(template)-1; i++ {
		temp := string(template[i]) + string(template[i+1])
		v, exists := insertRule[temp]
		result += string(template[i])
		if exists {
			result += v
		}
	}

	result += string(template[len(template)-1])

	return result
}

func countChar(template string) map[string]int {
	charMap := make(map[string]int)
	for _, v := range template {
		key := string(v)
		_, exists := charMap[key]
		if !exists {
			charMap[key] = 1
		} else {
			charMap[key] += 1
		}
	}

	return charMap
}
