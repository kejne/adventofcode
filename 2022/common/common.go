package common

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ReadInput(inputFile string) []string {
	file, error := os.Open(inputFile)
	
	if error != nil {
		log.Fatal("Failed to open input!")
	}

	defer func() {
		if error := file.Close(); error != nil {
			log.Fatal("Error closing file!")
		}
	}()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func GetNumber(line string) int {
	depth, err := strconv.Atoi(line)
	if err != nil {
		fmt.Println("Input can only be numbers")
		os.Exit(1)
	}
	return depth
}

