package switchstatement

import "fmt"

func Run() {
	defaultSwitch()
	betterSwitch()
	fallthroughSwitch()
	typeSwitch()
}

func defaultSwitch() {
	switch 3 {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}
}

func betterSwitch() {
	i := 10

	switch {
	case i <= 10:
		fmt.Println("Less or equal to 10")
	case i <= 20:
		fmt.Println("Less or equal to 20")
	default:
		fmt.Println("another number")
	}
}

func fallthroughSwitch() {
	i := 10

	switch {
	case i <= 10:
		fmt.Println("Less or equal to 10")
	case i <= 20:
		fmt.Println("Less or equal to 20")
		fallthrough
	default:
		fmt.Println("its a big number")
	}
}

func typeSwitch() {
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is a string")
	case string:
		fmt.Println("i is a float64 ")
	default:
		fmt.Println("i is another type")
	}
}
