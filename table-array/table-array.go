package main

import "fmt"

func main() {
	var array [11][11]int
	for count := 1; count <= 10; count++ {
		for multiple := 1; multiple <= 10; multiple++ {
			array[count][multiple] = count * multiple
		}

		for count := 1; count < len(array); count++ {
			for multiple := 1; multiple < len(array[0]); multiple++ {
				fmt.Printf("%d x %d = %d\n", count, multiple, array[count][multiple])
			}
			fmt.Println()
		}

	}
}
