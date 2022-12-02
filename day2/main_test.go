package main

import "testing"

func TestOne(t *testing.T) {
	want := 15
	got := one([]byte("A Y\nB X\nC Z"))
	if got != want {
		t.Fail()
	}
}

func TestTwo(t *testing.T) {
	want := 12
	got := two([]byte("A Y\nB X\nC Z"))
	if got != want {
		t.Fail()
	}
}
