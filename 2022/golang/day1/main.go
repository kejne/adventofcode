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

  sort.Slice(elfCalorie, func(i, j int) bool {
    return elfCalorie[i] < elfCalorie[j]
  })
  
  fmt.Println(elfCalorie)
}