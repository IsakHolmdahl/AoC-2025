package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	part1("day11example.txt")
}

type node struct {
	name    string
	outputs []string
}

func part1(file string) {
	input, _ := os.ReadFile(file)
	lines := strings.Split(string(input), "\n")
	nodes := make([]node, len(lines)-1)
	for i, s := range lines[:len(lines)-1] {
		parts := strings.Split(s, " ")
		nodes[i] = node{strings.Trim(parts[0], ":"), parts[1:]}
	}
	fmt.Println(nodes)
}
