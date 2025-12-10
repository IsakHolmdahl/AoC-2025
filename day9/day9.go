package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	part2("day9.txt")
}

type pos struct {
	x int
	y int
}

func part1(file string) {
	input, _ := os.ReadFile(file)
	lines := strings.Split(string(input), "\n")
	topLeft, topRight, bottomLeft, bottomRight := pos{-1, -1}, pos{-1, -1}, pos{-1, -1}, pos{-1, -1}
	var distanceTopLeft, distanceTopRight, distanceBottomLeft, distanceBottomRight float64
	points := make([]pos, 0)
	gridX, gridY := -1, -1

	for _, row := range lines {
		if row == "" {
			continue
		}
		values := strings.Split(row, ",")
		x, _ := strconv.Atoi(values[0])
		y, _ := strconv.Atoi(values[1])
		if x > gridX {
			gridX = x
		}
		if y > gridY {
			gridY = y
		}
		coord := pos{x, y}
		points = append(points, coord)
	}

	for _, point := range points {
		if topLeft.x == -1 {
			topLeft = point
			distanceTopLeft = calcDistanceNoSqrt(point, pos{0, 0})
			topRight = point
			distanceTopRight = calcDistanceNoSqrt(point, pos{gridX - 1, 0})
			bottomLeft = point
			distanceBottomLeft = calcDistanceNoSqrt(point, pos{0, gridY - 1})
			bottomRight = point
			distanceBottomRight = calcDistanceNoSqrt(point, pos{gridX - 1, gridY - 1})
			fmt.Println(topLeft)
			continue
		}
		xyDistanceTopLeft := calcDistanceNoSqrt(point, pos{0, 0})
		xyDistanceTopRight := calcDistanceNoSqrt(point, pos{gridX - 1, 0})
		xyDistanceBottomLeft := calcDistanceNoSqrt(point, pos{0, gridY - 1})
		xyDistanceBottomRight := calcDistanceNoSqrt(point, pos{gridX - 1, gridY - 1})
		if xyDistanceTopLeft < distanceTopLeft {
			distanceTopLeft = xyDistanceTopLeft
			topLeft = point
		}
		if xyDistanceTopRight < distanceTopRight {
			distanceTopRight = xyDistanceTopRight
			topRight = point
		}
		if xyDistanceBottomLeft < distanceBottomLeft {
			distanceBottomLeft = xyDistanceBottomLeft
			bottomLeft = point
		}
		if xyDistanceBottomRight < distanceBottomRight {
			distanceBottomRight = xyDistanceBottomRight
			bottomRight = point
		}
	}

	areaTLBR := (bottomRight.x - topLeft.x + 1) * (bottomRight.y - topLeft.y + 1)
	areaBLTR := (topRight.x - bottomLeft.x + 1) * (bottomLeft.y - topRight.y + 1)
	areaTLTR := (topRight.x - topLeft.x + 1) * (intAbs(topRight.y-topLeft.y) + 1)
	areaBLBR := (bottomRight.x - bottomLeft.x + 1) * (intAbs(bottomLeft.y-bottomRight.y) + 1)

	areas := []int{areaTLBR, areaBLTR, areaTLTR, areaBLBR}
	slices.Sort(areas)
	fmt.Print(areas[3])
}

func calcDistanceNoSqrt(start, end pos) float64 {
	return math.Pow(float64(end.x-start.x), 2) + math.Pow(float64(end.y-start.y), 2)
}

func intAbs(i int) int {
	if i < 0 {
		return i * -1
	} else {
		return i
	}
}

func intMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func intMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

type linkedPos struct {
	this   pos
	linked pos
}

// Go trough each point
// When past point 2, check area with the point two steps back.
// Check that no other point is in that area

// func part2Redo(file string) {
// 	input, _ := os.ReadFile(file)
// 	lines := strings.Split(string(input), "\n")
// 	points := make([][]pos, 0)
// 	for i, row := range lines {
// 		values := strings.Split(row, ",")
// 		x, _ := strconv.Atoi(values[0])
// 		y, _ := strconv.Atoi(values[1])
// 		coord := pos{x, y}
// 		points = append(points, linkedPos{coord, lastVisitedPoint})
// 		lastVisitedPoint = coord
// 	}
// }

