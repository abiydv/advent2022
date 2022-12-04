package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type elf struct {
	position int
	total    int
}

func totals(d []byte) []elf {
	scanner := bufio.NewScanner(bytes.NewReader(d))
	pos, total := 1, 0
	elfs := []elf{}

	for scanner.Scan() {
		l, err := strconv.Atoi(scanner.Text())
		total = total + l
		if err != nil { // empty line
			elfs = append(elfs, elf{position: pos, total: total})
			pos, total = pos+1, 0 // increment position, reset total
		}
	}
	if total != 0 { // append the last elf to the slice
		elfs = append(elfs, elf{position: pos, total: total})
	}

	sort.Slice(elfs, func(i int, j int) bool { return elfs[i].total > elfs[j].total })

	return elfs
}

func two(d []byte) ([]int, int) {
	elfs := totals(d)
	return []int{elfs[0].position, elfs[1].position, elfs[2].position}, elfs[0].total + elfs[1].total + elfs[2].total
}
func one(d []byte) (int, int) {
	elfs := totals(d)
	return elfs[0].position, elfs[0].total
}

func main() {
	data, err := os.ReadFile("input.txt") // read input from file
	if err != nil {
		log.Fatal(err)
	}

	position, maxTotal := one(data)
	fmt.Printf("Part 1: Most calories carried by Elf %d: %d\n", position, maxTotal)
	positions, maxTotal := two(data)
	fmt.Printf("Part 2: Top 3 Elves carrying the most Calories: %d, they are carrying %d calories in total.\n", positions, maxTotal)
}
