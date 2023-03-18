package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totPoints := 0
	for scanner.Scan() {
		strategy := strategy(scanner.Text())
		totPoints += strategy.evalPoints()
	}
	fmt.Println(totPoints)
}

type strategy string

func (s strategy) evalPoints() int {
	return s.basePoint() + s.battle()
}

func (s strategy) battle() int {
	wins := map[string]string{
		"X": "C",
		"Y": "A",
		"Z": "B",
	}
	draw := map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
	}

	points := 0
	for spr, defeats := range wins {
		if strings.Contains(string(s), spr) {
			if strings.Contains(string(s), defeats) {
				points = 6
			} else if strings.Contains(string(s), draw[spr]) {
				points = 3
			}
		}
	}
	return points
}

func (s strategy) basePoint() int {
	pointMap := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	for spr, points := range pointMap {
		if strings.Contains(string(s), spr) {
			return points
		}
	}

	return 0
}
