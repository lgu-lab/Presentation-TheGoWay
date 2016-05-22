package main

import "fmt"

func computeTable(count int, slice []int) {
	for multiple := 1; multiple <= 10; multiple++ {
		slice[multiple] = count * multiple
	}
}

const (
	countMax = 1
	multMax  = 5
)

func main() {
	var array [countMax + 1][multMax + 1]int
	for count := 1; count <= countMax; count++ {
		for multiple := 1; multiple <= multMax; multiple++ {
			array[count][multiple] = count * multiple
		}
	}
	for count := 1; count < len(array); count++ {
		for multiple := 1; multiple < len(array[0]); multiple++ {
			fmt.Printf("%d x %d = %d\n", count, multiple, array[count][multiple])
		}
		fmt.Println()
	}
}
