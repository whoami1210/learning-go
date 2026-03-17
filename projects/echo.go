package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Type something! ")
	message, _ := reader.ReadString('\n')
	message = strings.TrimSpace(message)
	
	fmt.Printf("You type: %s\n", message)
}
