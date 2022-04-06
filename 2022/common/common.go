package common

import (
	"bufio"
	"log"
	"os"
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
