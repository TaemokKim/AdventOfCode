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

	horizontal := 0
	depth := 0

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		value, _ := strconv.Atoi(words[1])
		if words[0] == "forward" {
			horizontal += value
		} else if words[0] == "down" {
			depth += value
		} else if words[0] == "up" {
			depth -= value
		}
	}

	fmt.Println("part1: ", horizontal*depth)

}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	horizontal := 0
	depth := 0
	aim := 0

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		value, _ := strconv.Atoi(words[1])
		if words[0] == "forward" {
			horizontal += value
			depth += value * aim
		} else if words[0] == "down" {
			aim += value
		} else if words[0] == "up" {
			aim -= value
		}
	}

	fmt.Println("part2: ", horizontal*depth)

}
