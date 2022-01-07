package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	result := 0

	scanner.Scan()
	prev, _ := strconv.Atoi(scanner.Text())

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		if prev < num {
			result += 1
		}
		prev = num
	}

	fmt.Println("part1: ", result)

}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	result := 0

	window := make([]int, 0, 3)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		window = append(window, num)

		if len(window) == 3 {
			break
		}
	}

	if len(window) != 3 {
		fmt.Println("need more than 2 datas")
	}

	index := 0

	sum := func(a []int) int {
		return a[0] + a[1] + a[2]
	}

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		prev := sum(window)

		window[index] = num
		cur := sum(window)

		if prev < cur {
			result += 1
		}
		prev = cur
		index = (index + 1) % 3
	}

	fmt.Println("part2: ", result)

}
