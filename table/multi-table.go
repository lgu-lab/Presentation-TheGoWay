package main

import "fmt"

const (
	countMax = 10
	multMax  = 10
)

func main() {
	for count := 1; count <= countMax; count++ {
		for multiple := 1; multiple <= multMax; multiple++ {
			fmt.Printf("%d x %d = %d\n", count, multiple, count*multiple)
		}
		fmt.Println()
		if n := count; n == countMax {
			fmt.Printf("Done writing %d tables\n.", n)
		}
	}
}
