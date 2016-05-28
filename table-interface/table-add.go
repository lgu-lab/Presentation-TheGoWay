package main

import "fmt"

// AddTable provides the compute capability for Addition tables
type AddTable struct {
}

// Create will create and populate the structure values
func (AddTable) Create(number int, size int) *TableValues {
	t := TableValues{number: number, size: size}
	t.values = make([]int, t.size)
	for multiple := 0; multiple < t.size; multiple++ {
		t.values[multiple] = t.number + (multiple + 1)
	}

	return &t
}

// Print will print the table using the correct operation syntax
func (AddTable) Print(t *TableValues) {
	for index, value := range t.values {
		fmt.Printf("%d + %d = %d\n", t.number, index+1, value)
	}
}
