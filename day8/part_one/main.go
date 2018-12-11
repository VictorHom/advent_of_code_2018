package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type node struct {
	childrenCount int
	metadataCount int
	children      *[]node
	metadata      []int
}

func createTree(n *[]string, total *int) node {
	if len(*n) == 0 {
		return node{}
	}
	numberOfChilren, _ := strconv.Atoi((*n)[0])
	numberOfMetadataEntries, _ := strconv.Atoi((*n)[1])
	*n = (*n)[2:]

	children := make([]node, 0)
	for i := 0; i < numberOfChilren; i++ {
		children = append(children, createTree(n, total))
	}

	metaDataEntries := make([]int, 0)
	for j := 0; j < numberOfMetadataEntries; j++ {
		entry, _ := strconv.Atoi((*n)[0])
		*total = *total + entry
		metaDataEntries = append(metaDataEntries, entry)
		*n = (*n)[1:]
	}

	return node{children: &children, metadata: metaDataEntries}
}

func findMetaDataIndexSum(headNode node, metaDataIndexSum *int) int {
	metaDataEntries := headNode.metadata
	children := headNode.children
	sum := 0
	if len(*children) == 0 {
		for i := 0; i < len(metaDataEntries); i++ {
			sum += metaDataEntries[i]
		}
		*metaDataIndexSum = *metaDataIndexSum + sum
		return sum
	}

	for i := 0; i < len(metaDataEntries); i++ {
		index := metaDataEntries[i]
		if index > len(*children) == false {
			findMetaDataIndexSum((*children)[index-1], metaDataIndexSum)
		}
	}

	return *metaDataIndexSum
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
	total := 0
	headNode := createTree(&numbers, &total)

	metaDataIndexSum := 0
	m := findMetaDataIndexSum(headNode, &metaDataIndexSum)

	fmt.Println(headNode)
	fmt.Println(total)
	fmt.Println(metaDataIndexSum)
	fmt.Println(m)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
