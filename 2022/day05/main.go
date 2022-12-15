package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var r, _ = regexp.Compile(`\d+`)

func main() {
	buf, _ := os.ReadFile("input.txt")
	fs := strings.Split(string(buf), "\n\n")
	first, second := strings.Split(fs[0], "\n"), strings.Split(fs[1], "\n")
	stacks := make([][]rune, 0)
	for _, line := range first {
		stacks = append(stacks, []rune(line))
	}
	for _, line := range second {
		cmds := parseCommand(line)
		fmt.Println(cmds)
		fmt.Println(stacks)
		stacks = executeCommand(stacks, cmds)
	}
	for _, line := range stacks {
		fmt.Print(string(line[len(line)-1]))
	}
	fmt.Println()
}

func parseCommand(line string) []int {
	cmds := make([]int, 0)
	for _, s := range r.FindAllString(line, -1) {
		n, _ := strconv.Atoi(s)
		cmds = append(cmds, n)
	}
	return cmds
}

func executeCommand(stacks [][]rune, cmds []int) [][]rune {
	move, from, to := cmds[0], cmds[1]-1, cmds[2]-1
	carry := make([]rune, 0)
	for i := 0; i < move; i++ {
		elem := stacks[from][len(stacks[from])-1]
		carry = append([]rune{elem}, carry...)
		stacks[from] = stacks[from][:len(stacks[from])-1]
	}
	for _, c := range carry {
		stacks[to] = append(stacks[to], c)
	}
	return stacks
}