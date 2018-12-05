package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	b, err := os.Open("./input.txt")
	if err != nil {
		fmt.Print(err)
	}

	defer b.Close()
	total := 0
	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		row := scanner.Text()
		num, _ := strconv.Atoi(row)
		fmt.Println(num)
		total += num
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)

}
