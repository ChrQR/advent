package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	listA := []int{}
	listB := []int{}
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), "   ")
		num1, err1 := strconv.Atoi(fields[0])
		num2, err2 := strconv.Atoi(fields[1])
		if err1 != nil || err2 != nil {
			log.Fatal(err1, err2)
		}
		listA = append(listA, num1)
		listB = append(listB, num2)
	}
	sort.Slice(listA, func(i, j int) bool {
		return listA[i] < listA[j]
	})
  sort.Slice(listB, func(i, j int) bool {
		return listB[i] < listB[j]
	})
	distance := 0
	for i := range listA {
		if listA[i] > listB[i] {
			distance += listA[i] - listB[i]
		} else {
			distance += listB[i] - listA[i]
		}
	}
	fmt.Printf("The total distance is: %d", distance)
}
