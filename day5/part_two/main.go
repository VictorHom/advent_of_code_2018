package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func reactive(a string, b string) bool {
	if a == b {
		return false
	}
	return strings.ToLower(a) == b || strings.ToUpper(a) == b
}

func collapseReact(polymerUnits []string) int {
	ptr := 0
	lookaheadPtr := ptr + 1
	for ptr < len(polymerUnits) && lookaheadPtr < len(polymerUnits) {
		if !reactive(polymerUnits[lookaheadPtr], polymerUnits[lookaheadPtr-1]) {
			ptr++
			lookaheadPtr = ptr + 1
		} else {
			polymerUnits = append(polymerUnits[:lookaheadPtr-1], polymerUnits[lookaheadPtr+1:]...)
			ptr = 0
			lookaheadPtr = ptr + 1
		}

	}
	return len(polymerUnits)
}

func removeUnits(p []string, c string) []string {
	str := strings.Replace(strings.Join(p, ""), c, "", -1)
	str = strings.Replace(str, strings.ToLower(c), "", -1)
	return strings.Split(str, "")
}

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	polymerUnits := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		polymerUnits = strings.Split(string(row), "")
	}

	alphabets := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Println(len(alphabets))

	for _, char := range alphabets {
		fmt.Println(string(char))
		updatedPolymerUnits := removeUnits(polymerUnits, string(char))
		fmt.Println(collapseReact(updatedPolymerUnits))
	}

	fmt.Println(len(polymerUnits))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
