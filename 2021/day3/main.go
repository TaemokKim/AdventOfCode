package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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
	inputValue := scanner.Text()
	length := len(inputValue)

	most := make([]int, length)

	addMost(most, inputValue)

	for scanner.Scan() {
		inputValue := scanner.Text()

		addMost(most, inputValue)
	}

	gammaRate := 0
	epsilonRate := 0
	for i, v := range most {
		realValue := int(math.Pow(2, float64(length-i-1)))
		if v > 0 {
			gammaRate += realValue
		} else {
			epsilonRate += realValue
		}
	}

	fmt.Println("part1: ", gammaRate*epsilonRate)
}

func addMost(most []int, inputValue string) {
	for pos, char := range inputValue {
		if char == '1' {
			most[pos] += 1
		} else {
			most[pos] -= 1
		}
	}
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	inputValue := scanner.Text()

	zeroList := []string{}
	oneList := []string{}
	most := 0
	index := 0

	most = addList(inputValue, index, most, &oneList, &zeroList)

	for scanner.Scan() {
		most = addList(scanner.Text(), index, most, &oneList, &zeroList)
	}

	pivotList := []string{}
	if most >= 0 {
		pivotList = make([]string, len(oneList))
		copy(pivotList, oneList)
	} else {
		pivotList = make([]string, len(zeroList))
		copy(pivotList, zeroList)
	}
	oxygenRating := oxygen(pivotList, index+1)

	if most > 0 {
		pivotList = make([]string, len(zeroList))
		copy(pivotList, zeroList)
	} else {
		pivotList = make([]string, len(oneList))
		copy(pivotList, oneList)
	}
	co2Rating := co2(pivotList, index+1)

	fmt.Println("part2: ", oxygenRating*co2Rating)
}

func oxygen(valueList []string, index int) int {
	if len(valueList) == 1 {
		oxygenRating := 0
		length := len(valueList[0])
		for i, v := range valueList[0] {
			realValue := int(math.Pow(2, float64(length-i-1)))
			if v == '1' {
				oxygenRating += realValue
			}
		}

		return oxygenRating
	}

	most := 0
	oneList := []string{}
	zeroList := []string{}

	for _, inputValue := range valueList {
		most = addList(inputValue, index, most, &oneList, &zeroList)
	}

	pivotList := []string{}
	if most >= 0 {
		pivotList = make([]string, len(oneList))
		copy(pivotList, oneList)
	} else {
		pivotList = make([]string, len(zeroList))
		copy(pivotList, zeroList)
	}

	return oxygen(pivotList, index+1)
}

func co2(valueList []string, index int) int {
	if len(valueList) == 1 {
		co2Rating := 0
		length := len(valueList[0])
		for i, v := range valueList[0] {
			realValue := int(math.Pow(2, float64(length-i-1)))
			if v == '1' {
				co2Rating += realValue
			}
		}

		return co2Rating
	}

	most := 0
	oneList := []string{}
	zeroList := []string{}

	for _, inputValue := range valueList {
		most = addList(inputValue, index, most, &oneList, &zeroList)
	}

	pivotList := []string{}
	if most >= 0 {
		pivotList = make([]string, len(zeroList))
		copy(pivotList, zeroList)
	} else {
		pivotList = make([]string, len(oneList))
		copy(pivotList, oneList)
	}

	return co2(pivotList, index+1)
}

func addList(inputValue string, index int, most int, oneList *[]string, zeroList *[]string) int {
	mostValue := 0
	if inputValue[index] == '1' {
		mostValue = most + 1
		*oneList = append(*oneList, inputValue)
	} else {
		mostValue = most - 1
		*zeroList = append(*zeroList, inputValue)
	}

	return mostValue
}
