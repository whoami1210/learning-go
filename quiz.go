package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	score := 0
	
	// here is question 1
	fmt.Println("=== Welcome to the Go Quiz Game! ===")
	fmt.Print("Question 1: What is the capital of France? ")
	answer1, _ := reader.ReadString('\n')
	answer1 = strings.TrimSpace(answer1)
	
	if strings.ToLower(answer1) == "paris" {
		fmt.Println("✓ Correct!")
		score++
	} else {
		fmt.Println("✗ Wrong! The answer is Paris.")
	}
	
	// here is question 2
	fmt.Print("\nQuestion 2: What is 124 + 459? ")
	answer2, _ := reader.ReadString('\n')
	answer2 = strings.TrimSpace(answer2)
	
	if answer2 == "583" {
		fmt.Println("✓ Correct!")
		score++
	} else {
		fmt.Println("✗ Wrong! The answer is 583.")
	}
	
	// here is question 3
	fmt.Print("\nQuestion 3: What is the largest planet in our solar system? ")
	answer3, _ := reader.ReadString('\n')
	answer3 = strings.TrimSpace(answer3)
	
	if strings.ToLower(answer3) == "jupiter" {
		fmt.Println("✓ Correct!")
		score++
	} else {
		fmt.Println("✗ Wrong! The answer is Jupiter.")
	}
	
	// final score
	fmt.Printf("\n=== Your Score: %d/3 ===\n", score)
	
	if score == 3 {
		fmt.Println("Wow! You're amazing!" )
	} else if score == 2 {
		fmt.Println("Great job!" )
	} else if score == 1 {
		fmt.Println("Not bad! But you need to practice ")
	} else {
		fmt.Println("Keep learning! ")
	}
}
