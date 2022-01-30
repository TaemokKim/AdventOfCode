package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
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

	scanner := bufio.NewScanner(f)

	count := 0
	for scanner.Scan() {
		digits := strings.Split(scanner.Text(), "|")

		fourDigits := strings.Split(digits[1], " ")

		for _, v := range fourDigits {
			length := len(v)

			if length == 2 || length == 3 || length == 4 || length == 7 {
				count += 1
			}
		}
	}

	fmt.Println("part1: ", count)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	sum := 0
	for scanner.Scan() {
		digits := strings.Split(scanner.Text(), "|")

		tenDigits := strings.Split(digits[0], " ")
		fourDigits := strings.Split(digits[1], " ")

		list := [10]string{}
		for _, v := range tenDigits {
			if len(v) == 2 {
				list[1] = v
			} else if len(v) == 3 {
				list[7] = v
			} else if len(v) == 4 {
				list[4] = v
			} else if len(v) == 7 {
				list[8] = v
			}
		}

		unmatched_4_1 := ""
		for _, c := range list[4] {
			matched := false
			for _, c2 := range list[1] {
				if c == c2 {
					matched = true

					break
				}
			}

			if matched == false {
				unmatched_4_1 = unmatched_4_1 + string(c)
			}
		}

		for _, v := range tenDigits {
			if len(v) == 5 {
				if matchLen(v, list[1]) == 2 {
					list[3] = v
				} else {
					count := 0
					for _, c := range unmatched_4_1 {
						for _, c2 := range v {
							if c2 == c {
								count += 1
								break
							}
						}
					}

					if count == 2 {
						list[5] = v
					} else {
						list[2] = v
					}
				}
			}
		}

		for _, v := range tenDigits {
			if len(v) == 6 {
				count := 0
				matched := false
				for _, c := range v {
					matched = false
					for _, c2 := range list[1] {
						if c == c2 {
							matched = true
							break
						}
					}

					if matched {
						count += 1

					}
				}
				if count == 1 {
					list[6] = v
				} else {
					count = 0

					for _, c := range v {
						matched = false
						for _, c2 := range list[3] {
							if c == c2 {
								matched = true
								break
							}
						}

						if matched {
							count += 1

						}
					}
					if count == 5 {
						list[9] = v
					} else {
						list[0] = v
					}
				}

			}
		}

		for i, v := range list {
			list[i] = sortString(v)
		}

		appendValue := ""
		for _, v := range fourDigits {
			if len(v) == 0 {
				continue
			}

			value := sortString(v)

			for i, v2 := range list {
				if value == v2 {
					appendValue = appendValue + strconv.Itoa(i)
					break
				}
			}
		}

		fourValue, _ := strconv.Atoi(appendValue)
		sum += fourValue
	}

	fmt.Println("part2: ", sum)
}

func sortString(input string) string {
	runeList := []rune(input)
	chars := []string{}
	for _, c := range runeList {
		chars = append(chars, string(c))
	}
	sort.Strings(chars)

	return strings.Join(chars, "")
}

func matchLen(target, value string) int {
	count := 0

	for _, c := range value {
		for _, c2 := range target {
			if c == c2 {
				count += 1
				break
			}
		}
	}

	return count
}
