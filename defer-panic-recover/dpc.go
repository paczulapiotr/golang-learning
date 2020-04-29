package dpr

import (
	"fmt"
	"log"
)

// defered functions are executed after func ends and before it returns value
// panic has pretty much functionality of 'throw' keyword
// recover returns value if program 'panicked' earlier, can be used in defered functions too

func Run() {
	fmt.Println("First message")
	defer func() {
		fmt.Println("Defered message")
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			//panic(err)
		}
	}()

	fmt.Println("Middle message")

	panic("Panic message")
}
