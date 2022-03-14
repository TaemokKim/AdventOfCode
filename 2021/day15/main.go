package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type coord struct {
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

	cavern := [][]int{}
	for scanner.Scan() {
		cavernLine := []int{}
		for _, v := range scanner.Text() {
			risk, _ := strconv.Atoi(string(v))
			cavernLine = append(cavernLine, risk)
		}

		cavern = append(cavern, cavernLine)
	}

	lowestTotalRisk := sumRisk(cavern)

	fmt.Println("part1:", lowestTotalRisk)

	cavern = appendCavern(cavern)
	lowestTotalRisk = sumRisk(cavern)

	fmt.Println("part2:", lowestTotalRisk)
}

func sumRisk(cavern [][]int) int {
	risk := make([][]int, len(cavern))
	for y := range risk {
		risk[y] = make([]int, len(cavern[y]))
		for x := range risk[y] {
			risk[y][x] = math.MaxInt
		}
	}

	risk[0][0] = cavern[0][0]

	searchList := []coord{}
	searchList = append(searchList, coord{0, 0})

	for len(searchList) > 0 {
		c := searchList[0]
		x := c.x
		y := c.y

		if x-1 >= 0 {
			r := risk[y][x] + cavern[y][x-1]
			if risk[y][x-1] > r {
				risk[y][x-1] = r
				searchList = append(searchList, coord{x - 1, y})
			}
		}

		if y-1 >= 0 {
			r := risk[y][x] + cavern[y-1][x]
			if risk[y-1][x] > r {
				risk[y-1][x] = r
				searchList = append(searchList, coord{x, y - 1})
			}
		}

		if x+1 < len(cavern[0]) {
			r := risk[y][x] + cavern[y][x+1]
			if risk[y][x+1] > r {
				risk[y][x+1] = r
				searchList = append(searchList, coord{x + 1, y})
			}
		}

		if y+1 < len(cavern) {
			r := risk[y][x] + cavern[y+1][x]
			if risk[y+1][x] > r {
				risk[y+1][x] = r
				searchList = append(searchList, coord{x, y + 1})
			}
		}

		searchList = searchList[1:]
	}

	return risk[len(risk)-1][len(risk[0])-1] - risk[0][0]
}

func appendCavern(cavern [][]int) [][]int {
	hCavern := [][]int{}

	for _, line := range cavern {
		newLine := []int{}

		for i := 0; i < 5; i++ {
			for _, v := range line {
				r := v + i
				if r > 9 {
					r = r % 9
				}
				newLine = append(newLine, r)
			}
		}

		hCavern = append(hCavern, newLine)
	}

	vCavern := [][]int{}

	for i := 0; i < 5; i++ {
		for _, line := range hCavern {
			cavernLine := []int{}
			for _, v := range line {
				r := v + i
				if r > 9 {
					r = r % 9
				}

				cavernLine = append(cavernLine, r)
			}
			vCavern = append(vCavern, cavernLine)
		}
	}

	return vCavern
}
