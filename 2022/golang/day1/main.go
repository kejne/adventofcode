package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	elfCalorie := groupCalCarried(scanner)

	fmt.Println(getTopCalories(3, elfCalorie))
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
