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
	totalSafeLevels, err := PartOne("input.txt")
	if err != nil {
		log.Fatalf("Error processing file: %v", err)
	}
	fmt.Println(totalSafeLevels)

	totalToleratedLevels, err := PartTwo("your_file.txt")
	if err != nil {
		log.Fatalf("Error processing file: %v", err)
	}
	fmt.Println(totalToleratedLevels)
}

func PartOne(filename string) (int, error) {
	levels, err := ExtractLevels(filename)
	if err != nil {
		return 0, err
	}
	return CalculateSafeLevels(levels)
}

func PartTwo(filename string) (int, error) {
	levels, err := ExtractLevels(filename)
	if err != nil {
		return 0, err
	}
	return CalculateToleratedLevels(levels)
}

func ExtractLevels(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	levels := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		levels = append(levels, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return levels, nil
}

func CalculateSafeLevels(levels []string) (int, error) {
	res := 0
	for _, lvl := range levels {
		nums, err := convertLevelToIntSlice(lvl)
		if err != nil {
			return 0, err
		}
		if IsSafeLevel(nums) {
			res++
		}
	}
	return res, nil
}

func convertLevelToIntSlice(lvl string) ([]int, error) {
	vals := strings.Fields(lvl)
	rv := make([]int, len(vals))
	for i, val := range vals {
		num, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		rv[i] = num
	}

	return rv, nil
}

func IsSafeLevel(nums []int) bool {
	increasing, decreasing := true, true
	for i := 0; i < len(nums)-1; i++ {
		diff := int(math.Abs(float64(nums[i+1] - nums[i])))
		if diff > 3 || diff == 0 {
			return false
		}

		if nums[i] > nums[i+1] {
			increasing = false
		}
		if nums[i] < nums[i+1] {
			decreasing = false
		}
	}

	return increasing || decreasing
}

// TODO: Optimize this code to reduce memory usage by revisiting data structures and processing logic.
func CalculateToleratedLevels(levels []string) (int, error) {
	res := 0
	for _, lvl := range levels {
		nums, err := convertLevelToIntSlice(lvl)
		if err != nil {
			return 0, err
		}
		if IsToleratedLevel(nums) {
			res++
		}
	}
	return res, nil
}

// TODO: Optimize this code to reduce memory usage by revisiting data structures and processing logic.
func IsToleratedLevel(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		modified := make([]int, 0, len(nums)-1)
		modified = append(modified, nums[:i]...)
		modified = append(modified, nums[i+1:]...)

		if IsSafeLevel(modified) {
			return true
		}
	}
	return false
}
