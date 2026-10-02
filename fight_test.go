package main

import "testing"

func TestRollDropBelowChanceSucceeds(t *testing.T) {
	if !rollDrop(0.5, 0.49) {
		t.Fatal("roll below chance should succeed")
	}
}

func TestRollDropAtOrAboveChanceFails(t *testing.T) {
	if rollDrop(0.5, 0.5) {
		t.Fatal("roll equal to chance should fail")
	}
	if rollDrop(0.5, 0.9) {
		t.Fatal("roll above chance should fail")
	}
}

func TestRollDropZeroChanceNeverSucceeds(t *testing.T) {
	if rollDrop(0, 0) {
		t.Fatal("zero chance should never succeed, even with roll 0")
	}
}

func TestRollDropCertainChanceAlwaysSucceeds(t *testing.T) {
	if !rollDrop(1, 0.999999) {
		t.Fatal("100% chance should succeed for any roll below 1")
	}
}
