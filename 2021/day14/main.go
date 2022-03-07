package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

const interval int = 10

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

	charMap := make(map[string]uint64)

	divide(template, insertRule, charMap, stepCount)

	countChar(string(template[len(template)-1]), charMap)

	max := uint64(0)
	min := uint64(math.MaxUint64)
	for _, v := range charMap {
		if v < min {
			min = uint64(v)
		}

		if v > max {
			max = uint64(v)
		}
	}

	fmt.Println("Step", stepCount, ": ", max, min, max-min)
}

func divide(template string, insertRule map[string]string, charMap map[string]uint64, stepCount int) {
	result := step(template, insertRule)

	if stepCount == 1 {
		result = result[0 : len(result)-1]
		countChar(result, charMap)
	} else {
		rest := len(result) % interval

		index := 0
		for ; index < len(result)-interval-rest+1; index += (interval - 1) {
			divide(result[index:index+interval], insertRule, charMap, stepCount-1)
		}

		if rest > 0 {
			divide(result[index:len(result)], insertRule, charMap, stepCount-1)
		}
	}
}

func step(template string, insertRule map[string]string) string {
	result := ""
	for i := 0; i < len(template)-1; i++ {
		result = result + string(template[i]) + insertRule[template[i:i+2]]
	}

	result += string(template[len(template)-1])

	return result
}

func countChar(template string, charMap map[string]uint64) {
	for _, v := range template {
		charMap[string(v)]++
	}
}
