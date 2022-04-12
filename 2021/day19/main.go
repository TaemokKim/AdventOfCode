package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type beacon struct {
	x int
	y int
	z int
}

type scanner struct {
	diff       beacon
	beaconList []beacon
}

func (s *scanner) getRotation(index int) scanner {
	newS := scanner{}
	beaconList := []beacon{}

	switch index {
	case 0:
		for _, b := range s.beaconList {
			beacon := beacon{b.x, b.y, b.z}
			beaconList = append(beaconList, beacon)
		}
	case 1:
		for _, b := range s.beaconList {
			beacon := beacon{b.y, b.x, b.z}
			beaconList = append(beaconList, beacon)
		}
	case 2:
		for _, b := range s.beaconList {
			beacon := beacon{b.z, b.y, b.x}
			beaconList = append(beaconList, beacon)
		}
	case 3:
		for _, b := range s.beaconList {
			beacon := beacon{b.x, b.z, b.y}
			beaconList = append(beaconList, beacon)
		}
	case 4:
		for _, b := range s.beaconList {
			beacon := beacon{-b.x, b.y, b.z}
			beaconList = append(beaconList, beacon)
		}
	case 5:
		for _, b := range s.beaconList {
			beacon := beacon{b.x, -b.y, b.z}
			beaconList = append(beaconList, beacon)
		}
	case 6:
		for _, b := range s.beaconList {
			beacon := beacon{b.x, b.y, -b.z}
			beaconList = append(beaconList, beacon)
		}
	case 7:
		for _, b := range s.beaconList {
			beacon := beacon{-b.x, -b.y, b.z}
			beaconList = append(beaconList, beacon)
		}
	case 8:
		for _, b := range s.beaconList {
			beacon := beacon{-b.x, b.y, -b.z}
			beaconList = append(beaconList, beacon)
		}
	case 9:
		for _, b := range s.beaconList {
			beacon := beacon{b.x, -b.y, -b.z}
			beaconList = append(beaconList, beacon)
		}
	case 10:
		for _, b := range s.beaconList {
			beacon := beacon{-b.y, b.x, b.z}
			beaconList = append(beaconList, beacon)
		}
	case 11:
		for _, b := range s.beaconList {
			beacon := beacon{b.y, -b.x, b.z}
			beaconList = append(beaconList, beacon)
		}
	case 12:
		for _, b := range s.beaconList {
			beacon := beacon{b.y, b.x, -b.z}
			beaconList = append(beaconList, beacon)
		}
	case 13:
		for _, b := range s.beaconList {
			beacon := beacon{-b.y, -b.x, b.z}
			beaconList = append(beaconList, beacon)
		}
	case 14:
		for _, b := range s.beaconList {
			beacon := beacon{-b.y, b.x, -b.z}
			beaconList = append(beaconList, beacon)
		}
	case 15:
		for _, b := range s.beaconList {
			beacon := beacon{b.y, -b.x, -b.z}
			beaconList = append(beaconList, beacon)
		}
	case 16:
		for _, b := range s.beaconList {
			beacon := beacon{-b.z, b.y, b.x}
			beaconList = append(beaconList, beacon)
		}
	case 17:
		for _, b := range s.beaconList {
			beacon := beacon{b.z, -b.y, b.x}
			beaconList = append(beaconList, beacon)
		}
	case 18:
		for _, b := range s.beaconList {
			beacon := beacon{b.z, b.y, -b.x}
			beaconList = append(beaconList, beacon)
		}
	case 19:
		for _, b := range s.beaconList {
			beacon := beacon{-b.z, -b.y, b.x}
			beaconList = append(beaconList, beacon)
		}
	case 20:
		for _, b := range s.beaconList {
			beacon := beacon{-b.z, b.y, -b.x}
			beaconList = append(beaconList, beacon)
		}
	case 21:
		for _, b := range s.beaconList {
			beacon := beacon{b.z, -b.y, -b.x}
			beaconList = append(beaconList, beacon)
		}
	case 22:
		for _, b := range s.beaconList {
			beacon := beacon{-b.x, b.z, b.y}
			beaconList = append(beaconList, beacon)
		}
	case 23:
		for _, b := range s.beaconList {
			beacon := beacon{b.x, -b.z, b.y}
			beaconList = append(beaconList, beacon)
		}
	case 24:
		for _, b := range s.beaconList {
			beacon := beacon{b.x, b.z, -b.y}
			beaconList = append(beaconList, beacon)
		}
	case 25:
		for _, b := range s.beaconList {
			beacon := beacon{-b.x, -b.z, b.y}
			beaconList = append(beaconList, beacon)
		}
	case 26:
		for _, b := range s.beaconList {
			beacon := beacon{-b.x, b.z, -b.y}
			beaconList = append(beaconList, beacon)
		}
	case 27:
		for _, b := range s.beaconList {
			beacon := beacon{b.x, -b.z, -b.y}
			beaconList = append(beaconList, beacon)
		}
	case 28:
		for _, b := range s.beaconList {
			beacon := beacon{-b.y, -b.x, -b.z}
			beaconList = append(beaconList, beacon)
		}
	case 29:
		for _, b := range s.beaconList {
			beacon := beacon{-b.z, -b.y, -b.x}
			beaconList = append(beaconList, beacon)
		}
	case 30:
		for _, b := range s.beaconList {
			beacon := beacon{-b.x, -b.z, -b.y}
			beaconList = append(beaconList, beacon)
		}
	case 31:
		for _, b := range s.beaconList {
			beacon := beacon{-b.x, -b.y, -b.z}
			beaconList = append(beaconList, beacon)
		}
	case 32:
		for _, b := range s.beaconList {
			beacon := beacon{b.z, b.y, b.x}
			beaconList = append(beaconList, beacon)
		}
	case 33:
		for _, b := range s.beaconList {
			beacon := beacon{-b.z, b.y, b.x}
			beaconList = append(beaconList, beacon)
		}
	case 34:
		for _, b := range s.beaconList {
			beacon := beacon{b.z, -b.y, b.x}
			beaconList = append(beaconList, beacon)
		}
	case 35:
		for _, b := range s.beaconList {
			beacon := beacon{b.z, b.y, -b.x}
			beaconList = append(beaconList, beacon)
		}
	case 36:
		for _, b := range s.beaconList {
			beacon := beacon{-b.z, -b.y, b.x}
			beaconList = append(beaconList, beacon)
		}
	case 37:
		for _, b := range s.beaconList {
			beacon := beacon{-b.z, b.y, -b.x}
			beaconList = append(beaconList, beacon)
		}
	case 38:
		for _, b := range s.beaconList {
			beacon := beacon{b.z, -b.y, -b.x}
			beaconList = append(beaconList, beacon)
		}
	case 39:
		for _, b := range s.beaconList {
			beacon := beacon{-b.z, -b.y, -b.x}
			beaconList = append(beaconList, beacon)
		}
	case 40:
		for _, b := range s.beaconList {
			beacon := beacon{b.z, b.x, b.y}
			beaconList = append(beaconList, beacon)
		}
	case 41:
		for _, b := range s.beaconList {
			beacon := beacon{-b.z, b.x, b.y}
			beaconList = append(beaconList, beacon)
		}
	case 42:
		for _, b := range s.beaconList {
			beacon := beacon{b.z, -b.x, b.y}
			beaconList = append(beaconList, beacon)
		}
	case 43:
		for _, b := range s.beaconList {
			beacon := beacon{b.z, b.x, -b.y}
			beaconList = append(beaconList, beacon)
		}
	case 44:
		for _, b := range s.beaconList {
			beacon := beacon{-b.z, -b.x, b.y}
			beaconList = append(beaconList, beacon)
		}
	case 45:
		for _, b := range s.beaconList {
			beacon := beacon{-b.z, b.x, -b.y}
			beaconList = append(beaconList, beacon)
		}
	case 46:
		for _, b := range s.beaconList {
			beacon := beacon{b.z, -b.x, -b.y}
			beaconList = append(beaconList, beacon)
		}
	case 47:
		for _, b := range s.beaconList {
			beacon := beacon{-b.z, -b.x, -b.y}
			beaconList = append(beaconList, beacon)
		}

	}

	newS.diff = s.diff
	newS.beaconList = beaconList
	return newS
}

