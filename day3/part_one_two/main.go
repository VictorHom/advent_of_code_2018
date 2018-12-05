package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Rect to represent the fabric location and patch
type Rect struct {
	id int
	x  int
	y  int
	w  int
	h  int
}

func overlappingArea(rect1 Rect, rect2 Rect) bool {
	// get leftmost sides
	if rect2.x < rect1.x {
		rect1.x, rect2.x = rect2.x, rect1.x
		rect1.w, rect2.w = rect2.w, rect1.w
	}
	if rect2.y < rect1.y {
		rect1.y, rect2.y = rect2.y, rect1.y
		rect1.h, rect2.h = rect2.h, rect1.h
	}

	if rect1.x+rect1.w > rect2.x && rect1.y+rect1.h > rect2.y {
		return true
	}
	return false
}

func split(r rune) bool {
	return r == '#' || r == '@' || r == ':' || r == 'x' || r == ','
}
func main() {
	claims := make([]Rect, 1)
	reInsideWhtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		a := strings.FieldsFunc(string(row), split)

		b := reInsideWhtsp.ReplaceAllString(strings.Join(a, " "), " ")
		s := strings.Split(b, " ")
		id, _ := strconv.Atoi(s[0])
		x, _ := strconv.Atoi(s[1])
		y, _ := strconv.Atoi(s[2])
		w, _ := strconv.Atoi(s[3])
		h, _ := strconv.Atoi(s[4])
		r := Rect{id: id, x: x, y: y, w: w, h: h}
		claims = append(claims, r)
	}

	// fmt.Println(claims)
	var fabric [1000][1000]int
	overlapCount := 0
	for _, claim := range claims {
		for i := claim.x; i < claim.x+claim.w; i++ {
			for j := claim.y; j < claim.y+claim.h; j++ {
				if fabric[i][j] == 0 {
					fabric[i][j] = 1
				} else if fabric[i][j] == 1 {
					overlapCount++
					fabric[i][j]++
				}
			}
		}
	}
	fmt.Println(overlapCount)
	// fmt.Println(claims)
	for _, claim := range claims {
		hasOverlap := false
		for i := claim.x; i < claim.x+claim.w; i++ {
			for j := claim.y; j < claim.y+claim.h; j++ {
				if fabric[i][j] > 1 {
					hasOverlap = true
				}
			}
		}
		if hasOverlap {
			// check next claim
			hasOverlap = false
		} else {
			fmt.Println(claim.id)
		}
	}

	// for every claim just check on the other to see if overlap
	// my overlap fxn is off; the result is too low
	// for i, claim := range claims {
	// 	for j := i + 1; j < len(claims); j++ {
	// 		if overlappingArea(claim, claims[j]) {
	// 			overlapCount++
	// 		}
	// 	}
	// }

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
