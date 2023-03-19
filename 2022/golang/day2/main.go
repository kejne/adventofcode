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
		strategy := Strategy(scanner.Text())
		totPoints += strategy.evalPoints(day2)
	}
	fmt.Println(totPoints)
}

type Strategy string 

type Day int

const (
	day1 Day = iota
	day2
)

type Move int

const (
	rock Move = iota
	paper
	scissors
)

func (m Move) Points() int {
	points := []int{1,2,3}
	return points[int(m)]
}

func (s Strategy) evalPoints(day Day) int {
	if day == day1 {
		return s.basePoint() + s.battle()
	} else if day == day2 {
		return s.evalMove().Points() + s.OutcomePoint()
	} else {
		panic("Not implemented yet!")
	}
}

func (s Strategy) evalMove() Move {
	moves := map[string]Move{
		"A X": scissors,
		"A Y": rock,
		"A Z": paper,
		"B X": rock,
		"B Y": paper,
		"B Z": scissors,
		"C X": paper,
		"C Y": scissors,
		"C Z": rock,
	}

	for strategy,move := range moves {
		if strings.Contains(string(s), strategy) {
			return move
		}
	}

	panic("No move found for row <" + s + ">")

}

func (s Strategy) battle() int {
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

func (s Strategy) OutcomePoint() int {
	pointMap := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	for spr, points := range pointMap {
		if strings.Contains(string(s), spr) {
			return points
		}
	}

	return 0
}


func (s Strategy) basePoint() int {
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
