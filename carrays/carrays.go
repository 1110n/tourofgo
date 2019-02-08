package carrays

import (
	"fmt"
	"os"

	"my.scripts/tour/cio"
)

func CreateIntArray() {
	fmt.Println("Syntax: type [n]T")
	var i [10]int
	for n := 0; n < 10; n++ {
		i[n] = n * n
	}
	fmt.Println(i)
	fmt.Println("Syntax: cubes := [4]int{1, 8, 27, 64}")
	cubes := [4]int{1, 8, 27, 64}
	fmt.Println(cubes)
}

func CreateStringArray() {}

func ReverseArrayElems() {}

func AccessSpecificElem() {}

func EditArrayElem() {}

func StartTutorial() {
	fmt.Println("\t\t1. To create a new Int Array")
	fmt.Println("\t\t2. To create a new String Array")
	fmt.Println("\t\t3. To reverse the elements of an Array")
	fmt.Println("\t\t4. To access specific element of an Array")
	fmt.Println("\t\t5. To change the value of an array")
	fmt.Println("\t\t6. To Exit the Tutorial")
	command := cio.GetInt()
	if command == 1 {
		CreateIntArray()
	}
	if command == 2 {
		CreateStringArray()
	}
	if command == 3 {
		ReverseArrayElems()
	}
	if command == 4 {
		AccessSpecificElem()
	}
	if command == 5 {
		EditArrayElem()
	}
	if command == 6 {
		fmt.Println("Closing the Tutorial!")
		os.Exit(0)
	}
	StartTutorial()
}
