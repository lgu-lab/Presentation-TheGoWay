package main

import (
	"fmt"
	"testing"
)

const (
	n int = 1
)

var testAgainst = [...]int{1, 2, 3, 4, 5}

//TestCompute will validate the array computing logic
func TestCompute(t *testing.T) {
	var table *[multMax]int
	table = computeTable(n)
	if *table != testAgainst {
		fmt.Printf("Unequal: %v -> was expecting %v\n", table, testAgainst)
		t.Fail()
	}
}

// ExamplePrint will validate the output for the computed array
func ExamplePrint() {

	printTable(n, &testAgainst)
	//Output:
	//1 x 1 = 1
	//1 x 2 = 2
	//1 x 3 = 3
	//1 x 4 = 4
	//1 x 5 = 5
}
