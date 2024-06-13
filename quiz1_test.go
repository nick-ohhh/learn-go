package main

import (
	"fmt"
	"testing"
)

func TestSolveMeAdd(t *testing.T) {
	sum := SolveMe(3, 6)
	expected := 9

	if sum != uint32(expected) {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleSolveMe() {
	sum := SolveMe(4, 8)
	fmt.Println(sum)
	// Output: 12
}
