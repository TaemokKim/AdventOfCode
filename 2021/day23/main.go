package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type amphipods struct {
	dst  int
	move int
	fix  bool
}

func main() {
	f, err := os.Open("input.txt")
	//f, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	diagram := [][]amphipods{}
	diagram = append(diagram, make([]amphipods, 11))

	for s.Scan() {
		r, _ := regexp.Compile("#[A-Z]#[A-Z]#[A-Z]#[A-Z]#")
		input := r.FindString(s.Text())
		if len(input) > 0 {
			data := strings.Split(input, "#")

			list := []amphipods{}
			list = append(list, amphipods{})
			list = append(list, amphipods{})

			for _, a := range data {
				switch a {
				case "A":
					list = append(list, amphipods{2, 0, false})
					list = append(list, amphipods{})
				case "B":
					list = append(list, amphipods{4, 0, false})
					list = append(list, amphipods{})
				case "C":
					list = append(list, amphipods{6, 0, false})
					list = append(list, amphipods{})
				case "D":
					list = append(list, amphipods{8, 0, false})
					list = append(list, amphipods{})
				}
			}
			list = append(list, amphipods{})

			diagram = append(diagram, list)
		}
	}

	for x := 2; x < 10; x++ {
		if x%2 > 0 {
			continue
		}

		for y := len(diagram) - 1; y > 0; y-- {
			if diagram[y][x].dst == x {
				diagram[y][x].fix = true
			} else {
				break
			}
		}
	}

	hashMap := make(map[string]bool)
	energy := calcEnergy(diagram, hashMap, math.MaxInt)

	fmt.Println(energy)
}

func calcEnergy(diagram [][]amphipods, hashMap map[string]bool, energyLimit int) int {
	queue := [][][]amphipods{}
	queue = append(queue, diagram)

	minEnergy := energyLimit

	for len(queue) > 0 {
		d := queue[0]
		queue = queue[1:]

		if len(queue) > 1 {
			energy := calcEnergy(d, hashMap, minEnergy)

			if minEnergy > energy {
				minEnergy = energy
			}

			continue
		}

		h := getHash(d)
		_, exists := hashMap[h]
		if exists {
			continue
		} else {
			hashMap[h] = true
		}

		energy := getEnergy(d)

		if energy > minEnergy {
			continue
		}

		if checkSorted(d) {
			minEnergy = energy
			continue
		}

		for x := 0; x < 4; x++ {
			for y := 1; y < len(diagram); y++ {
				indexX := x*2 + 2

				if d[y][indexX].fix || d[y][indexX].dst == 0 {
					continue
				}

				block := false
				for y2 := y - 1; y2 > 0; y2-- {
					if d[y2][indexX].dst > 0 {
						block = true
						break
					}
				}

				if block {
					continue
				}

				for x2 := indexX + 1; x2 < 11; x2++ {
					if (x2 < 10) && (x2%2 == 0) {
						continue
					}

					if d[0][x2].dst > 0 {
						break
					}

					d2 := deepCopy(d)

					d2[0][x2] = d2[y][indexX]
					d2[0][x2].move = x2 - indexX + y
					d2[y][indexX] = amphipods{}

					queue = append(queue, d2)
				}

				for x2 := indexX - 1; x2 >= 0; x2-- {
					if (x2 > 1) && (x2%2 == 0) {
						continue
					}
					if d[0][x2].dst > 0 {
						break
					}

					d2 := deepCopy(d)

					d2[0][x2] = d2[y][indexX]
					d2[0][x2].move = indexX - x2 + y
					d2[y][indexX] = amphipods{}

					queue = append(queue, d2)
				}
			}
		}

		for x := 0; x < 11; x++ {
			if d[0][x].dst == 0 {
				continue
			}

			dst := d[0][x].dst

			block := false
			if dst > x {
				for x2 := x + 1; x2 < dst; x2++ {
					if d[0][x2].dst > 0 {
						block = true
						break
					}
				}
			} else {
				for x2 := x - 1; x2 > dst; x2-- {
					if d[0][x2].dst > 0 {
						block = true
						break
					}
				}
			}

			if block {
				continue
			}

			for y := len(d) - 1; y > 0; y-- {
				if d[y][dst].dst == 0 {
					d2 := deepCopy(d)

					d2[y][dst] = d2[0][x]
					d2[y][dst].move += int(math.Abs(float64(dst-x))) + y
					d2[y][dst].fix = true
					d2[0][x] = amphipods{}

					queue = append(queue, d2)

					break
				} else if !d[y][dst].fix {
					break
				}
			}
		}
	}

	return minEnergy
}

func getHash(d [][]amphipods) string {
	hash := ""
	for y := 0; y < len(d); y++ {
		for x := 0; x < len(d[y]); x++ {
			hash += strconv.Itoa(d[y][x].dst)
			hash += strconv.Itoa(d[y][x].move)
		}
	}

	return hash
}

func deepCopy(d [][]amphipods) [][]amphipods {
	d2 := make([][]amphipods, len(d))
	for i := 0; i < len(d); i++ {
		d2[i] = make([]amphipods, len(d[i]))
		copy(d2[i], d[i])
	}

	return d2
}

func checkSorted(d [][]amphipods) bool {
	isSorted := true
	for y := 1; y < len(d); y++ {
		isSorted = isSorted && (d[y][2].dst == 2) && (d[y][4].dst == 4) && (d[y][6].dst == 6) && (d[y][8].dst == 8)
	}

	return isSorted
}

func getEnergy(d [][]amphipods) int {
	totalEnergy := 0
	for y := 1; y < len(d); y++ {
		totalEnergy += d[y][2].move*1 +
			d[y][4].move*10 +
			d[y][6].move*100 +
			d[y][8].move*1000
	}

	return totalEnergy
}
