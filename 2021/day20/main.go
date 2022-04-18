package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	//	f, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	algo := make(map[int]int)

	index := 0

	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			break
		}

		for _, v := range line {
			if v == '.' {
				algo[index] = 0
			} else if v == '#' {
				algo[index] = 1
			}

			index++
		}
	}

	inputImage := [][]int{}
	for s.Scan() {
		inputLine := []int{0, 0}

		for _, v := range s.Text() {
			if v == '.' {
				inputLine = append(inputLine, 0)
			} else if v == '#' {
				inputLine = append(inputLine, 1)
			}
		}

		inputLine = append(inputLine, []int{0, 0}...)

		inputImage = append(inputImage, inputLine)
	}

	padding := make([]int, len(inputImage[0]))

	inputImage = append([][]int{padding}, inputImage...)
	inputImage = append([][]int{padding}, inputImage...)
	inputImage = append(inputImage, padding)
	inputImage = append(inputImage, padding)

	enhanceImage := enhance(inputImage, algo, 1)
	enhanceImage = enhance(enhanceImage, algo, 2)

	count := getCount(enhanceImage)

	fmt.Println("part1", count)

	for i := 3; i < 51; i++ {
		enhanceImage = enhance(enhanceImage, algo, i)
	}

	count = getCount(enhanceImage)
	fmt.Println("part2", count)
}

func printImage(title string, image [][]int) {
	fmt.Println(title)
	for y := 0; y < len(image); y++ {
		fmt.Println(image[y])
	}
	fmt.Println()
}

func getCount(inputImage [][]int) int {
	count := 0

	for y := 0; y < len(inputImage); y++ {
		for x := 0; x < len(inputImage[y]); x++ {
			count += inputImage[y][x]
		}
	}

	return count
}

func enhance(inputImage [][]int, algo map[int]int, count int) [][]int {
	outputImage := [][]int{}
	paddingValue := 0

	if inputImage[0][0] == 0 {
		paddingValue = algo[0]
	} else {
		paddingValue = algo[len(algo)-1]
	}

	for y := 2; y < len(inputImage); y++ {
		outLine := []int{paddingValue, paddingValue}

		for x := 2; x < len(inputImage[y]); x++ {
			value := strconv.Itoa(inputImage[y-2][x-2]) + strconv.Itoa(inputImage[y-2][x-1]) + strconv.Itoa(inputImage[y-2][x])
			value += strconv.Itoa(inputImage[y-1][x-2]) + strconv.Itoa(inputImage[y-1][x-1]) + strconv.Itoa(inputImage[y-1][x])
			value += strconv.Itoa(inputImage[y][x-2]) + strconv.Itoa(inputImage[y][x-1]) + strconv.Itoa(inputImage[y][x])

			ret, _ := strconv.ParseInt(value, 2, 64)

			outLine = append(outLine, algo[int(ret)])
		}

		outLine = append(outLine, []int{paddingValue, paddingValue}...)

		outputImage = append(outputImage, outLine)
	}

	padding := make([]int, len(outputImage[0]))
	for i, _ := range padding {
		padding[i] = paddingValue
	}

	outputImage = append([][]int{padding}, outputImage...)
	outputImage = append([][]int{padding}, outputImage...)
	outputImage = append(outputImage, padding)
	outputImage = append(outputImage, padding)

	return outputImage
}
