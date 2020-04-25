package collections

import (
	"fmt"
	"learning/collections/arrays"
	"learning/collections/slices"
)

// Run collections usecase, contains both arrays and slices example
func Run() {
	fmt.Print("Collections usecases: \n")
	arrays.Run()
	slices.Run()
}
