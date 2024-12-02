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

	file, err := os.Open("../p1/input.txt")
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
	totalSimilarity := 0
	for _, v := range listA {
		occurances := 0
		for _, n := range listB {
			if v == n {
				occurances += 1
			}
		}
		totalSimilarity += v * occurances
	}
	fmt.Printf("The total similarity is: %d", totalSimilarity)
}
