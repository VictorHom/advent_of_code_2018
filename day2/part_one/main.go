package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getRepeatCountLetters(row string) (bool, bool) {
	m := make(map[string]int)
	var twoCount = false
	var threeCount = false

	for _, val := range row {
		m[string(val)]++
	}

	for k := range m {
		if m[k] == 2 {
			twoCount = true
		}
		if m[k] == 3 {
			threeCount = true
		}
	}
	return twoCount, threeCount
}

func main() {
	twoCount := 0
	threeCount := 0
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		containsTwo, containsThree := getRepeatCountLetters(row)
		if containsTwo {
			twoCount++
		}
		if containsThree {
			threeCount++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(twoCount * threeCount)
}
