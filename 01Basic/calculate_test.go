package main

import "testing"

func testSumWithOne(t *testing.T) {
	if sumWithOne(1) != 2 {
		t.Error("Expected 1 + 1 = 2")
	}
}
