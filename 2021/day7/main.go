package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
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

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	hPositionString := strings.Split(scanner.Text(), ",")

	hPosition := []int{}
	maxPosition := 0

	for _, v := range hPositionString {
		position, _ := strconv.Atoi(v)

		hPosition = append(hPosition, position)

		if maxPosition < position {
			maxPosition = position
		}
	}

	sort.Slice(hPosition, func(i, j int) bool {
		return hPosition[i] < hPosition[j]
	})

	minValue := math.MaxInt
	position := math.MinInt
	for i := 0; i < maxPosition+1; i++ {
		if position >= i {
			continue
		}

		position = i
		sum := 0
		for _, v := range hPosition {
			sum += int(math.Abs(float64(v - position)))
		}

		if minValue > sum {
			minValue = sum
		}
	}

	fmt.Println("part1: ", minValue)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	hPositionString := strings.Split(scanner.Text(), ",")

	hPosition := []int{}
	maxPosition := 0
	for _, v := range hPositionString {
		position, _ := strconv.Atoi(v)

		hPosition = append(hPosition, position)

		if maxPosition < position {
			maxPosition = position
		}
	}

	sort.Slice(hPosition, func(i, j int) bool {
		return hPosition[i] < hPosition[j]
	})

	minValue := math.MaxInt
	position := math.MinInt
	for i := 0; i < maxPosition+1; i++ {
		if position >= i {
			continue
		}

		position = i
		sum := 0
		for _, v := range hPosition {
			n := int(math.Abs(float64(v - position)))
			sum += n * (n + 1) / 2
		}

		if minValue > sum {
			minValue = sum
		}
	}

	fmt.Println("part2: ", minValue)
}
