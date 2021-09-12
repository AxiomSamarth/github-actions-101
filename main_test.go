package main

import "testing"

func TestSum(t *testing.T) {
	sum := Sum(3, 4)
	if sum != 7 {
		t.Error("The Sum function returned wrong output")
	}
}
