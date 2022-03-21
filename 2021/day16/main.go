package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	VERSION = 1 + iota
	TYPE
	LENGTH_TYPE
	NUMBER_SUB
	LENGTH_SUB
	SUB_PACKET
	LITERAL
)

var hex2dec map[string][]string

type parseResult struct {
	sumVersion  int64
	totalLength int64
	packet      []string
	value       int64
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
	scanner.Split(bufio.ScanRunes)

	hex2dec = make(map[string][]string)
	hex2dec["0"] = []string{"0", "0", "0", "0"}
	hex2dec["1"] = []string{"0", "0", "0", "1"}
	hex2dec["2"] = []string{"0", "0", "1", "0"}
	hex2dec["3"] = []string{"0", "0", "1", "1"}
	hex2dec["4"] = []string{"0", "1", "0", "0"}
	hex2dec["5"] = []string{"0", "1", "0", "1"}
	hex2dec["6"] = []string{"0", "1", "1", "0"}
	hex2dec["7"] = []string{"0", "1", "1", "1"}
	hex2dec["8"] = []string{"1", "0", "0", "0"}
	hex2dec["9"] = []string{"1", "0", "0", "1"}
	hex2dec["A"] = []string{"1", "0", "1", "0"}
	hex2dec["B"] = []string{"1", "0", "1", "1"}
	hex2dec["C"] = []string{"1", "1", "0", "0"}
	hex2dec["D"] = []string{"1", "1", "0", "1"}
	hex2dec["E"] = []string{"1", "1", "1", "0"}
	hex2dec["F"] = []string{"1", "1", "1", "1"}

	packet := []string{}

	result := parsePacket(scanner, packet)

	fmt.Println("part1:", result.sumVersion)
	fmt.Println("part2:", result.value)
}

func parsePacket(scanner *bufio.Scanner, packet []string) parseResult {
	status := VERSION
	sumVersion := int64(0)
	numberSub := int64(0)
	lengthSub := int64(0)

	totalLength := int64(0)
	valueString := []string{}
	value := int64(0)

	last := false

	resultParsing := parseResult{sumVersion, totalLength, []string{}, value}

	operand := func(a, b int64) int64 {
		return a
	}

	for true {
		switch status {
		case VERSION:
			if len(packet) >= 3 {
				totalLength += 3
				sumVersion += string2Dec(packet[0:3], 2)

				packet = packet[3:]

				status = TYPE

			} else {
				data, ret := readData(scanner)
				if ret {
					packet = append(packet, data...)
				} else {
					last = true
				}
			}
		case TYPE:
			if len(packet) >= 3 {
				totalLength += 3
				packetType := string2Dec(packet[0:3], 2)

				packet = packet[3:]
				switch packetType {
				case 0:
					operand = func(a, b int64) int64 {
						return a + b
					}

					status = LENGTH_TYPE
				case 1:
					operand = func(a, b int64) int64 {
						return a * b
					}

					status = LENGTH_TYPE
				case 2:
					operand = func(a, b int64) int64 {
						if a < b {
							return a
						} else {
							return b
						}
					}

					status = LENGTH_TYPE
				case 3:
					operand = func(a, b int64) int64 {
						if a > b {
							return a
						} else {
							return b
						}
					}

					status = LENGTH_TYPE
				case 4:
					status = LITERAL
				case 5:
					operand = func(a, b int64) int64 {
						if a > b {
							return 1
						} else {
							return 0
						}
					}

					status = LENGTH_TYPE
				case 6:
					operand = func(a, b int64) int64 {
						if a < b {
							return 1
						} else {
							return 0
						}
					}

					status = LENGTH_TYPE
				case 7:
					operand = func(a, b int64) int64 {
						if a == b {
							return 1
						} else {
							return 0
						}
					}

					status = LENGTH_TYPE
				}

			} else {
				data, ret := readData(scanner)
				if ret {
					packet = append(packet, data...)
				} else {
					last = true
				}
			}

		case LENGTH_TYPE:
			if len(packet) >= 1 {
				totalLength += 1
				lengthType := string2Dec(packet[0:1], 2)

				packet = packet[1:]

				if lengthType == 0 {
					status = LENGTH_SUB
				} else if lengthType == 1 {
					status = NUMBER_SUB
				} else {
					fmt.Println("lengthType error:", lengthType)
					return resultParsing
				}
			} else {
				data, ret := readData(scanner)
				if ret {
					packet = append(packet, data...)
				} else {
					last = true
				}
			}

		case NUMBER_SUB:
			if len(packet) >= 11 {
				totalLength += 11
				numberSub = string2Dec(packet[0:11], 2)

				status = SUB_PACKET
				result := parsePacket(scanner, packet[11:])
				sumVersion += result.sumVersion
				totalLength += result.totalLength
				packet = result.packet
				value = result.value

				numberSub--
				if numberSub == 0 {
					last = true
				}
			} else {
				data, ret := readData(scanner)
				if ret {
					packet = append(packet, data...)
				} else {
					last = true
				}
			}

		case LENGTH_SUB:
			if len(packet) >= 15 {
				totalLength += 15
				lengthSub = string2Dec(packet[0:15], 2)

				status = SUB_PACKET
				result := parsePacket(scanner, packet[15:])
				sumVersion += result.sumVersion
				totalLength += result.totalLength
				packet = result.packet
				value = result.value

				lengthSub -= result.totalLength
				if lengthSub < 1 {
					last = true
				}
			} else {
				data, ret := readData(scanner)
				if ret {
					packet = append(packet, data...)
				} else {
					last = true
				}
			}

		case SUB_PACKET:
			if numberSub > 0 {
				result := parsePacket(scanner, packet)
				sumVersion += result.sumVersion
				totalLength += result.totalLength
				packet = result.packet
				value = calc(operand, value, result.value)

				numberSub--
				if numberSub == 0 {
					last = true
				}
			} else if lengthSub > 0 {
				result := parsePacket(scanner, packet)
				sumVersion += result.sumVersion
				totalLength += result.totalLength
				packet = result.packet
				value = calc(operand, value, result.value)

				lengthSub -= result.totalLength
				if lengthSub < 1 {
					last = true
				}
			} else {
				last = true
			}
		case LITERAL:
			if len(packet) > 4 {
				totalLength += 5

				if packet[0] == "0" {
					last = true
				}

				valueString = append(valueString, packet[1:5]...)

				packet = packet[5:]
			} else {
				data, ret := readData(scanner)
				if ret {
					packet = append(packet, data...)
				} else {
					last = true
				}
			}
		}

		if last {
			if status == LITERAL {
				value = string2Dec(valueString, 2)
			}

			break
		}
	}

	resultParsing.sumVersion = sumVersion
	resultParsing.totalLength = totalLength
	resultParsing.packet = packet
	resultParsing.value = value

	return resultParsing
}

func string2Dec(packet []string, base int) int64 {
	packetString := strings.Join(packet, "")
	ret, _ := strconv.ParseInt(packetString, base, 64)

	return ret
}

func readData(scanner *bufio.Scanner) ([]string, bool) {
	ret := scanner.Scan()
	hex := string(scanner.Text())
	value := hex2dec[hex]

	return value, ret
}

func calc(f func(int64, int64) int64, a, b int64) int64 {
	return f(a, b)
}
