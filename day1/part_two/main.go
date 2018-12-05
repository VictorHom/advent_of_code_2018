package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	freqCount := make(map[int]int)
	freq := 0
	freqCount[freq] = 1
	notFound := false
	count := 0
	for !notFound {
		file, err := os.Open("./input.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		count++
		for scanner.Scan() {
			row := scanner.Text()
			n, _ := strconv.Atoi(row)
			freq = freq + n
			_, ok := freqCount[freq]
			if ok {
				fmt.Println(freq)
				fmt.Println("count", count)
				notFound = true
				os.Exit(1)
			} else {
				freqCount[freq] = 0
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
