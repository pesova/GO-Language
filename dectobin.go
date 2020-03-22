package main

import "fmt"

func main() {
	var Number int

	for {
		fmt.Printf("Enter decimal number\n")
		fmt.Scan(&Number)

		fmt.Printf("%d, in binary format = %b\n", Number, Number)

		fmt.Printf("\nplease press Ctrl + C, to exit loop\n\n")
	}
}
