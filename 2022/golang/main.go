package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	daystr := os.Args[1]
	day, _ := strconv.Atoi(daystr)
	fmt.Printf("Day %d result is: %s\n", day, startCalculation(day))
}

func startCalculation(day int) string {
	input := fmt.Sprintf("inputs/day%d.input", day)
	file, _ := os.Open(input)
	defer file.Close()
	return getCalculator(file, day).GetResult()
}

func getCalculator(file *os.File, day int) AdventCalculator {
	scanner := bufio.NewScanner(file)
	switch day {
	case 1:
		return NewDay1Calculator(scanner)
	case 2:
		return NewDay2Calculator(scanner)
	case 3:
		return NewDay3Calculator(scanner)
	default:
		return NewDay1Calculator(scanner)
	}
}
