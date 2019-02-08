package engine

import (
	"fmt"

	"my.scripts/tour/carrays"
	"my.scripts/tour/cconcurrency"
	"my.scripts/tour/cio"
	"my.scripts/tour/cstructs"
)

func Route(cmd int) {
	switch cmd {
	case 1:
		cconcurrency.StartTutorial()
	case 2:
		cstructs.StartTutorial()
	case 3:
		carrays.StartTutorial()
	default:
		fmt.Println("Not a known command.", cmd)
	}
}

func Initialize() {
	fmt.Println("Choose a Tutorial from the Tour!")
	fmt.Println("\t1. Concurrency")
	fmt.Println("\t2. Structs")
	fmt.Println("\t3. Arrays")
	command := cio.GetInt()
	Route(command)
}
