package main

import (
	"bufio"
	"sort"
	"strconv"
)

type Day1Calculator struct {
	input *bufio.Scanner
}

func NewDay1Calculator(input *bufio.Scanner) AdventCalculator {
	return Day1Calculator{
		input: input,
	}
}

func (c Day1Calculator) GetResult() string {
	calories := groupCalCarried(c.input)
	topCals := getTopCalories(3, calories)
	return strconv.Itoa(topCals)
}

func groupCalCarried(scanner *bufio.Scanner) []int {
	var elfCalorie []int

	currentElf := 0
	for scanner.Scan() {
		if calorie, err := strconv.Atoi(scanner.Text()); err == nil {
			currentElf += calorie
		} else {
			elfCalorie = append(elfCalorie, currentElf)
			currentElf = 0
		}
	}
	return elfCalorie
}

func getTopCalories(elfs int, elfCalorie []int) int {

	sort.Slice(elfCalorie, func(i, j int) bool {
		return elfCalorie[i] > elfCalorie[j]
	})

	totCal := 0
	for _, cal := range elfCalorie[:3] {
		totCal += cal
	}

	return totCal
}
