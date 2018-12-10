package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type node struct {
	childrenCount int
	metadataCount int
	children      *[]node
	metadata      []int
}

func createTree(n []string) node {
	return node{childrenCount: 0, metadataCount: 0}
}

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	numbers := make([]string, 0)
	for scanner.Scan() {
		data := scanner.Text()
		numbers = append(numbers, strings.Split(string(data), " ")...)
	}
	headNode := createTree(numbers)
	fmt.Println(headNode)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
