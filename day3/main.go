package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func exists(a string, s []string) bool {
	for i := range s {
		if a == s[i] {
			return true
		}
	}
	return false
}

func common(a []string) string {
	if len(a) != 3 {
		log.Printf("invalid group %s\n", a)
		return ""
	}

	first := strings.Split(a[0], "")
	second := strings.Split(a[1], "")
	third := strings.Split(a[2], "")

	for i := range third {
		if exists(third[i], first) && exists(third[i], second) {
			return third[i]
		}
	}
	return ""
}

func priority(a string) int {
	lc := strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	for p := range lc {
		if a == lc[p] {
			return p + 1
		}
	}

	uc := strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
	for q := range uc {
		if a == uc[q] {
			return q + 27
		}
	}
	return -1
}

func one(d []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(d))
	sum := 0

	for scanner.Scan() {
		items := scanner.Text()
		c1, c2 := strings.Split(items[:len(items)/2], ""), strings.Split(items[len(items)/2:], "")
		var com string
		for i := range c2 {
			if exists(c2[i], c1) {
				com = c2[i]
				break
			} else {
				com = ""
			}
		}
		if com == "" {
			log.Fatalf("no common characters in %s %s", c1, c2)
		}
		p := priority(com)
		if p == -1 {
			log.Fatalf("unable to determine priority score for %s", com)
		}
		sum = sum + p
	}
	return sum
}

func two(d []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(d))
	idx, sum := 1, 0
	group := []string{}
	for scanner.Scan() {
		items := scanner.Text()
		if idx%3 == 0 {
			group = append(group, items)
			com := common(group)
			if com == "" {
				log.Fatalf("no common characters in %s", group)
			}
			p := priority(com)
			if p == -1 {
				log.Fatalf("unable to determine priority score for %s", com)
			}
			sum = sum + p
			group = nil // clear for holding next group
		} else {
			group = append(group, items)
		}
		idx = idx + 1
	}
	return sum
}

func main() {
	data, err := os.ReadFile("input.txt") // read input from file
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: Sum of priorities of items that appear in both rugsacks : %d\n", one(data))
	fmt.Printf("Part 2: Sum of priorities of badge items : %d\n", two(data))
}
