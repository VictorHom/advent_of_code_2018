package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type step struct {
	first  string
	second string
}

func getFirstLetter(instructions []step) string {
	firstStepLetter := ""
	for i := 0; i < len(instructions); i++ {
		first := instructions[i].first
		firstStep := true
		for j := i; j < len(instructions); j++ {
			if first == instructions[j].second {
				firstStep = false
			}
		}
		if firstStep {
			firstStepLetter = first
			break
		}
	}
	// fmt.Println(firstStepLetter)
	return firstStepLetter
}

func findAllStepsWithFirstLetter(instructions *[]step, lastStepLetter string) []step {
	temporarySteps := make([]step, 0)
	for i := 0; i < len(*instructions); {
		// fmt.Println(lastStepLetter)
		if (*instructions)[i].first == lastStepLetter {
			temporarySteps = append(temporarySteps, (*instructions)[i])
			*instructions = append((*instructions)[:i], (*instructions)[i+1:]...)
		} else {
			i++
		}
	}

	sort.SliceStable(temporarySteps, func(i int, j int) bool {
		return temporarySteps[i].first > temporarySteps[j].first
	})

	return temporarySteps
}

func sortSteps(steps []step) []step {
	sort.SliceStable(steps, func(i int, j int) bool {
		return steps[i].first > steps[j].first
	})

	return steps
}

type worker struct {
	available bool
	workingOn string
	timeSpent int
}

func getMin(nums []int) int {
	min := nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}

func cleanUpWorkers(queue []string, workers *[]worker) {
	for _, w := range *workers {
		isTaskStillThere := false
		for _, task := range queue {
			if task == w.workingOn {
				isTaskStillThere = true
			}
		}
		if isTaskStillThere == false {
			// if task is no longer there, it means it's complete and we can clean up
			w.available = true
			w.workingOn = ""
		}
	}
}

func updateWorkersWithMinTime(time int, workers *[]worker) {
	tasks := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, w := range *workers {
		w.timeSpent = w.timeSpent + time
		if w.timeSpent >= strings.Index(tasks, w.workingOn)+1 {
			// we are done
			w.available = true
			w.workingOn = ""
		}
	}
}

func calcTime(queue []string, totalTime *int, workers *[]worker, seenTasks *map[string]int) {
	tasks := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	cleanUpWorkers(queue, workers)
	for len(queue) > 0 {
		task := queue[0]
		if _, ok := (*seenTasks)[task]; ok {
			// already seen, no reason to do any more calculations
			queue = queue[1:]
			continue
		}
		for _, worker := range *workers {
			times := make([]int, 0)
			if worker.available && len(queue) > 0 {
				worker.available = false
				worker.workingOn = task
				times = append(times, strings.Index(tasks, task)+1)
			}
			*totalTime = *totalTime + getMin(times)
			updateWorkersWithMinTime(getMin(times), workers)
			queue = queue[1:]
		}
	}
}

func reverse(input string) string {
	// fmt.Println("input", input)
	s := strings.Split(input, "")
	// fmt.Println(s)
	reversed := make([]string, 0)
	for j := len(s) - 1; j >= 0; j-- {
		reversed = append(reversed, s[j])
	}
	return strings.Join(reversed, "")
}

func main() {

	file, err := os.Open("./small.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	instructions := make([]step, 0)
	stepsWithOrder := make([]step, 0)
	fmt.Println(stepsWithOrder)

	for scanner.Scan() {
		row := string(scanner.Text())
		steps := strings.Split(row, " must be finished before step ")
		first := strings.Replace(steps[0], "Step ", "", -1)
		second := strings.Replace(steps[1], " can begin.", "", -1)
		s := step{first: first, second: second}
		instructions = append(instructions, s)
	}
	totalTime := 0
	workers := make([]worker, 0)
	workers = append(workers, worker{available: true, workingOn: ""})
	workers = append(workers, worker{available: true, workingOn: ""})
	seenTasks := make(map[string]int)
	// map the post activity to the pre steps
	graph := make(map[string][]step)
	current := ""
	solution := make([]string, 0)
	queue := make([]string, 0)
	for _, instruction := range instructions {
		_, ok := graph[instruction.first]
		if !ok {
			graph[instruction.first] = make([]step, 0)
		}
		graph[instruction.second] = append(graph[instruction.second], instruction)
	}
	fmt.Println("graph")
	fmt.Println(graph)

	for len(graph) > 0 {
		for k, v := range graph {
			for idx, step := range v {
				if step.first == string(current) {
					graph[k] = append(graph[k][:idx], graph[k][idx+1:]...)
				}
			}
			if len(graph[k]) == 0 {
				// fmt.Println(solution)
				// fmt.Println("zero", k, v)
				queue = append(queue, k)
			}
			go calcTime(queue, &totalTime, &workers, &seenTasks)
		}
		fmt.Println("queue", queue)
		sort.Strings(queue)
		current = queue[0]
		delete(graph, current)
		solution = append(solution, queue[0])
		queue = make([]string, 0)
	}
	fmt.Println(strings.Join(solution, ""))

	// fmt.Println(reverse(sol))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
