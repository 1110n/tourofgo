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

func CreateStringArray() {
	countries := [4]string{"India", "Austrilia", "Japan", "US"}
	fmt.Println(countries)
}

func ReverseArrayElems() {}

func AccessSpecificElem() {
	countries := [7]string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"}
	fmt.Println(countries)
	fmt.Println("Printing the 3rd Element of the above Array")
	fmt.Println(countries[2])
}

func EditArrayElem() {
	tasks := [3]string{"LOGIN", "SIGNUP", "DASHBOARD"}
	fmt.Println(tasks)
	fmt.Println("Changing the order of first 2 tasks")
	tasks[0], tasks[1] = tasks[1], tasks[0]
	fmt.Println(tasks)
}

func StartTutorial() {
	fmt.Println("\t\t1. To create a new Int Array")
	fmt.Println("\t\t2. To create a new String Array")
	fmt.Println("\t\t3. To reverse the elements of an Array")
	fmt.Println("\t\t4. To access specific element of an Array")
	fmt.Println("\t\t5. To change the value of an array")
	fmt.Println("\t\t6. To Exit the Tour")
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
		fmt.Println("Closing the Tour!")
		os.Exit(0)
	}
	StartTutorial()
}
