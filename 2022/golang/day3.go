package main

import (
	"bufio"
	"strconv"
	"strings"
)

type Day3Calculator struct {
	input *bufio.Scanner
}

type Day3_2Calculator struct {
	input *bufio.Scanner
}

func NewDay3Calculator(input *bufio.Scanner) AdventCalculator {
	return Day3Calculator{
		input: input,
	}
}

func NewDay3_2Calculator(input *bufio.Scanner) AdventCalculator {
	return Day3_2Calculator{
		input: input,
	}
}

func (c Day3Calculator) GetResult() string {
	prios := strings.Split("0abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
	allPrios := 0
	for c.input.Scan() {
		r := c.input.Text()
		comp := len(r) / 2

		c1 := strings.Split(r, "")[:comp]
		c2 := strings.Split(r, "")[comp:]
		var both string
		for _, f := range c1 {
			for _, s := range c2 {
				if f == s && !strings.Contains(both, f) {
					both += f
				}
			}
		}
		for prio, l := range prios {
			if l == both {
				allPrios += prio
			}
		}
	}
	return strconv.Itoa(allPrios)
}

func (c Day3_2Calculator) GetResult() string {
	group := 1
	first := true
	var groupBadges []string
	var foundInAll string
	for c.input.Scan() {
		var sameAsGroup []string
		r := c.input.Text()
		if first {
			foundInAll = r
		} else {
			for _, i := range strings.Split(r, "") {
				if strings.Contains(foundInAll, i) {
					sameAsGroup = append(sameAsGroup, i)
				}
			}
		}
		grougasiiiirappend(groupBadges, sameAsGroup[0])
		group++
		if group%3 == 0 {
			first = true
			foundInAll = ""
		} else {
			first = false
		}
	}
	return "123"
}
