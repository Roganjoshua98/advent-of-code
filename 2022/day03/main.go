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
	x := partTwo(scanner)
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

func partTwo(scanner *bufio.Scanner) int {
	score := 0
	group := make([]map[rune]int, 0)
	i := 0
	for scanner.Scan() {
		s := scanner.Text()
		if i < 2 {
			m := make(map[rune]int)
			for _, c := range s {
				m[c] += 1
			}
			group = append(group, m)
			i += 1
		} else {
			for _, c := range s {
				_, ok1 := group[0][c]
				_, ok2 := group[1][c]
				if ok1 && ok2 {
					score += charPriority(c)
					break
				}
			}
			i = 0
			group = make([]map[rune]int, 0)
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