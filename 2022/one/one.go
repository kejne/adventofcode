package main

import (
	"fmt"
	"github.com/kejne/adventofcode/2022/common"
	"os"
	"strconv"
)

func calculateIncrements(lines []string) int {

	prevDepth := 99999999
	noOfIncrements := 0

	for _, line := range lines {
		depth, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Input can only be numbers")
			os.Exit(1)
		}

		if depth > prevDepth {
			noOfIncrements++
		}
		prevDepth = depth
	}
	return noOfIncrements
}

func main() {

	lines := common.ReadInput("input1.txt")

	noOfIncrements := calculateIncrements(lines)
	fmt.Printf("Amount of increments: %d ", noOfIncrements)
}
