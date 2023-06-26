package main

import (
	"testing"
)

func TestOne(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  int
	}{
		{"case 1", []byte("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), 7},
		{"case 2", []byte("bvwbjplbgvbhsrlpgdmjqwftvncz"), 5},
		{"case 3", []byte("nppdvjthqldpwncqszvftbrmjlhg"), 6},
		{"case 4", []byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), 10},
		{"case 5", []byte("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), 11},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := one(test.input)
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
		{"case 1", []byte("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), 19},
		{"case 2", []byte("bvwbjplbgvbhsrlpgdmjqwftvncz"), 23},
		{"case 3", []byte("nppdvjthqldpwncqszvftbrmjlhg"), 23},
		{"case 4", []byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), 29},
		{"case 5", []byte("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), 26},
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