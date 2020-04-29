package ifstatement

import (
	"fmt"
)

func Run() {
	number := 50
	guess := 70

	if guess < number {
		fmt.Println("Too low")
	} else if guess > number {
		fmt.Println("Too high")
	} else if guess == number {
		fmt.Println("You got it")
	}
}
