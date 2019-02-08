package cstructs

import (
	"fmt"
	"os"

	"my.scripts/tour/cio"
)

type User struct {
	name        string
	age         int
	description string
}

func CreateUserStruct() User {
	u := User{"Rishabh", 23, "Rishabh is a programmer!"}
	fmt.Println(u)
	return u
}

func StartTutorial() {
	fmt.Println("\t\t1. To create a user structure")
	fmt.Println("\t\t2. To Exit")
	command := cio.GetInt()
	if command == 1 {
		CreateUserStruct()
	}
	if command == 2 {
		fmt.Println("Closing the Tutorial!")
		os.Exit(0)
	}
	StartTutorial()
}
