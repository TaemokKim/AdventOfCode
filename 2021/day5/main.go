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

type line struct {
	p1 point
	p2 point
}

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

	lineList := []line{}

	maxX := 0
	maxY := 0

	for scanner.Scan() {
		textLine := scanner.Text()
		pointList := strings.Split(textLine, "->")
		point1 := strings.Split(pointList[0], ",")
		point2 := strings.Split(pointList[1], ",")

		p1 := point{}
		p1.x, _ = strconv.Atoi(strings.Trim(point1[0], " "))
		p1.y, _ = strconv.Atoi(strings.Trim(point1[1], " "))

		p2 := point{}
		p2.x, _ = strconv.Atoi(strings.Trim(point2[0], " "))
		p2.y, _ = strconv.Atoi(strings.Trim(point2[1], " "))

		if p1.x == p2.x || p1.y == p2.y {
			l := line{}
			if p1.x > p2.x || p1.y > p2.y {
				l.p1 = p2
				l.p2 = p1
			} else {
				l.p1 = p1
				l.p2 = p2
			}

			if maxX < l.p2.x {
				maxX = l.p2.x
			}
			if maxY < l.p2.y {
				maxY = l.p2.y
			}

			lineList = append(lineList, l)
		}
	}

	pointMap := make([][]int, maxY+1)
	for i := range pointMap {
		pointMap[i] = make([]int, maxX+1)

		for j := range pointMap[i] {
			pointMap[i][j] = 0
		}
	}

	for i := 0; i < len(lineList); i += 1 {
		line := lineList[i]

		for y := line.p1.y; y <= line.p2.y; y += 1 {
			for x := line.p1.x; x <= line.p2.x; x += 1 {
				pointMap[y][x] += 1
			}
		}
	}

	count := 0
	for y := 0; y < maxY; y += 1 {
		for x := 0; x < maxX; x += 1 {
			if pointMap[y][x] > 1 {
				count += 1
			}
		}
	}

	fmt.Println("part1", count)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	lineList := []line{}

	maxX := 0
	maxY := 0

	for scanner.Scan() {
		textLine := scanner.Text()
		pointList := strings.Split(textLine, "->")
		point1 := strings.Split(pointList[0], ",")
		point2 := strings.Split(pointList[1], ",")

		p1 := point{}
		p1.x, _ = strconv.Atoi(strings.Trim(point1[0], " "))
		p1.y, _ = strconv.Atoi(strings.Trim(point1[1], " "))

		p2 := point{}
		p2.x, _ = strconv.Atoi(strings.Trim(point2[0], " "))
		p2.y, _ = strconv.Atoi(strings.Trim(point2[1], " "))

		l := line{}
		l.p1 = p1
		l.p2 = p2

		if maxX < l.p1.x {
			maxX = l.p1.x
		}
		if maxY < l.p1.y {
			maxY = l.p1.y
		}
		if maxX < l.p2.x {
			maxX = l.p2.x
		}
		if maxY < l.p2.y {
			maxY = l.p2.y
		}

		lineList = append(lineList, l)
	}

	pointMap := make([][]int, maxY+1)
	for i := range pointMap {
		pointMap[i] = make([]int, maxX+1)

		for j := range pointMap[i] {
			pointMap[i][j] = 0
		}
	}

	for _, line := range lineList {
		if line.p1.x == line.p2.x {
			startY := 0
			endY := 0
			if line.p1.y < line.p2.y {
				startY = line.p1.y
				endY = line.p2.y
			} else {
				startY = line.p2.y
				endY = line.p1.y
			}

			for j := startY; j < endY+1; j += 1 {
				pointMap[j][line.p1.x] += 1
			}
		} else if line.p1.y == line.p2.y {
			startX := 0
			endX := 0
			if line.p1.x < line.p2.x {
				startX = line.p1.x
				endX = line.p2.x
			} else {
				startX = line.p2.x
				endX = line.p1.x
			}

			for j := startX; j < endX+1; j += 1 {
				pointMap[line.p1.y][j] += 1
			}
		} else {
			lineLength := 0
			incValueX := 0
			if line.p1.x < line.p2.x {
				incValueX = 1
				lineLength = line.p2.x - line.p1.x + 1
			} else {
				incValueX = -1
				lineLength = line.p1.x - line.p2.x + 1
			}

			incValueY := 0
			if line.p1.y < line.p2.y {
				incValueY = 1
			} else {
				incValueY = -1
			}

			indexX := line.p1.x
			indexY := line.p1.y
			for j := 0; j < lineLength; j += 1 {
				pointMap[indexY][indexX] += 1
				indexY += incValueY
				indexX += incValueX
			}
		}
	}

	count := 0
	for y := 0; y < maxY; y += 1 {
		for x := 0; x < maxX; x += 1 {
			if pointMap[y][x] > 1 {
				count += 1
			}
		}
	}

	fmt.Println("part2", count)
}
