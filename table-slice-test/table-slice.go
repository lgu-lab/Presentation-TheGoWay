package main

import "fmt"

func computeTable(number int, max int, slice []int) {
	for multiple := 0; multiple < max; multiple++ {
		slice[multiple] = number * (multiple + 1)
	}
}

func printTable(n int, slice []int) {
	for multiple, value := range slice {
		fmt.Printf("%d x %d = %d\n", n, multiple+1, value)
	}
}

const (
	countMax = 1
	multMax  = 5
)

func main() {
	var array [countMax][multMax]int
	for count := 0; count < countMax; count++ {
		computeTable(count+1, multMax, array[count][:])
	}

	for count := 0; count < countMax; count++ {
		printTable(count+1, array[count][:])
		fmt.Println()
		if n := count + 1; n == countMax {
			fmt.Printf("Done writing %d tables\n.", n)
		}
	}

}
