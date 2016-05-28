package main

import (
	"fmt"
	"testing"
)

const (
	number int = 1
	size   int = 5
)

//TestCompute will validate the array computing logic
func TestCompute(t *testing.T) {
	var testArr = [...]int{1, 2, 3, 4, 5}
	var testAgainst = MultTableValues{number: number, size: size, values: testArr[:]}
	var table *MultTableValues
	table = computeTable(number, size)
	if !testAgainst.Equals(table) {
		fmt.Printf("Unequal: %v -> was expecting %v\n", table, testAgainst)
		t.Fail()
	}
}

//TestCompute will validate the array computing logic
func TestSize(t *testing.T) {
	var testArr = [...]int{1, 2, 3, 4}
	var testAgainst = MultTableValues{number: number, size: size, values: testArr[:]}

	var table *MultTableValues
	table = computeTable(number, size)
	if testAgainst.Equals(table) {
		fmt.Printf("Size check failed: %v -> was expecting %v\n", table, testAgainst)
		t.Fail()
	}
}

// ExamplePrint will validate the output for the computed array
func ExamplePrint() {
	var testArr = [...]int{1, 2, 3, 4, 5}
	var testAgainst = MultTableValues{number: number, size: size, values: testArr[:]}

	printTable(&testAgainst)
	//Output:
	//1 x 1 = 1
	//1 x 2 = 2
	//1 x 3 = 3
	//1 x 4 = 4
	//1 x 5 = 5
}
