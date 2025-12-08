package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type ingredientRange struct {
	b int // beginning
	e int // end
}

func main() {
	part2("day5.txt")
}

func part1(file string) {
	ingRanges, ingredients := readInput(file)
	var freshCount int
	for _, ingredient := range ingredients {
		for _, ingRange := range ingRanges {
			if ingredient >= ingRange.b && ingredient <= ingRange.e {
				freshCount++
				break
			}
		}
	}
	fmt.Println(freshCount)
}

func part2(file string) {
	ingRanges, _ := readInput(file)
	sort.Slice(ingRanges, func(i, j int) bool { return ingRanges[i].b < ingRanges[j].b })
	fmt.Println(ingRanges)
	compactedRanges := compactSortedRanges(ingRanges)
	fmt.Println(compactedRanges)
	freshIngredients := 0
	for _, ingRange := range compactedRanges {
		freshIngredients += ingRange.e - ingRange.b + 1
	}
	fmt.Println(freshIngredients)
}

func compactSortedRanges(ranges []ingredientRange) []ingredientRange {
	foundAny := false
	var res []ingredientRange
	for i := 0; i < len(ranges); i++ {
		if i < len(ranges)-1 && ranges[i].e >= ranges[i+1].b {
			res = append(res, ingredientRange{ranges[i].b, int(math.Max(float64(ranges[i+1].e), float64(ranges[i].e)))})
			i++
			foundAny = true
		} else if i > 0 && ranges[i-1].e == ranges[i].e {
			continue
		} else {
			res = append(res, ranges[i])
		}
	}
	if foundAny {
		return compactSortedRanges(res)
	} else {
		return res
	}
}

func readInput(file string) ([]ingredientRange, []int) {
	bytes, e := os.ReadFile(file)
	if e != nil {
		fmt.Printf("Error reading file: %e\\n", e)
	}

	input := string(bytes)

	var ingRanges []ingredientRange
	var ingredients []int

	inputIterator := strings.SplitSeq(input, "\n")

	for s := range inputIterator {
		if s == "" {
			continue
		} else if strings.Contains(s, "-") {
			split := strings.Split(s, "-")
			beginning, _ := strconv.Atoi(split[0])
			end, _ := strconv.Atoi(split[1])
			ir := ingredientRange{int(beginning), int(end)}
			ingRanges = append(ingRanges, ir)
		} else {
			i, _ := strconv.Atoi(s)
			ingredients = append(ingredients, int(i))
		}
	}

	return ingRanges, ingredients

}
