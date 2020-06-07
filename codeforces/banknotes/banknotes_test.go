package main

import "testing"

func TestCalcMininmalChange(t *testing.T) {
	coins := []int{1, 5, 6, 9}
	amount := 11
	got := calcMininmalChange(coins, amount)
	expect := 2
	if got != expect {
		t.Errorf("Expected %v, got %v", expect, got)
	}

	coins = []int{1, 10, 100}
	amount = 4
	got = calcMininmalChange(coins, amount)
	expect = 4
	if got != expect {
		t.Errorf("Expected %v, got %v", expect, got)
	}

	coins = []int{1, 10, 100}
	amount = 15
	got = calcMininmalChange(coins, amount)
	expect = 6
	if got != expect {
		t.Errorf("Expected %v, got %v", expect, got)
	}

	coins = []int{1, 10, 100}
	amount = 4137
	got = calcMininmalChange(coins, amount)
	expect = 51
	if got != expect {
		t.Errorf("Expected %v, got %v", expect, got)
	}
}
