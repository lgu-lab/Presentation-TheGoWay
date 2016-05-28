package main

import "fmt"

// MultTable provides the compute capability for Multiplication tables
type MultTable struct {
}

// Create will create and populate the structure values
func (MultTable) Create(number int, size int) *TableValues {
	t := TableValues{number: number, size: size}
	t.values = make([]int, t.size)
	for multiple := 0; multiple < t.size; multiple++ {
		t.values[multiple] = t.number * (multiple + 1)
	}

	return &t
}

// Print will print the table using the correct operation syntax
func (MultTable) Print(t *TableValues) {
	for index, value := range t.values {
		fmt.Printf("%d x %d = %d\n", t.number, index+1, value)
	}
}
