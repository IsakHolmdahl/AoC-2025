package main

import (
	"fmt"
	"os"
	"slices"
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

}

func dfsSearch(nodesGraph []node, currentNode node, visitedNodes []string, foundExits int) int {
	res := 0
	for i, n := range currentNode.outputs {
		if slices.Contains(visitedNodes, n) {
			continue
		} else if n == "out" {
			res++
		} else {
			sliceCopy := make([]string, 0)
			copy(sliceCopy, visitedNodes)
			sliceCopy = append(sliceCopy, n)
			destinationNode := nodesGraph[slices.IndexFunc(nodesGraph, func(v node) bool {
				return v.name == n
			})]
			res += dfsSearch([]node)
		}
	}
}
