package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What's your name? ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error!")
		return
	}

	name = strings.TrimSpace(name)

	fmt.Printf("Hello, %s\n", name)

}
