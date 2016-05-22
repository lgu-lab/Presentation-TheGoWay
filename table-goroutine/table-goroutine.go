package main

import "fmt"

func computeTable(number int, max int, slice []int) {
	for multiple := 1; multiple <= max; multiple++ {
		slice[multiple] = number * multiple
	}
}

func printTable(n int, slice []int) {
	for multiple, value := range slice {
		fmt.Printf("%d x %d = %d\n", n, multiple+1, value)
	}
}

const (
	countMax = 3
	multMax  = 15
)

func main() {
	var array [countMax + 1][multMax + 1]int
	for count := 1; count <= countMax; count++ {
		go computeTable(count, multMax, array[count][:])
	}

	for count := 1; count <= countMax; count++ {
		printTable(count, array[count][1:])
	}

}
