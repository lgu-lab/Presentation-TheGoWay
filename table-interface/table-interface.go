package main

import (
	"fmt"
	"strconv"
)

// Computable marks objects that will comply with our requirements
type Computable interface {
	Create(int, int) *TableValues
	Print(*TableValues)
}

// TableValues keeps the computed values from multiplication
type TableValues struct {
	number, size int
	values       []int
}

// Equals will check for the equality of the two TableValues
func (t *TableValues) Equals(t2 *TableValues) bool {
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

func computeTable(number int, size int, computable Computable) *TableValues {
	return computable.Create(number, size)
}

func execute(countMax int, multMax int, computeStrategy Computable) {
	calcMap := make(map[int]*TableValues)
	for count := 1; count <= countMax; count++ {
		calcMap[count] = computeTable(count, multMax, computeStrategy)
	}

	for count := 1; count <= countMax; count++ {
		computeStrategy.Print(calcMap[count])
		fmt.Println()
		if n := count; n == countMax {
			fmt.Printf("Done writing %d tables\n.", n)
		}
	}
}

func main() {
	var input, countStr string
	var strategy Computable

	fmt.Println("Choose Option: A-> Add, M->Multiply")
	fmt.Scanln(&input)

	fmt.Println("Enter how many tables ( <10 ):")
	fmt.Scanln(&countStr)

	switch string([]rune(input)[0]) {
	case "A":
		strategy = AddTable{}
	case "M":
		strategy = MultTable{}
	}

	count, _ := strconv.Atoi(countStr)
	execute(count, 5, strategy)
}
