package main

import "fmt"

// MultTableValues keeps the computed values from multiplication
type MultTableValues struct {
	number, size int
	values       []int
}

// Equals will check for the equality of the two MultTableValues
func (t *MultTableValues) Equals(t2 *MultTableValues) bool {
	if t == nil || t2 == nil || t.number != t2.number || t.size != t2.size || len(t.values) != len(t2.values) {
		return false
	}
	// now check the contents
	for i, val := range t.values {
		if val != t2.values[i] {
			return false
		}
	}
	return true
}

// Init will create and populate the structure values
func (t *MultTableValues) Init() {
	t.values = make([]int, t.size)
	for multiple := 0; multiple < t.size; multiple++ {
		t.values[multiple] = t.number * (multiple + 1)
	}
}

func computeTable(number int, size int) *MultTableValues {
	table := MultTableValues{number: number, size: size}
	table.Init()

	return &table
}

func printTable(table *MultTableValues) {
	for index, value := range table.values {
		fmt.Printf("%d x %d = %d\n", table.number, index+1, value)
	}
}

const (
	countMax = 3
	multMax  = 5
)

func main() {
	calcMap := make(map[int]*MultTableValues)
	for count := 1; count <= countMax; count++ {
		calcMap[count] = computeTable(count, multMax)
	}

	for count := 1; count <= countMax; count++ {
		printTable(calcMap[count])
		fmt.Println()
		if n := count; n == countMax {
			fmt.Printf("Done writing %d tables\n.", n)
		}
	}

}
