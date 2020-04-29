package loops

import (
	"fmt"
)

func Run() {
	normalFor()
	extendedFor()
	label()
	rangeFor()
}

func normalFor() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func extendedFor() {
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i)
		fmt.Println(j)
	}
}

func whileFor() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func label() {
MyLoop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i+j > 15 {
				break MyLoop
			}
		}
	}
}

func rangeFor() {
	s := []int{1, 2, 3}
	for key, value := range s {
		fmt.Printf("Key: %v. Value: %v", key, value)
	}
}
