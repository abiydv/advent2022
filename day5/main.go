package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func moveCrates(stacks [][]string, instructions [][]string) [][]string {
	for _, i := range instructions {
		move, _ := strconv.Atoi(i[1])
		frm, _ := strconv.Atoi(i[3])
		to, _ := strconv.Atoi(i[5])

		for m := 0; m < move; m++ {
			if len(stacks[frm-1]) == 0 { // no crates to move
				break
			}
			stacks[to-1] = append([]string{stacks[frm-1][0]}, stacks[to-1]...)
			stacks[frm-1] = stacks[frm-1][1:]
		}
	}
	return stacks
}

func moveCratesAdv(stacks [][]string, instructions [][]string) [][]string {
	for _, i := range instructions {
		move, _ := strconv.Atoi(i[1])
		frm, _ := strconv.Atoi(i[3])
		to, _ := strconv.Atoi(i[5])

		// find the least of slice len or batch size to avoid out of range errors
		min := int(math.Min(float64(move), float64(len(stacks[frm-1]))))

		// use a new slice to hold the batch, to avoid changing the underlying array
		var batch []string
		batch = append(batch, stacks[frm-1][:min]...)
		stacks[to-1] = append(batch, stacks[to-1]...)
		stacks[frm-1] = stacks[frm-1][min:]
	}
	return stacks
}

func removeNonLetters(s string) string {
	return strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) {
			return -1
		}
		return r
	}, s)
}

func topCrates(crates [][]string) string {
	top := []string{}
	i := 0
	for _, v := range crates {
		if len(v[i]) == 0 && i+1 <= len(v) {
			top = append(top, removeNonLetters(v[i+1]))
			continue
		}
		top = append(top, removeNonLetters(v[0]))
	}
	return strings.Join(top, "")
}

func stacks(crates [][]string) [][]string {
	stacks := [][]string{}
	max := math.Max(float64(len(crates)), float64(len(crates[0]))) // operate on the max length to cater for non-symetrical arrays
	for i := 0; i < int(max); i++ {
		stack := []string{}
		for _, v := range crates {
			if v[i] == "   " {
				continue
			}
			stack = append(stack, v[i])
		}
		stacks = append(stacks, stack)
	}
	return stacks
}

func prepareCratesAndInstruc(d []byte) ([][]string, [][]string) {
	allCrates := [][]string{}
	instructions := [][]string{}
	scanner := bufio.NewScanner(bytes.NewReader(d))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Split(line, " ")[0] == "move" {
			instructions = append(instructions, strings.Split(line, " "))
			continue
		}
		if len(line) == 0 {
			continue
		}
		crates := []string{}
		for i := 0; i < len(line); i += 4 {
			crates = append(crates, line[i:i+3])
		}
		allCrates = append(allCrates, crates)
	}
	stacks := stacks(allCrates[:len(allCrates)-1]) // ignore the last numerical element [ 1 2 3 ]
	fmt.Printf("stacks => %+v\n", stacks)
	return stacks, instructions
}

func one(d []byte) string {
	stacks, instructions := prepareCratesAndInstruc(d)

	moved := moveCrates(stacks, instructions)
	fmt.Printf("final stacks => %+v\n", moved)

	topCrates := topCrates(moved)
	fmt.Printf("top crates => %+s\n", topCrates)

	return topCrates
}

func two(d []byte) string {
	stacks, instructions := prepareCratesAndInstruc(d)

	moved := moveCratesAdv(stacks, instructions)
	fmt.Printf("final stacks => %+v\n", moved)

	topCrates := topCrates(moved)
	fmt.Printf("top crates => %+s\n", topCrates)

	return topCrates
}

func main() {
	data, err := os.ReadFile("input.txt") // read input from file
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: After the rearrangement procedure, crates on top of each stack: %s\n", one(data))
	fmt.Printf("Part 2: After the advanced rearrangement procedure, crates on top of each stack: %s\n", two(data))
}
