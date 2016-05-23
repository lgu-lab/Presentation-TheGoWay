package main

import "fmt"

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
	for count, countArray := range array {
		for multiple, val := range countArray {
			fmt.Printf("%d x %d = %d\n", count, multiple, val)
		}
		fmt.Println()
		if n := count; n == countMax {
			fmt.Printf("Done writing %d tables\n.", n)
		}
	}
}
