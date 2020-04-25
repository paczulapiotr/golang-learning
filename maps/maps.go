package maps

import (
	"fmt"
)

// maps are passed by reference

func Run() {
	statePopulations := map[string]int{
		"California": 2342,
		"Texas":      1241,
		"Florida":    3222,
		"New York":   2230,
	}

	emptyMap := make(map[string]int)
	emptyMap["Chicago"] = 23333

	delete(statePopulations, "New York")

	fmt.Println(statePopulations)

	fmt.Println(emptyMap)

	newYorkPop, ok := statePopulations["New York"]

	if ok {
		fmt.Printf("New York population: %v\n", newYorkPop)
	} else {
		fmt.Println("New York population not available")
	}
}