func main() {
	//f, err := os.Open("input.txt")
	f, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	scannerList := []scanner{}
	for s.Scan() {
		line := s.Text()
		if line[0:3] == "---" {
			beaconList := scanBeacons(s)
			scannerList = append(scannerList, scanner{beacon{0, 0, 0}, beaconList})
		}
	}

	newScannerList := []scanner{}
	newScannerList = append(newScannerList, scannerList[0])

	if len(scannerList) > 1 {
		scannerList = scannerList[1:len(scannerList)]
	}

	scanner0 := newScannerList[0]
	i := 0

	for len(scannerList) > i {
		scanner1 := scannerList[i]

		for j := 0; j < 48; j++ {
			newScanner, matched := compareScanner(scanner0, scanner1.getRotation(j))
			if matched {
				newBeaconList := []beacon{}
				for _, b1 := range newScanner.beaconList {
					found := false
					for _, b0 := range scanner0.beaconList {
						if b1.x == b0.x && b1.y == b0.y && b1.z == b0.z {
							found = true
							break
						}
					}

					if !found {
						newBeaconList = append(newBeaconList, b1)
					}
				}

				newScannerList = append(newScannerList, newScanner)
				scanner0.beaconList = append(scanner0.beaconList, newBeaconList...)
				scannerList = append(scannerList[0:i], scannerList[i+1:len(scannerList)]...)
				i = -1
				break
			}
		}

		i++
	}

	fmt.Println(len(scanner0.beaconList))
}

func scanBeacons(s *bufio.Scanner) []beacon {
	beaconList := []beacon{}

	for s.Scan() {
		xyz := strings.Split(s.Text(), ",")
		if len(xyz) == 3 {
			x, _ := strconv.Atoi(xyz[0])
			y, _ := strconv.Atoi(xyz[1])
			z, _ := strconv.Atoi(xyz[2])
			beaconList = append(beaconList, beacon{x, y, z})
		} else {
			break
		}
	}

	return beaconList
}

func compareScanner(s0, s1 scanner) (scanner, bool) {
	s := scanner{}

	for i := 0; i < len(s0.beaconList); i++ {
		b0 := s0.beaconList[i]

		for j := 0; j < len(s1.beaconList); j++ {
			b1 := s1.beaconList[j]

			diff := beacon{b0.x - b1.x, b0.y - b1.y, b0.z - b1.z}

			count := 0
			for ii := 0; ii < len(s0.beaconList); ii++ {
				b0 = s0.beaconList[ii]

				for jj := 0; jj < len(s1.beaconList); jj++ {
					b1 = s1.beaconList[jj]

					if (b0.x-b1.x == diff.x) && (b0.y-b1.y == diff.y) && (b0.z-b1.z == diff.z) {
						count++
						break
					}
				}
			}

			if count >= 12 {
				beaconList := []beacon{}
				for _, b := range s1.beaconList {
					beaconList = append(beaconList, beacon{b.x + diff.x, b.y + diff.y, b.z + diff.z})
				}
				s.diff = diff
				s.beaconList = beaconList

				return s, true
			}
		}
	}

	return s, false
}
