package main

import (
	"fmt"
	"github.com/kejne/adventofcode/2022/common"
)

func calculateIncrements(lines []string) int {

	prevDepth := 99999999
	noOfIncrements := 0

	for _, line := range lines {
		depth := common.GetNumber(line)

		if depth > prevDepth {
			noOfIncrements++
		}
		prevDepth = depth
	}
	return noOfIncrements
}

func calculateSlidingWindows(lines []string) []int {
	type window struct {
		values []int
		active bool
	}

	windows := make([]*window, 3)
	for _,window := range windows {
		window.values = make([]int, 0)
		window.active = false
	}

	windowSums := make([]int, 0)

	for _, line := range lines {
		depth := common.GetNumber(line)
		valueAdded := false
		for _, window := range windows {
			if !window.active {
				window.values = append(window.values, depth)
				window.active = true
			} else if !valueAdded {
				window.values = append(window.values, depth)
				valueAdded = true
			}
			
			
			if len(window.values) == 3 {
				var windowSum int
				for _, v := range window.values {
					windowSum += v
				}
				windowSums = append(windowSums, windowSum)
				window.values = make([]int, 0)
				window.active = false
			}
		}

	}
	return windowSums
}

func main() {

	lines := common.ReadInput("input1.txt")

	noOfIncrements := calculateIncrements(lines)
	fmt.Printf("Amount of increments: %d ", noOfIncrements)

	windowSums := calculateSlidingWindows(lines)
	fmt.Printf("WindowSums: %v", windowSums)
}
