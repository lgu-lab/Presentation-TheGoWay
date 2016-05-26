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
	for multiple, value := range array {
		fmt.Printf("%d x %d = %d\n", n, multiple+1, value)
	}
}

const (
	countMax = 3
	multMax  = 5
)

func main() {
	calcMap := make(map[int]*[multMax]int)
	for count := 0; count < countMax; count++ {
		calcMap[count] = computeTable(count + 1)
	}

	for count := 0; count < countMax; count++ {
		printTable(count+1, calcMap[count])
		fmt.Println()
		if n := count + 1; n == countMax {
			fmt.Printf("Done writing %d tables\n.", n)
		}
	}

}
