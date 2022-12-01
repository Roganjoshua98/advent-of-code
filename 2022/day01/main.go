package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	elves := make([][]int, 0)
	elf := make([]int, 0)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			elves = append(elves, elf)
			elf = make([]int, 0)
		} else {
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			elf = append(elf, i)
		}
	}
	fmt.Println(mostCalories(elves))
}

func mostCalories(elves [][]int) int {
	highest := 0
	for _, elf := range elves {
		sum := 0
		for _, cal := range elf {
			sum += cal
		}
		if sum > highest {
			highest = sum
		}
	}
	return highest
}