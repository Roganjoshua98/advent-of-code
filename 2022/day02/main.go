package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	score := 0
	points := secondPart()
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		score += points[s[0]][s[1]]
	}
	fmt.Println(score)
}

func firstPart() map[string]map[string]int {
	// 1 rock, 2 paper, 3 scissors
	// 0 lose, 3 draw, 6 win
	points := map[string]map[string]int{
        "A": map[string]int{
            "X": 4,
            "Y": 8,
            "Z": 3,

        },
        "B": map[string]int{
            "X": 1,
            "Y": 5,
            "Z": 9,

        },
		"C": map[string]int{
            "X": 7,
            "Y": 2,
            "Z": 6,

        },
    }
	return points
}

func secondPart() map[string]map[string]int {
	// 1 rock, 2 paper, 3 scissors
	// 0 lose, 3 draw, 6 win
	points := map[string]map[string]int{
        "A": map[string]int{
            "X": 3,
            "Y": 4,
            "Z": 8,

        },
        "B": map[string]int{
            "X": 1,
            "Y": 5,
            "Z": 9,

        },
		"C": map[string]int{
            "X": 2,
            "Y": 6,
            "Z": 7,

        },
    }
	return points
}