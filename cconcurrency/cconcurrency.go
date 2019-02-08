package cconcurrency

import (
	"fmt"
	"time"

	"my.scripts/tour/cio"
)

func RandomPrint(str string, rounds int) {
	for i := 0; i < rounds; i++ {
		fmt.Println(str)
		time.Sleep(1000 * time.Millisecond)
	}
}

func PrintWithConcurrency() {
	fmt.Println("Performing Print with Concurrency")
	go RandomPrint("Second Thread", 3)
	go RandomPrint("Third Thread", 3)
	RandomPrint("First Thread", 3)
}

func StartTutorial() {
	fmt.Println("\t\t1. To run printing with threads")
	fmt.Println("\t\t2. To Exit")
	command := cio.GetInt()
	switch command {
	case 1:
		PrintWithConcurrency()
	case 2:
		fmt.Println("Closing the Tutorial!")
	default:
		fmt.Println(" command not recognized ")
	}
}
