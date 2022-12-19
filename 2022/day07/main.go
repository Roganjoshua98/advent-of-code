package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Directory struct {
	Name string;
	Subdirectories []*Directory;
	Files []*File;
	Size int;
	Previous *Directory;
}

type File struct {
	Name string;
	Size int;
}

func main() {
	file, _ := os.ReadFile("input.txt")
	terminal := strings.Split(string(file), "\n")
	homeDir := Directory{Name: "/", Size: 0}
	currDir := &homeDir
	for i := 1; i < len(terminal); i++ {
		command := strings.Split(terminal[i], " ")
		switch command[1] {
		case "cd":
			currDir = cd(currDir, command[2])
		case "ls":
			var output [][]string
			output, i = parseOutput(i, terminal)
			currDir = ls(currDir, output)
		}
	}
	// fmt.Println(atMostTenThousand(homeDir))
	total, required := 70000000, 30000000
	leftover := total-homeDir.Size
	// total 70, size 50, leftover 20, required 30
	// target 30 - 20 = 10
	if leftover < required {
		fmt.Println(rm(homeDir, required-leftover, homeDir.Size))
	}
}

func cd(currDir *Directory, target string) *Directory {
	if target == ".." {
		return currDir.Previous
	} else {
		for _, d := range currDir.Subdirectories {
			if d.Name == target {
				return d
			}
		}
	}
	return currDir
}

func ls(currDir *Directory, output [][]string) *Directory {
	for _, line := range output {
		name := line[1]
		if line[0] == "dir" {
			dir := Directory{Name: name, Previous: currDir}
			currDir.Subdirectories = append(currDir.Subdirectories, &dir)
		} else {
			size, _ := strconv.Atoi(line[0])
			file := File{Name: name, Size: size}
			currDir.Files = append(currDir.Files, &file)
			currDir.Size += file.Size
			d := currDir.Previous
			for d != nil {
				d.Size += file.Size
				d = d.Previous
			}
		}
	}
	return currDir
}

func rm(dir Directory, target int, result int) int {
	for _, d := range dir.Subdirectories {
		result = rm(*d, target, result)
		if d.Size > target && d.Size < result {
			result = d.Size
		}
	}
	return result
}

func parseOutput(i int, terminal []string) ([][]string, int) {
	i += 1
	output := make([][]string, 0)
	for i < len(terminal) && terminal[i][0] != '$'{
		output = append(output, strings.Split(terminal[i], " "))
		i += 1
	}
	i -= 1
	return output, i
}

func atMostTenThousand(dir Directory) (total int) {
	for _, d := range dir.Subdirectories {
		total += atMostTenThousand(*d)
		if d.Size <= 100000 {
			total += d.Size
		}
	}
	return total
}