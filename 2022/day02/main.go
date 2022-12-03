package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	score := 0
	points := secondPart()
	for scanner.Scan() {
		s := scanner.Text()
		score += points[s]
	}
	fmt.Println(score)
}

func firstPart() map[string]int {
	// 1 rock, 2 paper, 3 scissors
	// 0 lose, 3 draw, 6 win
	p := map[string]int{
		"A X": 4, "A Y": 8, "A Z": 3,
		"B X": 4, "B Y": 8, "B Z": 3,
		"C X": 4, "C Y": 8, "C Z": 3,
	}
	return p
}

func secondPart() map[string]int {
	// 1 rock, 2 paper, 3 scissors
	// 0 lose, 3 draw, 6 win
	p := map[string]int{
		"A X": 3, "A Y": 4, "A Z": 8,
		"B X": 1, "B Y": 5, "B Z": 9,
		"C X": 2, "C Y": 6, "C Z": 7,
	}
	return p
}