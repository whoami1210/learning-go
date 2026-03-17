package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("=== Name Length Checker ===")
	fmt.Print("What is your name? ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	
	length := len(name)
	
	fmt.Printf("Your name '%s' has %d characters!\n", name, length)
	
	if length < 5 {
		fmt.Println("Short name!")
	} else if length < 10 {
		fmt.Println("Medium name!")
	} else {
		fmt.Println("Long name!")
	}
}