func part2(file string) {
	input, _ := os.ReadFile(file)
	lines := strings.Split(string(input), "\n")
	valuesFirstPoint := strings.Split(lines[0], ",")
	valuesLastPoint := strings.Split(lines[len(lines)-2], ",")
	xFirstPoint, _ := strconv.Atoi(valuesFirstPoint[0])
	yFirstPoint, _ := strconv.Atoi(valuesFirstPoint[1])
	xLastPoint, _ := strconv.Atoi(valuesLastPoint[0])
	yLastPoint, _ := strconv.Atoi(valuesLastPoint[1])
	firstPoint, lastPoint := linkedPos{pos{xFirstPoint, yFirstPoint}, pos{-1, -1}}, linkedPos{pos{xLastPoint, yLastPoint}, pos{-1, -1}}
	gridX, gridY := intMax(firstPoint.this.x, lastPoint.this.x), intMax(firstPoint.this.y, lastPoint.this.y)
	firstPoint.linked = lastPoint.this
	points := []linkedPos{firstPoint, lastPoint}
	lastVisitedPoint := firstPoint.this
	for _, row := range lines[1 : len(lines)-2] {
		values := strings.Split(row, ",")
		x, _ := strconv.Atoi(values[0])
		y, _ := strconv.Atoi(values[1])
		if x > gridX {
			gridX = x
		}
		if y > gridY {
			gridY = y
		}
		coord := pos{x, y}
		points = append(points, linkedPos{coord, lastVisitedPoint})
		lastVisitedPoint = coord
	}
	points[1].linked = lastVisitedPoint

	borderMap := make([][]string, gridY+1)

	for i := range gridY + 1 {
		borderMap[i] = make([]string, gridX+1)
	}

	fmt.Println(gridX)
	fmt.Println(gridY)
	for y, row := range borderMap {
		for x := range row {
			for _, point := range points {
				if point.this.x == x && point.this.y == y {
					if point.this.x == point.linked.x {
						var goingDown bool
						if point.this.y < point.linked.y {
							goingDown = true
						}
						for i := range intAbs(point.this.y - point.linked.y) {
							if goingDown {
								borderMap[y+i][x] = "|"
							} else {
								borderMap[y-i][x] = "|"
							}
						}
					} else {
						var goingRight bool
						if point.this.x < point.linked.x {
							goingRight = true
						}
						for i := range intAbs(point.this.x - point.linked.x) {
							if goingRight {
								borderMap[y][x+i] = "-"
							} else {
								borderMap[y][x-i] = "-"
							}
						}
					}
					borderMap[y][x] = "@"
				}
			}
			if borderMap[y][x] == "" {
				borderMap[y][x] = "."
			}
		}
	}

	for _, row := range borderMap {
		fmt.Println(row)
	}

	// biggestArea := 0

	// for i, point := range points[:len(points)-1] {
	// 	for _, oppositePoint := range points[i+1:] {
	// 		area := intAbs(point.this.x-oppositePoint.this.x) * intAbs(point.this.y-oppositePoint.this.y)
	// 		if area > biggestArea {
	// 			inside := rayTrace(borderMap, point, oppositePoint)
	// 			if inside {
	// 				biggestArea = area
	// 			}
	// 		}
	// 	}
	// }
	// fmt.Println(biggestArea)
}

func rayTrace(borderMap [][]string, cornerOne, cornerTwo linkedPos) bool {
	areaCoords := getAreaPoints(cornerOne.this, cornerTwo.this)
	for _, row := range areaCoords {
		for _, coord := range row {
			var isInside bool
			if borderMap[coord.y][coord.x] == "|" || borderMap[coord.y][coord.x] == "-" || (borderMap[coord.y][coord.x] == "@" && borderMap[coord.y][coord.x+1] == "-") {
				isInside = true
			} else {
				isInside = false
			}
			for x := range len(borderMap[1]) - coord.x - 1 {
				if borderMap[coord.y][coord.x+x+1] == "|" || borderMap[coord.y][coord.x+x+1] == "@" {
					isInside = !isInside
				}
			}
			if isInside == false {
				// fmt.Println(areaCoords)
				return false
			}
		}
	}
	return true
}

func getAreaPoints(cornerOne, cornerTwo pos) [][]pos {
	upperLeftCorner := pos{intMin(cornerOne.x, cornerTwo.x), intMin(cornerOne.y, cornerTwo.y)}
	bottomRightCorner := pos{intMax(cornerOne.x, cornerTwo.x), intMax(cornerOne.y, cornerTwo.y)}
	res := make([][]pos, intAbs(upperLeftCorner.y-bottomRightCorner.y)+1)

	for i := range res {
		res[i] = make([]pos, intAbs(upperLeftCorner.x-bottomRightCorner.x)+1)
		for j := range res[i] {
			res[i][j] = pos{upperLeftCorner.x + j, upperLeftCorner.y + i}
		}
	}
	return res
}
