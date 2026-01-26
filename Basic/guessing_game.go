package basic

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	//Generate a random number
	target := random.Intn(10) + 1

	fmt.Println("Wellcome to the Goessing Game!")
	fmt.Println("I have chosen a number between 1 and 10")

	var guess int

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("your guessing is correct")
			break
		} else if guess > target {
			fmt.Println("your guess is to high")
		} else {
			fmt.Println("your guess is to low")
		}
	}
}
