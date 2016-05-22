package main

import "fmt"

func main() {
	var array [11][11]int
	for count := 1; count <= 10; count++ {
		for multiple := 1; multiple <= 10; multiple++ {
			array[count][multiple] = count * multiple
		}
	}
	for count, countArray := range array {
		for multiple, val := range countArray {
			fmt.Printf("%d x %d = %d\n", count, multiple, val)
		}
		fmt.Println()
	}
}
