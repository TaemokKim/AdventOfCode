package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1(80)
	part2(256)
}

func part1(day int) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	init := strings.Split(scanner.Text(), ",")
	fishList := []int{}

	for _, v := range init {
		state, _ := strconv.Atoi(v)

		leftDay := day - state

		for leftDay > 0 {
			fishList = append(fishList, leftDay)
			leftDay -= 7
		}
	}

	index := 0

	if len(fishList) > 0 {
		for {
			leftDay := fishList[index]

			leftDay -= 9

			for leftDay > 0 {
				fishList = append(fishList, leftDay)
				leftDay -= 7
			}

			if index >= (len(fishList) - 1) {
				break
			}

			index += 1
		}
	}

	fmt.Println("part1: ", len(init)+len(fishList))
}

func part2(day int) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	init := strings.Split(scanner.Text(), ",")

	count := uint64(len(init))

	valueMap := make(map[int]uint64)
	for _, v := range init {
		state, _ := strconv.Atoi(v)

		if valueMap[state] != 0 {
			count += valueMap[state]
		} else {

			leftDay := day - state

			currentCount := uint64(0)
			for leftDay > 0 {
				currentCount = currentCount + newFishCount(leftDay) + 1
				leftDay -= 7
			}

			valueMap[state] = currentCount
			count += currentCount
		}
	}

	fmt.Println("part2: ", count)
}

func newFishCount(day int) uint64 {
	leftDay := day - 9
	count := uint64(0)

	for leftDay > 0 {
		count = count + newFishCount(leftDay) + 1
		leftDay -= 7
	}

	return count
}
