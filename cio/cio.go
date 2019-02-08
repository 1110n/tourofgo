package cio

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInt() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	readLine := strings.TrimSuffix(text, "\n")
	i, err := strconv.Atoi(readLine)
	if err != nil {
		return i
	}
	return i
}
