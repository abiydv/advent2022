package main

import (
	"log"
	"os"
	"testing"
)

func input() []byte {
	data, err := os.ReadFile("input_test.txt") // read test input from file
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func TestOne(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		maxSize int
		want    int
	}{
		{"example", input(), 100000, 95437},
		{"one", []byte("$ cd /\n$ ls\n50000 file1\n1000 file2\ndir a\ndir b\n$ cd a\n$ ls\n25000 fileina\n$ cd ..\n$ cd b\n$ ls\n25000 fileinb\n"), 100000, 50000},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := one(test.input, test.maxSize)
			if got != test.want {
				t.Errorf("Got %d, but Want %d", got, test.want)
			}
		})
	}
}

func TestTwo(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  int
	}{
		{"example", input(), 24933642},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := two(test.input)
			if got != test.want {
				t.Errorf("Got %d, but Want %d", got, test.want)
			}
		})
	}
}
