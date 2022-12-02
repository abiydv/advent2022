package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func one(d []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(d))
	scores := map[string]int{"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3}
	score := 0
	for scanner.Scan() {
		play := strings.Split(scanner.Text(), " ")
		score = score + scores[play[1]] // score for the shape played
		if scores[play[1]] == 1 && scores[play[0]] == 3 {
			score = score + 6 // score for WIN
		} else if scores[play[1]] == 3 && scores[play[0]] == 1 {
			score = score + 0 // score for LOSS
		} else if scores[play[1]] > scores[play[0]] {
			score = score + 6 // score for WIN
		} else if scores[play[1]] == scores[play[0]] {
			score = score + 3 // score for DRAW
		} else {
			score = score + 0 // score for LOSS
		}
	}
	return score
}

func two(d []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(d))
	scores := map[string]int{"A": 1, "B": 2, "C": 3, "X": 0, "Y": 3, "Z": 6}
	score := 0
	for scanner.Scan() {
		play := strings.Split(scanner.Text(), " ")
		score = score + scores[play[1]] // score for the outcome
		if scores[play[1]] == 6 && scores[play[0]] == 3 {
			score = score + 1 // score for WIN
		} else if scores[play[1]] == 6 {
			score = score + scores[play[0]] + 1 // score for WIN
		} else if scores[play[1]] == 0 && scores[play[0]] == 1 {
			score = score + 3 // score for LOSS
		} else if scores[play[1]] == 0 {
			score = score + scores[play[0]] - 1 // score for LOSS
		} else {
			score = score + scores[play[0]] // score for DRAW
		}
	}
	return score
}

func main() {

	data, err := os.ReadFile("input.txt") // read input from file
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: Total score : %d\n", one(data))
	fmt.Printf("Part 2: Total score : %d\n", two(data))
}
