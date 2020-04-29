package pointers

import (
	"fmt"
)

type bookshelf struct {
}

func (bs bookshelf) insert(book string) {
	fmt.Println("Inserted book: ", book)
}

func Run() {
	a := 10
	b := &a

	*b = 20

	fmt.Println(a)
	fmt.Println(b)
	bs := bookshelf{}
	bsPointer := &bs

	bs.insert("Value type book")
	bsPointer.insert("Pointer book")

}
