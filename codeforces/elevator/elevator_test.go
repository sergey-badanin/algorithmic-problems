package main

import (
	"fmt"
	"testing"
)

func TestFindMaxSum(t *testing.T) {
	given := []int{240, 216, 287}

	taken := make(map[int]bool)
	answ := make(map[int]bool)
	max := 0

	findMaxSum(given, 502, taken, 0, &max, answ)

	fmt.Println(max)
	fmt.Println(answ)
}

func TestFillElevator(t *testing.T) {
	given := []int{240, 216, 287}
	limit := 400
	got := fillElevator(given, limit)
	expect := 3
	if got != expect {
		t.Errorf("Expected: %v, got:%v", expect, got)
	}

	given = []int{176, 164, 226}
	limit = 350
	got = fillElevator(given, limit)
	expect = 2
	if got != expect {
		t.Errorf("Expected: %v, got:%v", expect, got)
	}

	given = []int{115, 130, 126}
	limit = 600
	got = fillElevator(given, limit)
	expect = 1
	if got != expect {
		t.Errorf("Expected: %v, got:%v", expect, got)
	}

	given = []int{100, 100, 100}
	limit = 100
	got = fillElevator(given, limit)
	expect = 3
	if got != expect {
		t.Errorf("Expected: %v, got:%v", expect, got)
	}
}
