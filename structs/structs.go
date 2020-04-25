package structs

import (
	"fmt"
	"reflect"
)

// public private case sensitive, also properties
// structs are value types

type doctor struct {
	number     int
	actorName  string
	companions []string
}

type animal struct {
	name   string `required,maxLength:"100"`
	origin string
}

type bird struct {
	animal
	maxSpeed uint16
}

func normal() {
	aDoctor := doctor{
		number:    3,
		actorName: "John Krasinski",
		companions: []string{
			"Anton Sokolov",
			"Maria Sharapova",
			"John Johnson",
		},
	}

	fmt.Println(aDoctor)
}

func embeddedWithReflect() {
	bird := bird{
		animal: animal{
			name:   "Pigeon",
			origin: "Poland",
		},
		maxSpeed: 60,
	}
	fmt.Println(bird)

	t := reflect.TypeOf(bird)
	field, _ := t.FieldByName("name")
	fmt.Println(field.Tag)
}

func annonymous() {
	// anonymous struct
	bDoctor := struct{ name string }{
		name: "Elizabeth Olsen",
	}

	fmt.Println(bDoctor)
}

func Run() {
	normal()
	annonymous()
	embeddedWithReflect()
}
