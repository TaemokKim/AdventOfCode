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
	bingoNumbers := scanner.Text()

	bingoNumberStringList := strings.Split(bingoNumbers, ",")
	bingoNumberList := make([]int, len(bingoNumberStringList))
	for i, v := range bingoNumberStringList {
		bingoNumberList[i], _ = strconv.Atoi(v)
	}

	minCount := len(bingoNumberList)
	lastNumber := -1
	sum := -1

	for {
		userValue := getUserValue(scanner)
		if userValue == nil {
			break
		}

		count, last, s := matchValue(bingoNumberList, userValue)
		if count >= 0 && count < minCount {
			minCount = count
			lastNumber = last
			sum = s
		}
	}

	fmt.Println("part1: ", lastNumber, sum, lastNumber*sum)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	bingoNumbers := scanner.Text()

	bingoNumberStringList := strings.Split(bingoNumbers, ",")
	bingoNumberList := make([]int, len(bingoNumberStringList))
	for i, v := range bingoNumberStringList {
		bingoNumberList[i], _ = strconv.Atoi(v)
	}

	maxCount := 0
	lastNumber := -1
	sum := -1

	for {
		userValue := getUserValue(scanner)
		if userValue == nil {
			break
		}

		count, last, s := matchValue(bingoNumberList, userValue)
		if count >= 0 && count > maxCount {
			maxCount = count
			lastNumber = last
			sum = s
		}
	}

	fmt.Println("part2: ", lastNumber, sum, lastNumber*sum)
}

func getUserValue(scanner *bufio.Scanner) [][]int {
	row := 0

	var userValue [][]int

	for scanner.Scan() {
		if userValue == nil {
			userValue = make([][]int, 5)
		}

		scanLine := scanner.Text()
		if len(scanLine) == 0 {
			break
		}

		userValue[row] = make([]int, 5)

		line := strings.Split(scanLine, " ")

		col := 0
		for _, v := range line {
			if v != "" {
				userValue[row][col], _ = strconv.Atoi(v)
				col += 1
			}
		}

		row += 1
	}

	return userValue
}

func matchValue(bingoNumbers []int, userValue [][]int) (int, int, int) {
	for count, number := range bingoNumbers {
		for row := range userValue {
			for col, v := range userValue[row] {
				if v == number {
					userValue[row][col] = -1
					if isBingo(userValue, row, col) {
						sum := getSum(userValue)

						return count, number, sum
					}
				}
			}
		}
	}

	return -1, -1, -1
}

func isBingo(userValue [][]int, row int, col int) bool {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += userValue[row][i]
	}

	if sum == -5 {
		return true
	}

	sum = 0
	for i := 0; i < 5; i++ {
		sum += userValue[i][col]
	}

	if sum == -5 {
		return true
	}

	return false
}

func getSum(userValue [][]int) int {
	sum := 0

	for row := range userValue {
		for _, colV := range userValue[row] {
			if colV > 0 {
				sum += colV
			}
		}
	}

	return sum
}
