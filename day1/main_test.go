package main

import (
	"bufio"
	"strings"
	"testing"
)

func equalNumSlices(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestFindMax(t *testing.T) {
	want := 3
	got := findMax([]int{1, 2, 3})
	if want != got {
		t.Fail()
	}
}

func TestFindMaxEmptySlice(t *testing.T) {
	want := 0
	got := findMax([]int{})
	if want != got {
		t.Fail()
	}
}

func TestFSumSliceSuccess(t *testing.T) {
	want := []int{6}
	got := fSumSlice(bufio.NewScanner(strings.NewReader("1\n2\n3\n")))
	if !equalNumSlices(want, got) {
		t.Fail()
	}
}
