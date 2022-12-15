package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	fmt.Println(findPacket(file, 4))
}

func findPacket(file []byte, size int) int {
	packet := file[0:size]
	duplicate := make(map[byte]int)
	flag := false
	// Process first 14 characters
	for _, b := range packet {
		if _, ok := duplicate[b]; ok {
			flag = true
		}
		duplicate[b] += 1
	}
	if !flag {
		return size
	}
	// For each subsequent packet, just remove packet[0] and append new char
	for i := size+1; i < len(file); i++ {
		flag = false
		removed, new := packet[0], file[i]
		packet = append(packet[1:], new)
		duplicate[removed] -= 1
		if duplicate[removed] == 0 {
			delete(duplicate, removed)
		}
		duplicate[new] += 1
		for _, v := range duplicate {
			if v > 1 {
				flag = true
				break
			}
		}
		if !flag {
			return i + 1
		}
	}
	return -1
}