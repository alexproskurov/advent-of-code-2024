package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//TODO: Optimize to reduce memory usage and enhance performance.
func main() {
	total, err := PartOne("input.txt")
	if err != nil {
			log.Fatalf("Error processing file: %v", err)
	}
	fmt.Println(total)

	total, err = PartTwo("input.txt")
	if err != nil {
		log.Fatalf("Error processing file: %v", err)
	}
	fmt.Println(total)
}

func PartOne(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	r := regexp.MustCompile(`mul\(\d+,\d+\)`)
	instructions := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions = append(instructions, r.FindAllString(scanner.Text(), -1)...)
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return CalculateSum(instructions)
}

func PartTwo(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	r := regexp.MustCompile(`mul\(\d+,\d+\)|don't\(\)|do\(\)`)
	instructions := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions = append(instructions, r.FindAllString(scanner.Text(), -1)...)
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return CalculateSum(instructions)
}

func CalculateSum(instructions []string) (int, error) {
	sum := 0
	skipOrWrite := true
	for _, val := range instructions {
		if strings.HasPrefix(val, "do") {
			if val == "don't()" {
				skipOrWrite = false
			} else {
				skipOrWrite = true
			}
			continue
		}

		if skipOrWrite {
			nums := strings.Split(val[4:len(val)-1], ",")

			num1, err := strconv.Atoi(nums[0])
			if err != nil {
				return 0, err
			}
			num2, err := strconv.Atoi(nums[1])
			if err != nil {
				return 0, err
			}

			sum += num1 * num2
		}
	}

	return sum, nil
}
