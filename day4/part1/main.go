package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func recordToSlice(r string) []string {
	split := strings.Split(r, "] ")
	d := make([]string, 0)
	dateTime := strings.Split(split[0], "[")
	time := strings.Split(strings.Join(dateTime, ""), " ")[1]
	d = append(d, strings.Join(dateTime, ""))
	d = append(d, split[1])
	d = append(d, time)
	return d
}

func getMinute(m string) string {
	return strings.Split(m, ":")[1]
}

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	records := make([][]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		recordSlice := recordToSlice(string(row))
		records = append(records, recordSlice)
	}

	// sort by the timetamp for each date
	sort.Slice(records, func(i, j int) bool {
		a, _ := time.Parse("2006-01-02 15:04", records[i][0])
		b, _ := time.Parse("2006-01-02 15:04", records[j][0])

		return a.Before(b)
	})
	fmt.Println(records)

	guardTotalSleep := make(map[string]int)
	guardIndividualMinuteCount := make(map[string]map[int]int)

	currentGuard := ""
	fallAsleepTime := ""
	for _, record := range records {
		if strings.Contains(record[1], "begins shift") {
			currentGuard = record[1]
		}
		if strings.Contains(record[1], "falls asleep") {
			fallAsleepTime = record[2]
		}
		if strings.Contains(record[1], "wakes up") {
			fallAsleepMinute, _ := strconv.Atoi(getMinute(fallAsleepTime))
			wakeAsleepMinute, _ := strconv.Atoi(getMinute(record[2]))
			for i := fallAsleepMinute; i < wakeAsleepMinute; i++ {
				if guardIndividualMinuteCount[currentGuard] == nil {
					guardIndividualMinuteCount[currentGuard] = make(map[int]int)
				} else {
					guardIndividualMinuteCount[currentGuard][i]++
				}
			}

			guardTotalSleep[currentGuard] += wakeAsleepMinute - fallAsleepMinute + 1
			fallAsleepTime = ""
		}
	}

	fmt.Println(guardTotalSleep)
	fmt.Println(guardIndividualMinuteCount)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
