package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("=== Vowel Counter ===")
	fmt.Print("Enter a word: ")
	word, _ := reader.ReadString('\n')
	word = strings.ToLower(strings.TrimSpace(word))
	
	vowels := 0
	for _, char := range word {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			vowels++
		}
	}
	
      fmt.Printf("The word '%s' has %d vowels!\n", word, vowels)
}
