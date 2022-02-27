package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func main() {
	part1()
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	coordList := []point{}

	for scanner.Scan() {
		inputLine := scanner.Text()
		if len(inputLine) == 0 {
			break
		}

		coord := strings.Split(inputLine, ",")
		p := point{}

		p.x, err = strconv.Atoi(coord[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		p.y, err = strconv.Atoi(coord[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		coordList = append(coordList, p)
	}

	for scanner.Scan() {
		instLine := strings.Split(scanner.Text(), " ")
		inst := strings.Split(instLine[2], "=")

		foldValue, err := strconv.Atoi(inst[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		if inst[0] == "x" {
			coordList = foldLeft(coordList, foldValue)
			fmt.Println("foldLeft: ", len(coordList))
		} else {
			coordList = foldUp(coordList, foldValue)
			fmt.Println("foldUp: ", len(coordList))
		}
	}

	printCoordList(coordList)
}

func foldLeft(coordList []point, left int) []point {
	visibleDots := []point{}
	leftDots := []point{}

	for _, v := range coordList {
		if v.x < left {
			leftDots = append(leftDots, v)
		}
	}

	for _, v := range coordList {
		if v.x > left {
			p := point{}
			p.x = left - (v.x - left)
			p.y = v.y

			if !isDuplicated(leftDots, p) {
				visibleDots = append(visibleDots, p)
			}
		}
	}

	visibleDots = append(visibleDots, leftDots...)

	return visibleDots
}

func foldUp(coordList []point, top int) []point {
	visibleDots := []point{}
	upDots := []point{}

	for _, v := range coordList {
		if v.y < top {
			upDots = append(upDots, v)
		}
	}

	for _, v := range coordList {
		if v.y > top {
			p := point{}
			p.x = v.x
			p.y = top - (v.y - top)

			if !isDuplicated(upDots, p) {
				visibleDots = append(visibleDots, p)
			}
		}
	}

	visibleDots = append(visibleDots, upDots...)

	return visibleDots
}

func isDuplicated(visibleDots []point, dot point) bool {
	for _, v := range visibleDots {
		if dot.x == v.x && dot.y == v.y {
			return true
		}
	}

	return false
}

func printCoordList(coordList []point) {
	width := 0
	height := 0

	for _, v := range coordList {
		if width < v.x {
			width = v.x
		}

		if height < v.y {
			height = v.y
		}
	}

	width += 1
	height += 1

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if (isDuplicated(coordList, point{x, y})) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}
