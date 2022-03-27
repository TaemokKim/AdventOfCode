package main

import (
	"bufio"
	"fmt"
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

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	split := strings.Split(scanner.Text(), ":")
	split = strings.Split(split[1], ",")
	targetX := strings.Split(split[0], "=")
	targetX = strings.Split(targetX[1], "..")
	targetY := strings.Split(split[1], "=")
	targetY = strings.Split(targetY[1], "..")

	x0, _ := strconv.Atoi(targetX[0])
	x1, _ := strconv.Atoi(targetX[1])
	y0, _ := strconv.Atoi(targetY[0])
	y1, _ := strconv.Atoi(targetY[1])

	part1(x0, x1, y0, y1)
	part2(x0, x1, y0, y1)
}

func part1(x0, x1, y0, y1 int) {
	highest := y0 * (y0 + 1) / 2
	fmt.Println("part1:", highest)
}

func part2(xStart, xEnd, yStart, yEnd int) {
	x0 := xStart
	x1 := xEnd
	y0 := yStart
	y1 := yEnd

	count := getCount(x0, x1, y0, y1)

	x1 = x0 - 1
	x0 = int(math.Ceil(math.Sqrt(float64(x0*2)))) - 1

	y1 = -y0 - 1
	y0 = yEnd + 1

	for x := x0; x <= x1; x++ {
		for y := y0; y <= y1; y++ {
			if simulation(x, y, xStart, xEnd, yStart, yEnd) {
				count++
			}
		}
	}

	fmt.Println("part2:", count)
}

func getCount(x0, x1, y0, y1 int) int {
	return (x1 - x0 + 1) * (y1 - y0 + 1)
}

func simulation(x, y, xStart, xEnd, yStart, yEnd int) bool {
	vX := 0
	vY := 0

	for i, j := x, y; ; i, j = i-1, j-1 {
		if i < 0 {
			i = 0
		}

		vX += i
		vY += j

		if vX >= xStart && vX <= xEnd && vY >= yStart && vY <= yEnd {
			return true
		} else if vX > xEnd || vY < yStart {
			return false
		}
	}

	return false
}
