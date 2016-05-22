package main

import "fmt"

func main() {
	for count := 1; count <= 10; count++ {
		for multiple := 1; multiple <= 10; multiple++ {
			fmt.Printf("%d x %d = %d\n", count, multiple, count*multiple)
		}
		fmt.Println()
	}
}
