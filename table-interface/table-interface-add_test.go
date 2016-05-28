package main

import (
	"fmt"
	"testing"
)

//TestCompute will validate the array computing logic
func TestAddCompute(t *testing.T) {
	testArr := [...]int{2, 3, 4, 5, 6}
	number := 1
	size := len(testArr)
	testAgainst := TableValues{number: number, size: size, values: testArr[:]}
	var table *TableValues
	var strategy = AddTable{}
	table = computeTable(number, size, strategy)
	if !testAgainst.Equals(table) {
		fmt.Printf("Unequal: %v -> was expecting %v\n", table, testAgainst)
		t.Fail()
	}
}

// ExamplePrint will validate the output for the computed array
func ExampleAddPrint() {
	testArr := [...]int{2, 3, 4, 5, 6}
	number := 1
	size := len(testArr)
	testAgainst := TableValues{number: number, size: size, values: testArr[:]}
	var strategy = AddTable{}

	strategy.Print(&testAgainst)
	//Output:
	//1 + 1 = 2
	//1 + 2 = 3
	//1 + 3 = 4
	//1 + 4 = 5
	//1 + 5 = 6
}
