package main

import (
	"testing"
)

func input() []byte {
	return []byte("1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000")
}

func TestOne(t *testing.T) {
	wantPos, wantTotal := 4, 24000
	gotPos, gotTotal := one(input())

	if gotPos != wantPos || gotTotal != wantTotal {
		t.Fail()
	}
}

func TestTwo(t *testing.T) {
	wantPositions, wantTotal := []int{4, 3, 5}, 45000
	gotPositions, gotTotal := two(input())

	if gotTotal != wantTotal || len(gotPositions) != 3 {
		t.Fail()
	}
	for i := range gotPositions {
		if gotPositions[i] != wantPositions[i] {
			t.Fail()
		}
	}
}
