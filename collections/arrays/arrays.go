package arrays

import (
	"fmt"
)

// Go arrays are values which means its elements will be copied in following example or in case of passing array into function

func Run() {
	// normal init
	var students [3]string
	students[0] = "John Krasinski"
	students[2] = "Emily Krasinski"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("There are %v students in total\n", len(students))

	// initializer list
	grades := [3]int{4, 5, 2}
	fmt.Printf("Grades: %v\n", grades)

	// 2d array
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)

	a := [...]int{1, 2, 3, 4, 5}
	b := a
	c := &b
	b[1] = 99
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
