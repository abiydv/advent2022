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

type node struct {
	name     string
	size     int
	in       *node
	contains map[string]*node
}

func one(d []byte, max int) int {
	dirs := prepareDirs(d)

	pickDirs := dirSmallerThan(dirs, 100000)
	var sum int
	for _, d := range pickDirs {
		sum += d.size
	}
	return sum
}

func dirSmallerThan(dirs []*node, limit int) []*node {
	subset := []*node{}
	for _, d := range dirs {
		s := size(d)
		if s <= limit {
			subset = append(subset, d)
		}
	}
	return subset
}

func two(d []byte) int {
	dirs := prepareDirs(d)

	pickDirs := dirBiggerThan(dirs, cleanupSpace(dirs[0]))

	return smallestDir(pickDirs).size
}

func cleanupSpace(root *node) int {
	total := 70000000
	required := 30000000
	used := root.size
	free := total - used
	return required - free
}

func dirBiggerThan(dirs []*node, limit int) []*node {
	subset := []*node{}
	for _, d := range dirs {
		s := size(d)
		if s >= limit {
			subset = append(subset, d)
		}
	}
	return subset
}

func smallestDir(dirs []*node) (ret *node) {
	min := dirs[0].size
	for _, d := range dirs {
		if d.size < min {
			min = d.size
			ret = d
		}
	}
	return
}

func prepareDirs(d []byte) (dirs []*node) {
	var c *node // hold the current node
	scanner := bufio.NewScanner(bytes.NewReader(d))

	for scanner.Scan() {
		l := strings.Fields(scanner.Text())

		if l[0] == "$" {
			if l[1] == "cd" {
				if l[2] == "/" { // root node
					c = &node{
						name:     l[2],
						size:     0,
						in:       nil,
						contains: make(map[string]*node),
					}
					dirs = append(dirs, c)
				} else if l[2] == ".." {
					c = c.in // move back 1 level
				} else {
					c = c.contains[l[2]] // move into  1 level
				}
			}
		} else if l[0] == "dir" {
			c.contains[l[1]] = &node{
				name:     l[1],
				size:     0, // fill later
				in:       c,
				contains: make(map[string]*node),
			}
			dirs = append(dirs, c.contains[l[1]])
			//fmt.Printf("pick %+v\n", *c)
		} else {
			size, err := strconv.Atoi(l[0])
			if err != nil {
				fmt.Printf("err getting size in int %s", err.Error())
			}
			c.contains[l[1]] = &node{
				name:     l[1],
				size:     size,
				in:       c,
				contains: nil,
			}
		}
	}
	for _, d := range dirs {
		d.size = size(d)
	}
	return
}

func size(n *node) int {
	if len(n.contains) == 0 { // file node
		return n.size
	}
	var s int
	for _, c := range n.contains {
		s += size(c)
	}
	return s
}

func main() {
	data, err := os.ReadFile("input.txt") // read input from file
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: Sum of the total sizes of all the directories with a total size of at most 100000: %d\n", one(data, 100000))
	fmt.Printf("Part 2: Total size of the smallest directory that, if deleted, would free up enough space on the filesystem to run the update: %d\n", two(data))
}
