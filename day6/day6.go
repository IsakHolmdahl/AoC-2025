package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part2("day6.txt")
}

func part1(file string) {
	problems := parseInput(file)
	fmt.Println(problems)
	var problemsSum int
	for i := range len(problems[0]) {
		var isMultiply = false
		var problemsAnswer, _ = strconv.Atoi(problems[0][i])
		if problems[len(problems)-1][i] == "*" {
			isMultiply = true
		}
		for _, row := range problems[1 : len(problems)-1] {
			value, _ := strconv.Atoi(row[i])
			if isMultiply {
				problemsAnswer *= value
			} else {
				problemsAnswer += value

			}
		}

		problemsSum += problemsAnswer
	}
	fmt.Println(problemsSum)
}

func parseInput(file string) [][]string {
	inputBytes, _ := os.ReadFile(file)
	lines := strings.Split(string(inputBytes), "\n")
	fmt.Println(lines)
	problems := make([][]string, len(lines)-1)
	for i, e := range lines {
		if e == "" || e == " " {
			continue
		}
		parts := strings.Split(e, " ")
		problems[i] = make([]string, 0)
		for _, f := range parts {
			if f != "" && f != " " {
				problems[i] = append(problems[i], strings.Trim(f, " "))
			}
		}
	}
	return problems
}

func part2(file string) {
	values, operators := parseInputVertical(file)
	sum := 0
	for i, problemValues := range values {
		problemResult := problemValues[0]
		for _, value := range problemValues[1:] {
			if strings.Trim(operators[i], " ") == "+" {
				problemResult += value
			} else {
				problemResult *= value
			}
		}
		fmt.Println(problemResult)
		sum += problemResult
	}
	fmt.Println(sum)
}

func parseInputVertical(file string) ([][]int, []string) {
	inputBytes, _ := os.ReadFile(file)
	textLines := strings.Split(string(inputBytes), "\n")
	var values, operators, problemIndex = make([][]int, 1), make([]string, 0), 0
	values[0] = make([]int, 0)
	for x := range len(textLines[0]) {
		var foundDigit = false
		var value = 0
		for _, line := range textLines[:len(textLines)-2] {
			if line[x] != ' ' {
				digit, _ := strconv.Atoi(string(line[x]))
				foundDigit = true
				value = (value * 10) + digit
			}
		}
		if foundDigit {
			values[problemIndex] = append(values[problemIndex], value)
		} else {
			problemIndex++
			values = append(values, make([]int, 0))
		}
	}
	operatorsSlice := strings.Split(textLines[len(textLines)-2], " ")
	for _, c := range operatorsSlice {
		if c != " " && c != "" {
			operators = append(operators, c)
		}
	}
	return values, operators

}

