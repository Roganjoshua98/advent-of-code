package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	fmt.Println(partOne(file))
	fmt.Println(partTwo(file))
}

func partOne(file *os.File) int {
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		pairs := lineToPairs(scanner.Text())
		if (pairs[1][0] >= pairs[0][0] && pairs[1][1] <= pairs[0][1]) || 
			(pairs[0][0] >= pairs[1][0] && pairs[0][1] <= pairs[1][1]) {
				total += 1
			}
	}
	return total
}

func partTwo(file *os.File) int {
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		pairs := lineToPairs(scanner.Text())
		// Easier to prove if they *don't* overlap
		min, max := pairs[0], pairs[1]
		if pairs[1][0] < pairs[0][1] {
			min, max = pairs[1], pairs[0]
		}
		if !(min[1] < max[0]) {
			total += 1
		}
	}
	return total
}

func lineToPairs(s string) [][]int {
	pairs := strings.Split(s, ",")
	result := make([][]int, 0)
	for _, pair := range pairs {
		p := strings.Split(pair, "-")
		pa, _ := strconv.Atoi(p[0])
		pb, _ := strconv.Atoi(p[1])
		result = append(result, []int{pa, pb})
	}
	return result
}