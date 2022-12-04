package main

import "testing"

func input() []byte {
	return []byte("2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8")
}

func TestOne(t *testing.T) {
	want := 2
	got := one(input())
	if got != want {
		t.Fail()
	}
}

func TestTwo(t *testing.T) {
	want := 4
	got := two(input())
	if got != want {
		t.Fail()
	}
}
