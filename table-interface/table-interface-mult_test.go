package main

import (
	"fmt"
	"testing"
)

//TestMultCompute will validate the array computing logic
func TestMultCompute(t *testing.T) {
	testArr := [...]int{1, 2, 3, 4, 5}
	number := 1
	size := len(testArr)
	testAgainst := TableValues{number: number, size: size, values: testArr[:]}
	var table *TableValues
	strategy := MultTable{}
	table = computeTable(number, size, strategy)
	if !testAgainst.Equals(table) {
		fmt.Printf("Unequal: %v -> was expecting %v\n", table, testAgainst)
		t.Fail()
	}
}

// ExamplePrint will validate the output for the computed array
func ExampleMultPrint() {
	testArr := [...]int{1, 2, 3, 4, 5}
	number := 1
	size := len(testArr)
	testAgainst := TableValues{number: number, size: size, values: testArr[:]}
	strategy := MultTable{}

	strategy.Print(&testAgainst)
	//Output:
	//1 x 1 = 1
	//1 x 2 = 2
	//1 x 3 = 3
	//1 x 4 = 4
	//1 x 5 = 5
}
