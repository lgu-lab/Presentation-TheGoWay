package main

import "fmt"

func computeTable(number int, max int, slice []int) {
	for multiple := 1; multiple <= max; multiple++ {
		slice[multiple] = number * multiple
	}
}

func printTable(n int, slice []int) {
	for multiple := 1; multiple < len(slice); multiple++ {
		fmt.Printf("%d x %d = %d\n", n, multiple, slice[multiple])
	}
}

const (
	countMax = 1
	multMax  = 5
)

func main() {
	var array [countMax + 1][multMax + 1]int
	for count := 1; count <= countMax; count++ {
		computeTable(1, multMax, array[count][:])
	}

	for count := 1; count <= countMax; count++ {
		printTable(count, array[count][:multMax-1])
	}

}
