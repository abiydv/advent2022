package main

import (
	"fmt"
	"log"
	"os"
)

// charPresent returns true if a char is present in a string
func charPresent(val string, slc string) bool {
	for _, v := range slc {
		if val == string(v) {
			return true
		}
	}
	return false
}

// uniqueValues returns true if a string has unique chars
func uniqueChars(slc string) bool {
	for k, v := range slc {
		// brute force comparison with all other elements except itself
		if charPresent(string(v), slc[:k]+slc[k+1:]) {
			return false
		}
	}
	fmt.Printf("unique chars in %v \n", slc)
	return true
}

// find start-of-message marker
func two(d []byte) int {
	// start counter from 1
	for k := 1; k < len(d); k++ {
		if k-14 < 0 {
			// not enough chars to compare
			continue
		}
		if uniqueChars(string(d[k-14 : k])) {
			return k
		}
	}
	return -1
}

// finds start-of-packet marker
func one(d []byte) int {
	// start counter from 1
	for k := 1; k < len(d); k++ {
		if k-4 < 0 {
			// not enough chars to compare
			continue
		}
		if uniqueChars(string(d[k-4 : k])) {
			return k
		}
	}
	return -1
}

func main() {
	data, err := os.ReadFile("input.txt") // read input from file
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: First start-of-packet marker detected after processing %d characters\n", one(data))
	fmt.Printf("Part 2: First start-of-message marker detected after processing %d characters\n", two(data))
}
