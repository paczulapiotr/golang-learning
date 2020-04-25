package iotauc

import (
	"fmt"
)

// iota enum
const (
	isAdmin = byte(1 << iota)
	isEmployee
	isCustomer

	itDepartment
	nonItDepartment

	developer
	qa
	leader
)

// Run iota usecase
func Run() {
	const userOne = byte(isAdmin | itDepartment | developer)
	const userTwo = byte(isEmployee | itDepartment | developer | qa | leader)
	const userThree = byte(isCustomer)

	fmt.Printf("User one roles: %.8b \n", userOne)
	fmt.Printf("User one roles: %.8b \n", userOne)
	fmt.Printf("User two roles: %.8b \n", userTwo)
	fmt.Printf("User three roles: %.8b \n", userThree)
	checkIfAdmin("User one", userOne)
	checkIfAdmin("User two", userTwo)
	checkIfAdmin("User three", userThree)
}

func checkIfAdmin(username string, roles byte) {
	if isAdmin&roles == isAdmin {
		fmt.Printf("%s is an admin\n", username)
	} else {
		fmt.Printf("%s is not an admin\n", username)
	}
}
