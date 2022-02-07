package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	stack := []rune{}
	closeMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	pointMap := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		for _, v := range line {
			if v == '(' || v == '[' || v == '{' || v == '<' {
				stack = append(stack, v)
			} else {
				if closeMap[stack[len(stack)-1]] == v {
					stack = stack[:len(stack)-1]
				} else {
					sum += pointMap[v]
					break
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	fmt.Println("part1: ", sum)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	closeMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	pointMap := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	scoreList := []int{}
	for scanner.Scan() {
		stack := []rune{}
		line := scanner.Text()

		for _, v := range line {
			if v == '(' || v == '[' || v == '{' || v == '<' {
				stack = append(stack, v)
			} else {
				if closeMap[stack[len(stack)-1]] == v {
					stack = stack[:len(stack)-1]
				} else {
					stack = nil
					break
				}
			}
		}

		score := 0
		for i := len(stack) - 1; i >= 0; i-- {
			score = score*5 + pointMap[closeMap[stack[i]]]
		}

		if score > 0 {
			scoreList = append(scoreList, score)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	sort.Slice(scoreList, func(i, j int) bool {
		return scoreList[i] < scoreList[j]
	})

	middleScore := scoreList[len(scoreList)/2]

	fmt.Println("part2: ", middleScore)
}
