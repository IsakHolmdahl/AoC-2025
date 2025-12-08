package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	day2()
}

func day1() {
	ranges := readInput()
	sum := 0
	for _, element := range ranges {
		sum += getSumFromRange(element)
	}
	fmt.Println(sum)
}

func day2() {
	ranges := readInput()
	sum := big.NewInt(0)
	for _, element := range ranges {
		res := getSumFromRange2(element)
		sum = new(big.Int).Add(res, sum)
	}
	fmt.Println(sum)
}

func readInput() []string {
	data, err := os.ReadFile("day2.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	ranges := strings.Split(string(data), ",")
	return ranges
}

func getSumFromRange(r string) int {
	sum := 0
	split := strings.Split(r, "-")
	beginning, _ := strconv.Atoi(split[0])
	end, _ := strconv.Atoi(split[1])

	for i := beginning; i <= end; i++ {
		toString := strconv.Itoa(i)
		if len(toString)%2 != 0 {
			continue
		}
		firstHalf, secondHalf := toString[:len(toString)/2], toString[len(toString)/2:]
		if firstHalf == secondHalf {
			sum += i
		}
	}
	return sum
}

func getSumFromRange2(r string) *big.Int {
	sum := big.NewInt(0)
	split := strings.Split(strings.Trim(r, "\n"), "-")
	beginning, _ := strconv.Atoi(split[0])
	end, _ := strconv.Atoi(split[1])
	for i := beginning; i <= end; i++ {
		toString := strconv.Itoa(i)
		if recursiveCheck(toString, len(toString)) {
			sum = new(big.Int).Add(big.NewInt(int64(i)), sum)
		}
	}
	return sum
}

func recursiveCheck(str string, splits int) bool {
	if splits == 1 {
		return false
	}
	if len(str)%splits != 0 {
		return recursiveCheck(str, splits-1)
	}
	substrings := splitString(str, splits)
	for _, element := range substrings[1:] {
		if element != substrings[0] {
			return recursiveCheck(str, splits-1)
		}
	}
	return true
}

func splitString(str string, splits int) []string {
	res := make([]string, splits)
	for i := range splits {
		res[i] = str[(len(str)/splits)*i : (len(str)/splits)*(i+1)]
	}
	return res
}
