// 463 players; last marble is worth 71787 points

package main

import "fmt"

func shiftTwoSpotsFromCurrent(marbleIndex *int, marbles []int) {
	newIdx := *marbleIndex + 2
	if newIdx > len(marbles)-1 {
		*marbleIndex = 1
	} else {
		*marbleIndex = *marbleIndex + 2
	}
}
func shiftSevenCounterSpotsFromCurrent(marbleIndex *int, marbles []int) {

}
func main() {
	players := make([]int, 10)
	// 463
	lastMarbleScore := 1618
	// 71787
	currentMarbleScore := 0
	playerTurn := 0
	marbles := make([]int, 0)
	currentMarbleIndex := 0

	for currentMarbleScore != lastMarbleScore {
		if currentMarbleScore == 0 {
			marbles = append(marbles, currentMarbleScore)
		} else {
			if currentMarbleScore%23 == 0 {
				players[playerTurn] = players[playerTurn] + currentMarbleScore
				shiftSevenCounterSpotsFromCurrent(&currentMarbleIndex, marbles)
			}
			shiftTwoSpotsFromCurrent(&currentMarbleIndex, marbles)
			rightEnd := marbles[currentMarbleIndex:]
			marbles = append(marbles[:currentMarbleIndex], currentMarbleScore)
			marbles = append(marbles, rightEnd...)
			fmt.Println(marbles)

		}
		currentMarbleScore++
		playerTurn++
	}
}
