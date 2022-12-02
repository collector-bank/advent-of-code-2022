package main

import (
	"strings"
)

func Day02part01(input string) int {
	rows := strings.Split(input, "\n")

	score := 0
	for _, row := range rows {
		temp := strings.Split(row, " ")
		bot := int(temp[0][0] - 'A')
		you := int(temp[1][0] - 'X')

		score += you + 1
		if (bot == 0 && you == 1) || (bot == 1 && you == 2) || (bot == 2 && you == 0) {
			score += 6
		} else if bot == you {
			score += 3
		}
	}

	return score
}

func Day02part02(input string) int {
	rows := strings.Split(input, "\n")

	score := 0
	for _, row := range rows {
		temp := strings.Split(row, " ")
		bot := int(temp[0][0] - 'A')
		win := int(temp[1][0] - 'X')

		if win == 0 {
			// lose
			you := bot - 1
			if you < 0 {
				you = 2
			}
			score += you + 1
		} else if win == 1 {
			// draw
			you := bot
			score += you + 1
			score += 3
		} else {
			// win
			you := (bot + 1) % 3
			score += you + 1
			score += 6
		}
	}

	return score
}
