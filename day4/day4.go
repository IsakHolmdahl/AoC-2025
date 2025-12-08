package main

import (
	"fmt"
	"os"
	"strings"
)

type pos struct {
	x int
	y int
}

func main() {
	part2("day4.txt")
}

func part1(file string) {
	matrix := create2dSlice(file)
	accessibleRolls := make(map[pos]bool)
	for y, row := range matrix {
		for x, e := range row {
			if e != '@' {
				continue
			}
			if getPaperSum(matrix, y, x) < 4 {
				accessibleRolls[pos{x: x, y: y}] = true
			}
		}
	}
	fmt.Println(len(accessibleRolls))
}

func part2(file string) {
	matrix := create2dSlice(file)
	sum := 0
	for true {
		accessibleRolls := make(map[pos]bool)
		for y, row := range matrix {
			for x, e := range row {
				if e != '@' {
					continue
				}
				if getPaperSum(matrix, y, x) < 4 {
					accessibleRolls[pos{x: x, y: y}] = true
				}
			}
		}
		sum += len(accessibleRolls)
		if len(accessibleRolls) == 0 {
			break
		}
		movePaper(matrix, accessibleRolls)
	}
	fmt.Println(sum)
}

func movePaper(matrix [][]rune, accessibleRolls map[pos]bool) {
	for k := range accessibleRolls {
		matrix[k.y][k.x] = '.'
	}
}

func create2dSlice(file string) [][]rune {
	raw, e := os.ReadFile(file)
	if e != nil {
		fmt.Println("Error reading file: ", e)
	}
	rows := strings.Split(string(raw), "\n")
	res := make([][]rune, len(rows)-1)
	for i, r := range rows {
		if strings.Trim(r, "\n") != "" {
			res[i] = []rune(r)
		}
	}
	return res
}

func getPaperSum(matrix [][]rune, y int, x int) int {
	paperSum := 0
	if y != 0 {
		if matrix[y-1][x] == '@' {
			paperSum++
		}
		if x != 0 && matrix[y-1][x-1] == '@' {
			paperSum++
		}
		if x != len(matrix[y])-1 && matrix[y-1][x+1] == '@' {
			paperSum++
		}
	}
	if x != 0 && matrix[y][x-1] == '@' {
		paperSum++
	}

	if y != len(matrix)-1 {
		if matrix[y+1][x] == '@' {
			paperSum++
		}
		if x != len(matrix[y])-1 && matrix[y+1][x+1] == '@' {
			paperSum++
		}
		if x != 0 && matrix[y+1][x-1] == '@' {
			paperSum++
		}
	}
	if x != len(matrix[y])-1 && matrix[y][x+1] == '@' {
		paperSum++
	}
	return paperSum
}
