package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type cube struct {
	on   bool
	minX int
	maxX int
	minY int
	maxY int
	minZ int
	maxZ int
}

func main() {
	f, err := os.Open("input.txt")
	//f, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	cubeOnList := []cube{}
	cubeAllOnList := []cube{}

	for s.Scan() {
		newCube := makeCube(s.Text())
		cubeAllOnList = checkCube(cubeAllOnList, newCube)

		outOfBounds, c := checkBound(newCube)
		if !outOfBounds {
			continue
		}

		cubeOnList = checkCube(cubeOnList, c)
	}

	count := 0
	for _, c := range cubeOnList {
		count += (c.maxX - c.minX + 1) * (c.maxY - c.minY + 1) * (c.maxZ - c.minZ + 1)
	}

	fmt.Println("on region", count)

	count = 0
	for _, c := range cubeAllOnList {
		count += (c.maxX - c.minX + 1) * (c.maxY - c.minY + 1) * (c.maxZ - c.minZ + 1)
	}

	fmt.Println("on all", count)
}

func makeCube(input string) cube {
	data := strings.Split(input, " ")
	on := bool(data[0] == "on")

	coord := strings.Split(data[1], ",")

	x := strings.Split(coord[0], "=")
	x = strings.Split(x[1], "..")
	minX, _ := strconv.Atoi(x[0])
	maxX, _ := strconv.Atoi(x[1])

	y := strings.Split(coord[1], "=")
	y = strings.Split(y[1], "..")
	minY, _ := strconv.Atoi(y[0])
	maxY, _ := strconv.Atoi(y[1])

	z := strings.Split(coord[2], "=")
	z = strings.Split(z[1], "..")
	minZ, _ := strconv.Atoi(z[0])
	maxZ, _ := strconv.Atoi(z[1])

	return cube{on, minX, maxX, minY, maxY, minZ, maxZ}
}

func checkBound(c cube) (bool, cube) {
	if !checkLimit(c.minX, c.maxX) {
		return false, c
	}

	if !checkLimit(c.minY, c.maxY) {
		return false, c
	}

	if !checkLimit(c.minZ, c.maxZ) {
		return false, c
	}

	if c.minX < -50 {
		c.minX = -50
	}

	if c.maxX > 50 {
		c.maxX = 50
	}

	if c.minY < -50 {
		c.minY = -50
	}

	if c.maxY > 50 {
		c.maxY = 50
	}

	if c.minZ < -50 {
		c.minZ = -50
	}

	if c.maxZ > 50 {
		c.maxZ = 50
	}

	return true, c
}

func checkLimit(min, max int) bool {
	if max < -50 || min > 50 {
		return false
	}

	return true
}

func checkCube(cubeList []cube, newCube cube) []cube {
	if newCube.on {
		return turnOn(cubeList, newCube)
	} else {
		return turnOff(cubeList, newCube)
	}
}

func turnOn(cubeList []cube, newCube cube) []cube {
	newCubeList := []cube{newCube}

	for _, c := range cubeList {
		tempCubeList := []cube{}
		for _, newC := range newCubeList {
			oV, oCube := isOverlap(c, newC)
			if oV {
				if newC.minX == oCube.minX && newC.maxX > oCube.maxX {
					cubeX := cube{true, oCube.maxX + 1, newC.maxX, newC.minY, newC.maxY, newC.minZ, newC.maxZ}
					tempCubeList = append(tempCubeList, cubeX)
				} else if newC.minX < oCube.minX && newC.maxX == oCube.maxX {
					cubeX := cube{true, newC.minX, oCube.minX - 1, newC.minY, newC.maxY, newC.minZ, newC.maxZ}
					tempCubeList = append(tempCubeList, cubeX)
				} else if newC.minX < oCube.minX && newC.maxX > oCube.maxX {
					cubeX := cube{true, newC.minX, oCube.minX - 1, newC.minY, newC.maxY, newC.minZ, newC.maxZ}
					tempCubeList = append(tempCubeList, cubeX)
					cubeX = cube{true, oCube.maxX + 1, newC.maxX, newC.minY, newC.maxY, newC.minZ, newC.maxZ}
					tempCubeList = append(tempCubeList, cubeX)
				}

				if newC.minY == oCube.minY && newC.maxY > oCube.maxY {
					cubeY := cube{true, oCube.minX, oCube.maxX, oCube.maxY + 1, newC.maxY, newC.minZ, newC.maxZ}
					tempCubeList = append(tempCubeList, cubeY)
				} else if newC.minY < oCube.minY && newC.maxY == oCube.maxY {
					cubeY := cube{true, oCube.minX, oCube.maxX, newC.minY, oCube.minY - 1, newC.minZ, newC.maxZ}
					tempCubeList = append(tempCubeList, cubeY)
				} else if newC.minY < oCube.minY && newC.maxY > oCube.maxY {
					cubeY := cube{true, oCube.minX, oCube.maxX, newC.minY, oCube.minY - 1, newC.minZ, newC.maxZ}
					tempCubeList = append(tempCubeList, cubeY)
					cubeY = cube{true, oCube.minX, oCube.maxX, oCube.maxY + 1, newC.maxY, newC.minZ, newC.maxZ}
					tempCubeList = append(tempCubeList, cubeY)
				}

				if newC.minZ == oCube.minZ && newC.maxZ > oCube.maxZ {
					cubeZ := cube{true, oCube.minX, oCube.maxX, oCube.minY, oCube.maxY, oCube.maxZ + 1, newC.maxZ}
					tempCubeList = append(tempCubeList, cubeZ)
				} else if newC.minZ < oCube.minZ && newC.maxZ == oCube.maxZ {
					cubeZ := cube{true, oCube.minX, oCube.maxX, oCube.minY, oCube.maxY, newC.minZ, oCube.minZ - 1}
					tempCubeList = append(tempCubeList, cubeZ)
				} else if newC.minZ < oCube.minZ && newC.maxZ > oCube.maxZ {
					cubeZ := cube{true, oCube.minX, oCube.maxX, oCube.minY, oCube.maxY, newC.minZ, oCube.minZ - 1}
					tempCubeList = append(tempCubeList, cubeZ)
					cubeZ = cube{true, oCube.minX, oCube.maxX, oCube.minY, oCube.maxY, oCube.maxZ + 1, newC.maxZ}
					tempCubeList = append(tempCubeList, cubeZ)
				}
			} else {
				tempCubeList = append(tempCubeList, newC)
			}
		}

		newCubeList = tempCubeList
	}

	cubeList = append(cubeList, newCubeList...)

	return cubeList
}

