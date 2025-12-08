package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	part2("day7.txt")
}

func part1(file string) {
	input, _ := os.ReadFile(file)
	lines := strings.Split(string(input), "\n")
	tachyonMap, splits := make(map[int]bool), 0
	for _, row := range lines {
		for x, c := range row {
			if c == '^' {
				if tachyonMap[x] == true {
					splits++
					tachyonMap[x] = false
					if x > 0 {
						tachyonMap[x-1] = true
					}
					if x < len(row)-1 {
						tachyonMap[x+1] = true
					}
				}
			} else if c == 'S' {
				tachyonMap[x] = true
			}
		}
	}
	fmt.Println(splits)
}

func part2(file string) {
	input, _ := os.ReadFile(file)
	lines := strings.Split(string(input), "\n")
	tachyonMap, timelines := make(map[int]int), 1
	tachyonMap[len(lines[0])/2] = 1
	for _, row := range lines {
		var splits int
		for x, c := range row {
			if c == '^' && tachyonMap[x] > 0 {
				splits += tachyonMap[x]
				tachyonMap[x+1] += tachyonMap[x]
				tachyonMap[x-1] += tachyonMap[x]
				tachyonMap[x] -= tachyonMap[x]
				// for range tachyonMap[x] {
				// 	splits++
				// 	tachyonMap[x]--
				// 	tachyonMap[x+1]++
				// 	tachyonMap[x-1]++
				// }
			}
		}
		if splits != 0 {
			timelines += splits
		}
	}
	fmt.Println(timelines)
}

func recursiveBruteForce(inputRows []string, timelines int, particleX int) int {
	for i, row := range inputRows {
		if row[particleX] == '^' {
			timelines += 1
			timelines += recursiveBruteForce(inputRows[i+1:], 0, particleX+1)
			particleX -= 1
		}
	}
	return timelines
}
