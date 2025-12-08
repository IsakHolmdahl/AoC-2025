package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	day2()
}

func day1() {
	lines, _ := getLines()
	zeroHits := 0
	currentValue := 50
	for _, element := range lines {
		dir := element[0]
		steps, _ := strconv.Atoi(element[1:])
		if dir == 'L' {
			currentValue -= steps
		} else {
			currentValue += steps
		}
		currentValue = currentValue % 100
		if currentValue == 0 {
			zeroHits += 1
		}
	}
	println(zeroHits)
}

func day2() {
	lines, _ := getLines()
	zeroHits := 0
	currentValue := 50
	for _, element := range lines {
		dir := string(element[0])
		steps, _ := strconv.Atoi(element[1:])
		if steps >= 100 {
			f := float64(steps)
			zeroHits += int(math.Floor(f / 100))
			steps = steps % 100
			if steps == 0 {
				continue
			}
		}
		if dir == "L" {
			if steps >= currentValue && currentValue != 0 {
				zeroHits++
			}
			currentValue -= steps
		} else {
			if steps >= (100 - currentValue) {
				zeroHits++
			}
			currentValue += steps
		}
		currentValue = nonNegativeMod(currentValue, 100)
	}
	println(zeroHits)
}

func nonNegativeMod(value int, mod int) int {
	res := value % mod
	if res < 0 {
		res = mod + res
	}
	return res
}

func getLines() ([]string, error) {
	data, err := os.ReadFile("day1.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	return lines[:len(lines)-1], nil

}
