package main

import (
	"fmt"
	"testing"
)

const (
	n    int = 1
	mult int = 3
)

var testAgainst = [...]int{0, 1, 2, 3}

//TestCompute will validate the array computing logic
func TestCompute(t *testing.T) {
	var table [mult + 1]int
	computeTable(n, mult, table[1:])
	if table != testAgainst {
		fmt.Printf("Unequal: %v -> was expecting %v\n", table, testAgainst)
		t.Fail()
	}
}

// ExamplePrint will validate the output for the computed array
func ExamplePrint() {

	printTable(n, testAgainst[1:])
	//Output:
	//1 x 1 = 1
	//1 x 2 = 2
	//1 x 3 = 3
}

func BenchmarkCompute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var table [mult + 1]int
		computeTable(i, mult, table[1:])
	}

}
