package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	defer f.Close()

	scanner := bufio.NewScanner(f)

	row0 := []int{}
	row1 := []int{}
	row2 := []int{}

	if scanner.Scan() {
		row0 = string2Int(scanner.Text())
	} else {
		log.Fatalf("Error while reading file: %s", scanner.Err())
	}

	if scanner.Scan() {
		row1 = string2Int(scanner.Text())
	} else {
		log.Fatalf("Error while reading file: %s", scanner.Err())
	}

	lowValues := []int{}
	rowLength := len(row0)

	for i := 1; i < rowLength-1; i++ {
		if row0[i] < row0[i+1] {
			if row0[i-1] > row0[i] && row0[i] < row1[i] {
				lowValues = append(lowValues, row0[i]+1)
			}
			i++
		}
	}

	for scanner.Scan() {
		row2 = string2Int(scanner.Text())

		for i := 1; i < rowLength-1; i++ {
			if row1[i] < row1[i+1] {
				if row1[i-1] > row1[i] && row1[i] < row2[i] && row0[i] > row1[i] {
					lowValues = append(lowValues, row1[i]+1)
				}
				i++
			}
		}

		row0 = row1
		row1 = row2
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	for i := 1; i < rowLength-1; i++ {
		if row1[i] < row1[i+1] {
			if row1[i-1] > row1[i] && row1[i] < row0[i] {
				lowValues = append(lowValues, row1[i]+1)
			}
			i++
		}
	}

	sum := 0
	for _, v := range lowValues {
		sum += v
	}

	fmt.Println("part1: ", sum)
}

type point struct {
	x int
	y int
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	heightMap := [][]int{}

	for scanner.Scan() {
		heightMap = append(heightMap, string2Int(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	lowLocation := []point{}

	rowLength := len(heightMap[0])

	for i := 1; i < rowLength-1; i++ {
		if heightMap[0][i] < heightMap[0][i+1] {
			if heightMap[0][i-1] > heightMap[0][i] && heightMap[0][i] < heightMap[1][i] {
				lowLocation = append(lowLocation, point{y: 0, x: i})
			}
			i++
		}
	}

	for row := 1; row < len(heightMap)-1; row++ {
		for i := 1; i < rowLength-1; i++ {
			if heightMap[row][i] < heightMap[row][i+1] {
				if heightMap[row][i-1] > heightMap[row][i] && heightMap[row][i] < heightMap[row+1][i] && heightMap[row-1][i] > heightMap[row][i] {
					lowLocation = append(lowLocation, point{y: row, x: i})
				}
				i++
			}
		}
	}

	rowLast := len(heightMap) - 1
	for i := 1; i < rowLength-1; i++ {
		if heightMap[rowLast][i] < heightMap[rowLast][i+1] {
			if heightMap[rowLast][i-1] > heightMap[rowLast][i] && heightMap[rowLast][i] < heightMap[rowLast-1][i] {
				lowLocation = append(lowLocation, point{y: rowLast, x: i})
			}
			i++
		}
	}

	basinSize := []int{}
	for _, v := range lowLocation {
		value := heightMap[v.y][v.x]

		heightMap[v.y][v.x] = 9

		count := 1

		if value < 8 {
			count += checkHeightMap(value+1, v.x-1, v.y, heightMap)
			count += checkHeightMap(value+1, v.x+1, v.y, heightMap)
			count += checkHeightMap(value+1, v.x, v.y-1, heightMap)
			count += checkHeightMap(value+1, v.x, v.y+1, heightMap)
		}

		basinSize = append(basinSize, count)
	}

	sort.Slice(basinSize, func(i, j int) bool {
		return basinSize[i] > basinSize[j]
	})

	result := basinSize[0] * basinSize[1] * basinSize[2]

	fmt.Println("part2: ", result)
}

func checkHeightMap(pivot, x, y int, heightMap [][]int) int {
	if x < 0 || x > len(heightMap[0])-1 {
		return 0
	}

	if y < 0 || y > len(heightMap)-1 {
		return 0
	}

	if heightMap[y][x] == 9 {
		return 0
	}

	if heightMap[y][x] >= pivot {
		heightMap[y][x] = 9

		count := 1

		if pivot < 8 {
			count += checkHeightMap(pivot+1, x-1, y, heightMap)
			count += checkHeightMap(pivot+1, x+1, y, heightMap)
			count += checkHeightMap(pivot+1, x, y-1, heightMap)
			count += checkHeightMap(pivot+1, x, y+1, heightMap)
		}

		return count
	}

	return 0
}

func string2Int(value string) []int {
	intValue := make([]int, len(value)+2)

	intValue[0] = 9
	for i, v := range value {
		intValue[i+1], _ = strconv.Atoi(string(v))
	}

	intValue[len(intValue)-1] = 9

	return intValue
}
