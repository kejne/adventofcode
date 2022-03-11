package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, error := os.Open("input1.txt")
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

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
