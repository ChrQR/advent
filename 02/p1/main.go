package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reports := bufio.NewScanner(file)
	safeReports := 0

	for reports.Scan() {
		fields := strings.Fields(reports.Text())
		previous := 0
		unsafe := false
		increasing := false
		decreasing := false
		diff := 0
		for i, v := range fields {
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			if i == 0 {
				previous = num
        continue
			}
			if i > 0 {
				diff = toPositive(num - previous)
				if diff > 3 || diff < 1 {
					unsafe = true
					break
				}
				if num > previous {
					increasing = true
				}
				if num < previous {
					decreasing = true
				}
				if increasing && decreasing {
					unsafe = true
					break
				}
        previous = num
			}
		}
		if !unsafe {
			safeReports++
		}
	}
	fmt.Printf("There are %d safe reports", safeReports)
}

func toPositive(n int) int {
	return int(math.Abs(float64(n)))
}
