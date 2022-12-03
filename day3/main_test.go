package main

import "testing"

func input() []byte {
	return []byte("vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw")
}

func TestOne(t *testing.T) {
	want := 157
	got := one(input())
	if got != want {
		t.Fail()
	}
}

func TestTwo(t *testing.T) {
	want := 70
	got := two(input())
	if got != want {
		t.Fail()
	}
}
