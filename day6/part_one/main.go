package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x    float64
	y    float64
	area int
}

func manhattanDistance(x int, y int, b point) float64 {
	return math.Abs(float64(x)-float64(b.x)) + math.Abs(float64(y)-float64(b.y))
}

func isWinner(i int, s []int) bool {
	for _, c := range s {
		if i == c {
			return false
		}
	}
	return true
}

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	points := make([]point, 0)
	var disqualified []int
	for scanner.Scan() {
		row := string(scanner.Text())
		coords := strings.Split(row, ", ")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		p := point{x: float64(x), y: float64(y)}
		points = append(points, p)
	}
	maxX := 0
	maxY := 0
	minX := 0
	minY := 0

	for _, p := range points {
		if int(p.x) > maxX {
			maxX = int(p.x)
		}
		if int(p.x) < minX || minX == 0 {
			minX = int(p.x)
		}
		if int(p.y) > maxY {
			maxY = int(p.y)
		}
		if int(p.y) < minY || minY == 0 {
			minY = int(p.y)
		}
	}
	// credit goes to https://github.com/atssteve/advent_of_code_2018/tree/master/day_6/part_1
	// then it turns out the grader was not working correctly after checking reddit
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			dup := false
			minD := 0.0
			idx := 0
			for i, p := range points {
				d := manhattanDistance(x, y, p)
				if i == 0 {
					minD = d
					continue
				}
				if d == 0 {
					idx = i
					break
				} else if d < minD && dup == true {
					minD = d
					dup = false
					idx = i
				} else if d < minD {
					minD = d
					idx = i
				} else if d == minD {
					dup = true
				} else {
					continue
				}
			}
			if dup == true {
				continue
			} else if x == 0 || y == 0 || x == maxX || y == maxY {
				points[idx].area++
				if isWinner(idx, disqualified) {
					disqualified = append(disqualified, idx)
				}
			} else {
				points[idx].area++
			}
		}
	}
	fmt.Println(points)

	largestArea := 0
	winner := 0
	for i, p := range points {
		if p.area > largestArea {
			if isWinner(i, disqualified) {
				winner = i
				largestArea = p.area
			}
		}
	}
	fmt.Println("------")
	fmt.Println(winner, points[winner])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