func turnOff(cubeList []cube, newCube cube) []cube {
	newCubeList := []cube{}

	for _, c := range cubeList {
		oV, oCube := isOverlap(c, newCube)
		if oV {
			if c.minX == oCube.minX && c.maxX > oCube.maxX {
				cubeX := cube{true, oCube.maxX + 1, c.maxX, c.minY, c.maxY, c.minZ, c.maxZ}
				newCubeList = append(newCubeList, cubeX)
			} else if c.minX < oCube.minX && c.maxX == oCube.maxX {
				cubeX := cube{true, c.minX, oCube.minX - 1, c.minY, c.maxY, c.minZ, c.maxZ}
				newCubeList = append(newCubeList, cubeX)
			} else if c.minX < oCube.minX && c.maxX > oCube.maxX {
				cubeX := cube{true, c.minX, oCube.minX - 1, c.minY, c.maxY, c.minZ, c.maxZ}
				newCubeList = append(newCubeList, cubeX)
				cubeX = cube{true, oCube.maxX + 1, c.maxX, c.minY, c.maxY, c.minZ, c.maxZ}
				newCubeList = append(newCubeList, cubeX)
			}

			if c.minY == oCube.minY && c.maxY > oCube.maxY {
				cubeY := cube{true, oCube.minX, oCube.maxX, oCube.maxY + 1, c.maxY, c.minZ, c.maxZ}
				newCubeList = append(newCubeList, cubeY)
			} else if c.minY < oCube.minY && c.maxY == oCube.maxY {
				cubeY := cube{true, oCube.minX, oCube.maxX, c.minY, oCube.minY - 1, c.minZ, c.maxZ}
				newCubeList = append(newCubeList, cubeY)
			} else if c.minY < oCube.minY && c.maxY > oCube.maxY {
				cubeY := cube{true, oCube.minX, oCube.maxX, c.minY, oCube.minY - 1, c.minZ, c.maxZ}
				newCubeList = append(newCubeList, cubeY)
				cubeY = cube{true, oCube.minX, oCube.maxX, oCube.maxY + 1, c.maxY, c.minZ, c.maxZ}
				newCubeList = append(newCubeList, cubeY)
			}

			if c.minZ == oCube.minZ && c.maxZ > oCube.maxZ {
				cubeZ := cube{true, oCube.minX, oCube.maxX, oCube.minY, oCube.maxY, oCube.maxZ + 1, c.maxZ}
				newCubeList = append(newCubeList, cubeZ)
			} else if c.minZ < oCube.minZ && c.maxZ == oCube.maxZ {
				cubeZ := cube{true, oCube.minX, oCube.maxX, oCube.minY, oCube.maxY, c.minZ, oCube.minZ - 1}
				newCubeList = append(newCubeList, cubeZ)
			} else if c.minZ < oCube.minZ && c.maxZ > oCube.maxZ {
				cubeZ := cube{true, oCube.minX, oCube.maxX, oCube.minY, oCube.maxY, c.minZ, oCube.minZ - 1}
				newCubeList = append(newCubeList, cubeZ)
				cubeZ = cube{true, oCube.minX, oCube.maxX, oCube.minY, oCube.maxY, oCube.maxZ + 1, c.maxZ}
				newCubeList = append(newCubeList, cubeZ)
			}
		} else {
			newCubeList = append(newCubeList, c)
		}
	}

	return newCubeList
}

func isOverlap(c1, c2 cube) (bool, cube) {
	overlap := cube{}
	oX := false
	oY := false
	oZ := false

	overlap.minX, overlap.maxX, oX = getOverlap(c1.minX, c1.maxX, c2.minX, c2.maxX)
	overlap.minY, overlap.maxY, oY = getOverlap(c1.minY, c1.maxY, c2.minY, c2.maxY)
	overlap.minZ, overlap.maxZ, oZ = getOverlap(c1.minZ, c1.maxZ, c2.minZ, c2.maxZ)

	return oX && oY && oZ, overlap
}

func getOverlap(min1, max1, min2, max2 int) (int, int, bool) {
	if min1 <= min2 && max1 >= min2 {
		if max2 < max1 {
			return min2, max2, true
		}

		return min2, max1, true
	} else if min2 <= min1 && max2 >= min1 {
		if max1 < max2 {
			return min1, max1, true
		}

		return min1, max2, true
	}

	return 0, 0, false
}
