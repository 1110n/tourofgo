package cmaps

import (
	"fmt"
	"os"

	"my.scripts/tour/cio"
)

type User struct {
	Name       string
	Age        int
	Profession string
}

func CreateMap() {
	var m map[int]User

}

func StartTutorial() {
	fmt.Println("\t\t1. To create a new map")
	fmt.Println("\t\t2. To create multiple slices")
	fmt.Println("\t\t3. To create Slies with builtin Make function")
	fmt.Println("\t\t4. To perform Append Function")
	fmt.Println("\t\t5. To perform Iteration via Range")
	fmt.Println("\t\t0. To Exit")
	command := cio.GetInt()
	switch command {
	case 1:
		CreateMap()
	case 0:
		os.Exit(0)
	default:
		fmt.Println(" command not recognized ")
	}
	StartTutorial()
}
