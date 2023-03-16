package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
  file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	
  maxCalorie := 0
  currentElf := 0
  for scanner.Scan() {
    if calorie, err := strconv.Atoi(scanner.Text()); err == nil {
      currentElf += calorie
    } else {
      if maxCalorie < currentElf {
        maxCalorie = currentElf
      }
      currentElf = 0
    }
	}
  fmt.Println(maxCalorie)
}