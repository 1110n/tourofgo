package cslices

import (
	"fmt"
	"os"

	"my.scripts/tour/cio"
)

func CreateSlice() {
	fmt.Println("A slice unlike an array is dynamically sized")
	fmt.Println("Syntax: type []T")
	days := [7]string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"}
	fmt.Println(days)
	fmt.Println("Creating slice from above array using: var a []string = days[1:4] ")
	var a []string = days[1:4]
	fmt.Println(a)
}

func CreateMultipleSlices() {
	fmt.Println("Creating Slices")
	var s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(s)
	fmt.Println("Length:", len(s))
	fmt.Println("Capacity:", cap(s))
	fmt.Print("Slicing: ")
	fmt.Println(s[3:5])
	fmt.Println("The Zero value of a slice is nil")
}

func SlicesWithMake() {
	b := make([]int, 0, 5)
	fmt.Println(b)
}

func AppendToSlice() {
	b := make([]int, 0, 5)
	fmt.Println(b)
	fmt.Printf("Before Append: length=%d, capacity=%d\n", len(b), cap(b))
	fmt.Println("Appending to the Slice")
	newb := append(b, 1, 2, 3, 4, 5, 6, 7)
	fmt.Printf("After Append: length=%d, capcaity=%d \n", len(newb), cap(newb))
}

func IterateWithRange() {
	holidays := []string{"Republic Day", "Independence Day", "Gandhi Jayanti"}
	for i, v := range holidays {
		fmt.Println(i)
		fmt.Println(v)
	}
}

func StartTutorial() {
	fmt.Println("\t\t1. To create a new slice")
	fmt.Println("\t\t2. To create multiple slices")
	fmt.Println("\t\t3. To create Slies with builtin Make function")
	fmt.Println("\t\t4. To perform Append Function")
	fmt.Println("\t\t5. To perform Iteration via Range")
	fmt.Println("\t\t0. To Exit")
	command := cio.GetInt()
	switch command {
	case 1:
		CreateSlice()
	case 2:
		CreateMultipleSlices()
	case 3:
		SlicesWithMake()
	case 4:
		AppendToSlice()
	case 5:
		IterateWithRange()
	case 0:
		os.Exit(0)
	default:
		fmt.Println(" command not recognized ")
	}
	StartTutorial()
}
