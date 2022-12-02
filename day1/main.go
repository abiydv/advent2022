package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func fSumSlice(scanner *bufio.Scanner) []int {
	var sum int
	var sumSlice []int

	for scanner.Scan() {
		d, err := strconv.Atoi(scanner.Text())
		sum = sum + d
		if err != nil { // empty line
			sumSlice = append(sumSlice, sum)
			sum = 0
		}
	}
	if sum != 0 { // append the last sum to the slice
		sumSlice = append(sumSlice, sum)
	}
	return sumSlice
}

func findMax(s []int) int {
	max := 0
	for i := range s {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}

func findPos(v int, s []int) int {
	idx := -1
	for i := range s {
		if s[i] == v {
			idx = i
		}
	}
	return idx
}

func findTopN(n int, s []int) []int {
	if n > len(s) || n == len(s) {
		return s
	}

	topN := []int{}
	ss := []int{}
	ss = append(ss, s[:]...) // don't modify the original slice

	for i := 0; i < n; i++ {
		m := findMax(ss)
		p := findPos(m, ss)
		topN = append(topN, m)
		ss = append(ss[:p], ss[p+1:]...)
	}
	return topN
}

func main() {
	data, err := os.ReadFile("input.txt") // read input from file
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(data))

	sumSlice := fSumSlice(scanner)

	// top := 1 // part 1
	top := 3 // part 2
	topVals := findTopN(top, sumSlice)

	var sum int
	var positions []int
	for i := range topVals {
		sum = sum + topVals[i]
		positions = append(positions, findPos(topVals[i], sumSlice))
	}
	fmt.Printf("Total calories %d carried by top %d Elfs %d\n", sum, top, positions)

}
