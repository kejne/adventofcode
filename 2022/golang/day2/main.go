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
	return s.basePoint()
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
