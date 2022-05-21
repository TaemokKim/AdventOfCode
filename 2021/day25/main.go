package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	//f, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	input := [][]rune{}

	for s.Scan() {
		state := []rune{}

		for _, c := range s.Text() {
			state = append(state, c)
		}

		input = append(input, state)
	}

	index := 1

	for step(input) {
		index++
	}

	//printInput(input)

	fmt.Println("stop: ", index)
}

func printInput(input [][]rune) {
	for _, line := range input {
		fmt.Println(string(line))
	}
}

func step(input [][]rune) bool {
	ret1 := checkEast(input)
	ret2 := checkSouth(input)

	return ret1 || ret2
}

func checkEast(input [][]rune) bool {
	width := len(input[0])
	
	move := false

	for y:=0; y<len(input); y++ {
		endX := 0
		if input[y][0] == '.' {
			endX = 1
		}

		for x:=width-1; x>=endX; x-- {
			if input[y][x] == '>' {
				if input[y][(x+1)%width] == '.' {
					input[y][x] = '.'
					input[y][(x+1)%width] = '>'
					x--
					move = true
				}
			}
		}
	}

	return move
}

func checkSouth(input [][]rune) bool {
	width := len(input[0])
	height := len(input)
	
	move := false

	for x:=0; x<width; x++ {	
		endY := 0
		if input[0][x] == '.' {
			endY = 1
		}

		for y:=height-1; y>=endY; y-- {
			if input[y][x] == 'v' {
				if input[(y+1)%height][x] == '.' {
					input[y][x] = '.'
					input[(y+1)%height][x] = 'v'
					y--
					move = true
				}
			}
		}
	}

	return move
}
