package main

import "fmt"

func computeTable(number int) *[multMax]int {
	var table [multMax]int
	for multiple := 0; multiple < multMax; multiple++ {
		table[multiple] = number * (multiple + 1)
	}

	return &table
}

func printTable(n int, array *[multMax]int) {
	for index, value := range array {
		fmt.Printf("%d x %d = %d\n", n, index+1, value)
	}
}

const (
	countMax = 3
	multMax  = 5
)

func main() {
	calcMap := make(map[int]*[multMax]int)
	for count := 1; count <= countMax; count++ {
		calcMap[count] = computeTable(count)
	}

	for count := 1; count <= countMax; count++ {
		printTable(count+1, calcMap[count])
		fmt.Println()
		if n := count; n == countMax {
			fmt.Printf("Done writing %d tables\n.", n)
		}
	}

}
