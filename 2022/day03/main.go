package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	x := partOne(scanner)
	fmt.Println(x)
}

func partOne(scanner *bufio.Scanner) int {
	score := 0
	for scanner.Scan() {
		s := scanner.Text()
		first, second := s[0:(len(s) / 2)], s[(len(s) / 2):]
		m := make(map[rune]int)
		for _, c := range first {
			m[c] += 1
		}
		for _, c := range second {
			if _, ok := m[c]; ok {
				score += charPriority(c)
				break
			}
		}
	}
	return score
}

func charPriority(c rune) int {
	if unicode.IsUpper(c) {
		return int(c - 38)
	} else {
		return int(c - 96)
	}
}