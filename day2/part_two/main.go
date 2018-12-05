package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func isSimilar(a string, b string) bool {
	countDifference := 0
	if len(a) != len(b) {
		return false
	}
	for i, x := range a {
		if string(x) != string(b[i]) {
			countDifference++
			if countDifference == 2 {
				return false
			}
		}
	}
	fmt.Println(a)
	fmt.Println(b)
	return true
}

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	list := make([]string, 1)
	for scanner.Scan() {
		row := scanner.Text()
		list = append(list, string(row))
	}
	for i := 0; i < len(list); i++ {
		a := list[i]
		for j := i + 1; j < len(list); j++ {
			isSimilar(a, list[j])
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
