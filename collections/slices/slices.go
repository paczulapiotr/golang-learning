package slices

import (
	"fmt"
)

// Slices are references unlike arrays

func Run() {
	fmt.Println()
	baseOperations()

	fmt.Println()
	slicingCollections()

	fmt.Println()
	makingSlices()

	fmt.Println()
	appending()

	fmt.Println()
	removingElement()
}

func baseOperations() {
	a := []int{1, 2, 3}
	b := a
	c := &b
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func slicingCollections() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := a[:]   // slice of all elements
	c := a[3:]  // slice after 3rd element to the end
	d := a[:6]  // slice from start up to 6th including
	e := a[3:6] // slice elements after 3rd up to 6th element including

	a[5] = 99

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}

func makingSlices() {
	a := make([]int, 3, 10) // type, length, capacity
	showSliceDetails(a)
}

func appending() {
	a := []int{}
	showSliceDetails(a)
	a = append(a, 1)
	showSliceDetails(a)
	a = append(a, 2, 3, 4, 5)
	showSliceDetails(a)
	a = append(a, []int{6, 7, 8, 9}...)
	showSliceDetails(a)
}

func removingElement() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	b := append(a[:2], a[3:]...)
	fmt.Println(a)
	fmt.Println(b)
}

func showSliceDetails(slice []int) {
	fmt.Println(slice)
	fmt.Printf("Length: %v\n", len(slice))
	fmt.Printf("Capacity: %v\n", cap(slice))
}
