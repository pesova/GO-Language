package main

import "fmt"
import "time"
import "math/rand"

func main() {
	fmt.Printf("guess the number\n")
	
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(10)

	var guess int 
	
	for {
		fmt.Printf("please input your guess: ")
		fmt.Scan(&guess)
		if guess > secretNumber {
			fmt.Printf("too Big\n")
		} else if guess < secretNumber {
			fmt.Printf("too small\n")
		} else{
			fmt.Printf("You win! ")
			break
		}
	}
}