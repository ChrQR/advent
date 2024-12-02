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

func isValidSequence(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	increasing := nums[1] > nums[0]
	for i := 1; i < len(nums); i++ {
		diff := toPositive(nums[i] - nums[i-1])

		if diff < 1 || diff > 3 {
			return false
		}

		if increasing && nums[i] <= nums[i-1] {
			return false
		}
		if !increasing && nums[i] >= nums[i-1] {
			return false
		}
	}
	return true
}

func isReportSafe(nums []int) bool {
	if isValidSequence(nums) {
		return true
	}

	for i := 0; i < len(nums); i++ {
		withoutCurrent := make([]int, 0, len(nums)-1)
		withoutCurrent = append(withoutCurrent, nums[:i]...)
		withoutCurrent = append(withoutCurrent, nums[i+1:]...)

		if isValidSequence(withoutCurrent) {
			return true
		}
	}

	return false
}

func main() {
	file, err := os.Open("../p1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reports := bufio.NewScanner(file)
	safeReports := 0

	for reports.Scan() {
		fields := strings.Fields(reports.Text())
		nums := make([]int, len(fields))

		for i, v := range fields {
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			nums[i] = num
		}

		if isReportSafe(nums) {
			safeReports++
		}
	}

	fmt.Printf("There are %d safe reports\n", safeReports)
}

func toPositive(n int) int {
	return int(math.Abs(float64(n)))
}
