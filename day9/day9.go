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

func part2(file string) {
	input, _ := os.ReadFile(file)
	lines := strings.Split(string(input), "\n")
	valuesFirstPoint := strings.Split(lines[0], ",")
	xFirstPoint, _ := strconv.Atoi(valuesFirstPoint[0])
	yFirstPoint, _ := strconv.Atoi(valuesFirstPoint[1])
	firstPoint := linkedPos{pos{xFirstPoint, yFirstPoint}, pos{-1, -1}}
	points := []linkedPos{firstPoint}
	verticallyJoinedPoints := make([]linkedPos, 0)
	lastVisitedPoint := firstPoint.this
	for _, row := range lines[1:] {
		if row == "" {
			continue
		}
		values := strings.Split(row, ",")
		x, _ := strconv.Atoi(values[0])
		y, _ := strconv.Atoi(values[1])
		coord := pos{x, y}
		points = append(points, linkedPos{coord, lastVisitedPoint})
		if coord.x == lastVisitedPoint.x {
			verticallyJoinedPoints = append(verticallyJoinedPoints, linkedPos{coord, lastVisitedPoint})
		}
		lastVisitedPoint = coord
	}
	points[0].linked = lastVisitedPoint
	if points[0].this.x == points[0].linked.x {
		verticallyJoinedPoints = append(verticallyJoinedPoints, points[0])
	}

	slices.SortFunc(verticallyJoinedPoints, func(a, b linkedPos) int {
		if a.this.x < b.this.x {
			return -1
		} else {
			return 1
		}
	})

	maxArea := 0
	for i, point := range points[:len(points)-1] {
		for _, oppositePoint := range points[i+1:] {
			if point.this.x == oppositePoint.this.x || point.this.y == oppositePoint.this.y {
				continue
			}
			upperLeft := pos{intMin(point.this.x, oppositePoint.this.x), intMin(point.this.y, oppositePoint.this.y)}
			lowerRight := pos{intMax(point.this.x, oppositePoint.this.x), intMax(point.this.y, oppositePoint.this.y)}
			middle := pos{(lowerRight.x - upperLeft.x) / 2, (lowerRight.y - upperLeft.y) / 2}
			area := (intAbs(point.this.x-oppositePoint.this.x) + 1) * (intAbs(point.this.y-oppositePoint.this.y) + 1)
			nothingInside := true
			for _, p := range points {
				if p.this.x > upperLeft.x && p.this.y > upperLeft.y && p.this.x < lowerRight.x && p.this.y < lowerRight.y {
					nothingInside = false
					break
				}
			}
			if !nothingInside {
				continue
			}
			isInsideVerticalUpperLeft := false
			isInsideVerticalLowerLeft := false
			isInsideVerticalUpperRight := false
			isInsideVerticalLowerRight := false
			isInsideMiddle := false
			for _, vPair := range verticallyJoinedPoints {
				if vPair.this.x > upperLeft.x+1 {
					if intMin(vPair.this.y, vPair.linked.y) <= upperLeft.y+1 && intMax(vPair.this.y, vPair.linked.y) >= upperLeft.y+1 {
						isInsideVerticalUpperLeft = !isInsideVerticalUpperLeft
					}
					if intMin(vPair.this.y, vPair.linked.y) <= lowerRight.y-1 && intMax(vPair.this.y, vPair.linked.y) >= lowerRight.y-1 {
						isInsideVerticalLowerLeft = !isInsideVerticalLowerLeft
					}
				}
				if vPair.this.x > lowerRight.x-1 {
					if intMin(vPair.this.y, vPair.linked.y) <= upperLeft.y+1 && intMax(vPair.this.y, vPair.linked.y) >= upperLeft.y+1 {
						isInsideVerticalUpperRight = !isInsideVerticalUpperRight
					}
					if intMin(vPair.this.y, vPair.linked.y) <= lowerRight.y-1 && intMax(vPair.this.y, vPair.linked.y) >= lowerRight.y-1 {
						isInsideVerticalLowerRight = !isInsideVerticalLowerRight
					}
				}
				if vPair.this.x > middle.x {
					if intMin(vPair.this.y, vPair.linked.y) <= middle.y && intMax(vPair.this.y, vPair.linked.y) >= middle.y {
						isInsideMiddle = !isInsideMiddle
					}
				}

			}
			if isInsideVerticalUpperLeft && isInsideVerticalLowerLeft && isInsideVerticalUpperRight && isInsideVerticalLowerRight && isInsideMiddle && area > maxArea {
				fmt.Println(point)
				fmt.Println(oppositePoint)
				maxArea = area
			}
		}

	}

	fmt.Println(maxArea)
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
