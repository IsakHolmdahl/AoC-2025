package main

import (
	"cmp"
	"fmt"
	"maps"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type box struct {
	x float64
	y float64
	z float64
}

type boxDistance struct {
	distance float64
	boxOne   box
	boxTwo   box
}

func main() {
	part2("day8.txt")
}

func part1(file string) {
	boxes := readInput(file)
	availableConnections := findDistances(boxes)
	slices.SortFunc(availableConnections, func(distanceOne, distanceTwo boxDistance) int {
		return cmp.Compare(distanceOne.distance, distanceTwo.distance)
	})
	shortestDistances := availableConnections[:1000]
	connections := make([]map[box]bool, 0)
	connections = append(connections, map[box]bool{shortestDistances[0].boxOne: true, shortestDistances[0].boxTwo: true})
	for _, distance := range shortestDistances[1:] {
		existingConnectionFound := -1
		for i, connection := range connections {
			if connection[distance.boxOne] && !connection[distance.boxTwo] {
				if existingConnectionFound > -1 {
					maps.Copy(connections[existingConnectionFound], connections[i])
					connections[i] = make(map[box]bool)
				} else {
					connections[i][distance.boxTwo] = true
					existingConnectionFound = i
				}
			} else if connection[distance.boxTwo] && !connection[distance.boxOne] {
				if existingConnectionFound > -1 {
					maps.Copy(connections[existingConnectionFound], connections[i])
					connections[i] = make(map[box]bool)
				} else {
					connections[i][distance.boxOne] = true
					existingConnectionFound = i
				}
			} else if connection[distance.boxTwo] && connection[distance.boxOne] {
				existingConnectionFound = -2
			}
		}
		if existingConnectionFound == -1 {
			connections = append(connections, map[box]bool{distance.boxOne: true, distance.boxTwo: true})
		}

	}
	slices.SortFunc(connections, func(mapOne, mapTwo map[box]bool) int {
		return cmp.Compare(len(mapTwo), len(mapOne))
	})
	fmt.Println(len(connections[0]) * len(connections[1]) * len(connections[2]))
}

func part2(file string) {
	boxes := readInput(file)
	availableConnections := findDistances(boxes)
	slices.SortFunc(availableConnections, func(distanceOne, distanceTwo boxDistance) int {
		return cmp.Compare(distanceOne.distance, distanceTwo.distance)
	})
	connections := make([]map[box]bool, 0)
	connections = append(connections, map[box]bool{availableConnections[0].boxOne: true, availableConnections[0].boxTwo: true})
	for _, distance := range availableConnections[1:] {
		existingConnectionFound := -1
		for i, connection := range connections {
			if connection[distance.boxOne] && !connection[distance.boxTwo] {
				if existingConnectionFound > -1 {
					maps.Copy(connections[existingConnectionFound], connections[i])
					connections = append(connections[:i], connections[i+1:]...)
				} else {
					connections[i][distance.boxTwo] = true
					existingConnectionFound = i
				}
			} else if connection[distance.boxTwo] && !connection[distance.boxOne] {
				if existingConnectionFound > -1 {
					maps.Copy(connections[existingConnectionFound], connections[i])
					connections = append(connections[:i], connections[i+1:]...)

				} else {
					connections[i][distance.boxOne] = true
					existingConnectionFound = i
				}
			} else if connection[distance.boxTwo] && connection[distance.boxOne] {
				existingConnectionFound = -2
			}
		}
		if existingConnectionFound == -1 {
			connections = append(connections, map[box]bool{distance.boxOne: true, distance.boxTwo: true})
		} else if len(connections[0]) == len(boxes) {
			fmt.Println(distance.boxOne.x * distance.boxTwo.x)
			break
		}
	}
}

func readInput(file string) []box {
	fileBytes, _ := os.ReadFile(file)
	lines := strings.Split(string(fileBytes), "\n")
	positions := make([]box, 0)
	for _, row := range lines {
		if row == "" {
			continue
		}
		values := strings.Split(row, ",")
		xPos, _ := strconv.ParseFloat(values[0], 64)
		yPos, _ := strconv.ParseFloat(values[1], 64)
		zPos, _ := strconv.ParseFloat(values[2], 64)
		positions = append(positions, box{xPos, yPos, zPos})
	}
	return positions
}

func findDistances(boxes []box) []boxDistance {
	distances := make([]boxDistance, 0)
	for i, box := range boxes[:len(boxes)-1] {
		for _, otherBox := range boxes[i+1:] {
			distances = append(distances, boxDistance{straightLineDistance(box, otherBox), box, otherBox})
		}
	}
	return distances
}

func straightLineDistance(boxOne box, boxTwo box) float64 {
	x := math.Pow((boxOne.x - boxTwo.x), 2.0)
	y := math.Pow((boxOne.y - boxTwo.y), 2.0)
	z := math.Pow((boxOne.z - boxTwo.z), 2.0)
	return math.Sqrt(x + y + z)
}
