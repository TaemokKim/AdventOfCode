package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type player struct {
	position int
	score    int
	universe int
}

func main() {
	f, err := os.Open("input.txt")
	//f, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	playerList := []player{}

	for s.Scan() {
		data := strings.Split(s.Text(), ": ")
		position, _ := strconv.Atoi(data[1])

		playerList = append(playerList, player{position, 0, 1})
	}

	part1(playerList)
	part2(playerList)
}

func part1(list []player) {
	playerList := make([]player, len(list))
	copy(playerList, list)

	count := 0
	result := 0
	turn := 0
	for {
		turn++
		count++
		playerList[0].position = (playerList[0].position + 8*turn - 2) % 10
		if playerList[0].position == 0 {
			playerList[0].position = 10
		}

		playerList[0].score += playerList[0].position

		if playerList[0].score >= 1000 {
			result = playerList[1].score * count * 3
			break
		}

		count++
		playerList[1].position = (playerList[1].position + 8*turn - 3) % 10
		if playerList[1].position == 0 {
			playerList[1].position = 10
		}

		playerList[1].score += playerList[1].position

		if playerList[1].score >= 1000 {
			result = playerList[0].score * count * 3
			break
		}
	}

	fmt.Println("part1", result)

}

func part2(list []player) {
	rollResult := []int{3, 4, 5, 6, 7, 8, 9}
	universe := []int{1, 3, 6, 7, 6, 3, 1}

	p1Wins, p2Wins := roll(rollResult, universe, list[0], list[1])

	if p1Wins > p2Wins {
		fmt.Println("part2", p1Wins)
	} else {
		fmt.Println("part2", p2Wins)
	}
}

func roll(rollResult, universe []int, player1, player2 player) (int, int) {
	p1Wins := 0
	p2Wins := 0

	for i := 0; i < len(rollResult); i++ {
		p1 := player1.position + rollResult[i]
		if p1 > 10 {
			p1 -= 10
		}

		s1 := player1.score + p1
		u1 := player1.universe * universe[i]

		if s1 >= 21 {
			p1Wins += (u1 * player2.universe)
			continue
		}

		updateP1 := player{p1, s1, u1}

		for j := 0; j < len(rollResult); j++ {
			p2 := player2.position + rollResult[j]
			if p2 > 10 {
				p2 -= 10
			}

			s2 := player2.score + p2
			u2 := player2.universe * universe[j]

			if s2 >= 21 {
				p2Wins += (u2 * player1.universe)
				continue
			}

			updateP2 := player{p2, s2, u2}

			w1, w2 := roll(rollResult, universe, updateP1, updateP2)
			p1Wins += w1
			p2Wins += w2
		}
	}

	return p1Wins, p2Wins
}
