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
	want := "CMZ"
	got := one(input())
	if got != want {
		t.Fail()
	}
}

func TestTwo(t *testing.T) {
	want := "MCD"
	got := two(input())
	if got != want {
		t.Fail()
	}
}
