package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What's your name? ")
	name, err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	name = strings.TrimSpace(name)

	fmt.Printf("Hello, %s\n", name)
}
