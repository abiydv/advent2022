package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func outerElements(a string) (int, int) {
	min, err := strconv.Atoi(strings.Split(a, "-")[0])
	if err != nil {
		log.Fatalf("cannot determine min outer element of range [min-max]: %s", a)
	}
	max, err := strconv.Atoi(strings.Split(a, "-")[1])
	if err != nil {
		log.Fatalf("cannot determine max outer element of range [min-max]: %s", a)
	}
	return min, max

}

func totalOverlap(a string, b string) bool {
	aMin, aMax := outerElements(a)
	bMin, bMax := outerElements(b)

	if aMin >= bMin && aMax <= bMax {
		return true
	} else if bMin >= aMin && bMax <= aMax {
		return true
	}
	return false
}

func partialOverlap(a string, b string) bool {
	aMin, aMax := outerElements(a)
	bMin, bMax := outerElements(b)

	if (aMin >= bMin && aMin <= bMax) || (aMax >= bMin && aMax <= bMax) {
		return true
	} else if (bMin >= aMin && bMin <= aMax) || (bMax >= aMin && bMax <= aMax) {
		return true
	}
	return false
}

func one(d []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(d))
	count := 0
	for scanner.Scan() {
		items := scanner.Text()
		assignments := strings.Split(items[:], ",")
		if totalOverlap(assignments[0], assignments[1]) {
			count = count + 1
		}
	}
	return count
}

func two(d []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(d))
	count := 0
	for scanner.Scan() {
		items := scanner.Text()
		assignments := strings.Split(items[:], ",")
		if partialOverlap(assignments[0], assignments[1]) {
			count = count + 1
		}
	}
	return count
}

func main() {
	data, err := os.ReadFile("input.txt") // read input from file
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: Assignment pairs where one range fully contains the other: %d\n", one(data))
	fmt.Printf("Part 2: Assignment pairs where one range partially contains the other: %d\n", two(data))
}
