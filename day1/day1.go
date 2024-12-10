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

//TODO: Refactor this code to match the structure and organization
// used in day2.go for consistency and clarity.
func main() {
	calc, err := ProcessFile("input.txt")
	if err != nil {
		log.Fatalf("Error processing file: %v", err)
	}
	DisplayCalculations(calc)
}

type Calculations struct {
	TotalDistance   int
	SimilarityScore int
}

type Pair struct {
	Left  int
	Right int
}

func ProcessFile(filename string) (*Calculations, error) {
	pairs, err := ExtractPairsFromFile(filename)
	if err != nil {
		return nil, err
	}

	totalDistance := CalculateTotalDistance(pairs)
	similarityScore := CalculateSimilarityScore(pairs)

	return &Calculations{
		TotalDistance:   totalDistance,
		SimilarityScore: similarityScore,
	}, nil
}

func ExtractPairsFromFile(filename string) ([]Pair, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var pairs []Pair
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pair, err := ConvertLineToPair(scanner.Text())
		if err != nil {
			return nil, err
		}
		pairs = append(pairs, pair)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return pairs, nil
}

func ConvertLineToPair(line string) (Pair, error) {
	vals := strings.Fields(line)
	if len(vals) > 2 {
		return Pair{}, fmt.Errorf("invalid line: %s", line)
	}

	left, err := strconv.Atoi(vals[0])
	if err != nil {
		return Pair{}, err
	}
	right, err := strconv.Atoi(vals[1])
	if err != nil {
		return Pair{}, err
	}

	return Pair{
		Left:  left,
		Right: right,
	}, nil
}

func CalculateTotalDistance(pairs []Pair) int {
	sum := 0
	for i := 0; i < len(pairs); i++ {
		sum += int(math.Abs(float64(pairs[i].Right - pairs[i].Left)))
	}
	return sum
}

func CalculateSimilarityScore(pairs []Pair) int {
	m := make(map[int]int, len(pairs))
	for _, pair := range pairs {
		m[pair.Right]++
	}

	sum := 0
	for _, pair := range pairs {
		sum += pair.Left * m[pair.Left]
	}

	return sum
}

func DisplayCalculations(c *Calculations) {
	fmt.Printf("Total Distance: %d\nSimilarity Score: %d\n", c.TotalDistance, c.SimilarityScore)
}
