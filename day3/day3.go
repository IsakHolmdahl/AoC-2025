package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(part2("day3.txt"))
}

func part1(file string) int {
	batteryRows := getInputRowAsSlice(file)
	var sum int = 0
	for _, row := range batteryRows {
		var highestFromLeft, highestFromRight, highestFromLeftIndex uint8 = 0, 0, 0
		for y, e := range row[:len(row)-1] {
			if e > highestFromLeft {
				highestFromLeft = e
				highestFromLeftIndex = uint8(y)
			}
		}
		for y := uint8(len(row)) - 1; y > highestFromLeftIndex; y-- {
			if row[y] > highestFromRight {
				highestFromRight = row[y]
			}
		}
		sum += int((highestFromLeft * 10) + highestFromRight)
	}
	return sum
}

func part2(file string) int {
	batteryRows := getInputRowAsSlice(file)
	sum := 0
	for _, row := range batteryRows {
		builtNumber := 0
		indexForHighest := 0
		for y := range 12 {
			highest := 0
			for x := indexForHighest; x < len(row)-11+y; x++ {
				if int(row[x]) > highest {
					highest = int(row[x])
					indexForHighest = x + 1
				}
			}
			builtNumber += highest * int(math.Pow10(11-y))
		}
		fmt.Println(builtNumber)
		sum += builtNumber
	}
	return sum
}

func getInputRowAsSlice(file string) [][]uint8 {
	input, e := os.ReadFile(file)
	if e != nil {
		fmt.Println("Error reading file", e)
		return nil
	}
	splits := strings.Split(string(input), "\n")
	for i, e := range splits {
		splits[i] = strings.Trim(e, "\n")
	}
	res := make([][]uint8, len(splits)-1)
	for i, _ := range res {
		res[i] = make([]uint8, len(splits[i]))
		for y, f := range splits[i] {
			runeToInt, _ := strconv.Atoi(string(f))
			res[i][y] = uint8(runeToInt)
		}
	}
	return res
}
