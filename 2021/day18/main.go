package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	//f, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	part1(f)

	// call the Seek method first
	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}

	part2(f)
}

func part1(f *os.File) {
	scanner := bufio.NewScanner(f)
	scanner.Scan()

	numbers := strings.Split(scanner.Text(), "")

	for scanner.Scan() {
		second := strings.Split(scanner.Text(), "")
		numbers = append([]string{"["}, numbers...)
		numbers = append(numbers, ",")
		numbers = append(numbers, second...)
		numbers = append(numbers, "]")

		numbers = reduce(numbers)
	}

	sum := calcMagnitude(numbers)

	fmt.Println(sum)
}

func part2(f *os.File) {
	scanner := bufio.NewScanner(f)
	scanner.Scan()

	numberList := [][]string{}

	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), "")
		numberList = append(numberList, numbers)
	}

	maxValue := 0

	for i := 0; i < len(numberList); i++ {
		for j := 0; j < len(numberList); j++ {
			if i == j {
				continue
			}

			numbers := append([]string{"["}, numberList[i]...)
			numbers = append(numbers, ",")
			numbers = append(numbers, numberList[j]...)
			numbers = append(numbers, "]")

			numbers = reduce(numbers)
			sum := calcMagnitude(numbers)
			if sum > maxValue {
				maxValue = sum
			}
		}
	}

	fmt.Println(maxValue)

}

func reduce(numbers []string) []string {
	isExploded := true
	isSplitted := true
	clean := true

	for isExploded || isSplitted {
		numbers, isExploded = reduceExplode(numbers)

		for clean {
			numbers, clean = cleanList(numbers)
		}
		clean = true

		numbers, isSplitted = reduceSplit(numbers)

		if isSplitted {
			for clean {
				numbers, clean = cleanList(numbers)
			}
			clean = true
		}
	}

	return numbers
}

func reduceExplode(numbers []string) ([]string, bool) {
	openCount := 0
	explodRight := 0
	resultNumbers := []string{}
	leftNumberIndex := 0
	isChanged := false

	for i := 0; i < len(numbers); i++ {
		v := numbers[i]

		if v == "[" {
			openCount++
		} else if v == "]" {
			openCount--
		} else if v != "," {
			if explodRight > 0 {
				rightNumber, _ := strconv.Atoi(v)
				rightNumber += explodRight
				v = strconv.Itoa(rightNumber)

				explodRight = 0
			}

			if openCount >= 5 && numbers[i+1] != "[" {
				isChanged = true

				resultNumbers = resultNumbers[0 : len(resultNumbers)-1]
				openCount--

				left, _ := strconv.Atoi(v)

				if leftNumberIndex > 0 {
					leftNumber, _ := strconv.Atoi(resultNumbers[leftNumberIndex])
					left += leftNumber
					resultNumbers[leftNumberIndex] = strconv.Itoa(left)
				}

				i++
				v = numbers[i]
				if v == "," {
					i++
					v = numbers[i]
				}

				explodRight, _ = strconv.Atoi(v)

				resultNumbers = append(resultNumbers, "0")
				leftNumberIndex = len(resultNumbers) - 1

				i++

				continue
			} else {
				leftNumberIndex = len(resultNumbers)
			}
		}

		resultNumbers = append(resultNumbers, v)
	}

	return resultNumbers, isChanged
}

func reduceSplit(numbers []string) ([]string, bool) {
	resultNumbers := []string{}
	isChanged := false

	for i := 0; i < len(numbers); i++ {
		v := numbers[i]
		if v != "[" && v != "]" && v != "," {
			currentNumber, _ := strconv.Atoi(v)

			if currentNumber >= 10 && isChanged == false {
				isChanged = true

				left := int(math.Floor(float64(currentNumber) / 2))
				right := int(math.Ceil(float64(currentNumber) / 2))

				leftNumber := strconv.Itoa(left)
				rightNumber := strconv.Itoa(right)

				resultNumbers = append(resultNumbers, "[")
				resultNumbers = append(resultNumbers, leftNumber)
				resultNumbers = append(resultNumbers, rightNumber)
				resultNumbers = append(resultNumbers, "]")

				continue
			}
		}

		resultNumbers = append(resultNumbers, v)
	}

	return resultNumbers, isChanged
}

func cleanList(numbers []string) ([]string, bool) {
	resultNumbers := []string{}

	numberCount := 0
	isChanged := false

	for i := 0; i < len(numbers); i++ {
		v := numbers[i]

		if v == "[" {
			numberCount = 0
		} else if v == "]" {
			if resultNumbers[len(resultNumbers)-1] == "[" {
				resultNumbers = resultNumbers[0 : len(resultNumbers)-1]
				isChanged = true

				continue
			} else if numberCount == 1 && resultNumbers[len(resultNumbers)-2] == "[" {
				temp := resultNumbers[len(resultNumbers)-1]
				if len(resultNumbers) > 2 {
					resultNumbers = resultNumbers[0 : len(resultNumbers)-2]
				} else {
					resultNumbers = []string{}
				}

				resultNumbers = append(resultNumbers, temp)

				numberCount = 0
				isChanged = true

				continue
			}

			numberCount = 0
		} else if v == "," {
			continue
		} else {
			numberCount++
		}

		resultNumbers = append(resultNumbers, v)
	}

	return resultNumbers, isChanged

}

func calcMagnitude(numbers []string) int {
	resultNumbers := []int{}

	for i := 0; i < len(numbers); i++ {
		v := numbers[i]

		if v == "]" && len(resultNumbers) > 1 {
			lastIndex := len(resultNumbers) - 1
			result := 3*resultNumbers[lastIndex-1] + 2*resultNumbers[lastIndex]
			resultNumbers = resultNumbers[0:lastIndex]
			resultNumbers[lastIndex-1] = result
		} else if v != "[" && v != "," {
			currentNumber, _ := strconv.Atoi(v)
			resultNumbers = append(resultNumbers, currentNumber)
		}
	}

	return resultNumbers[0]
}
