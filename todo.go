package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var todos []string
	
	fmt.Println("=== Todo List App ===")
	
	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Add task")
		fmt.Println("2. View all tasks")
		fmt.Println("3. Delete task")
		fmt.Println("4. Exit")
		fmt.Print("Choose option: ")
		
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		
		switch choice {
		case "1":
			fmt.Print("Enter task: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			
			if task == "" {
				fmt.Println("Task cannot be empty!")
				continue
			}
			
			todos = append(todos, task)
			fmt.Println("✓ Task added!")
			
		case "2":
			if len(todos) == 0 {
				fmt.Println("No tasks yet!")
			} else {
				fmt.Println("\n--- Your Tasks ---")
				for i, task := range todos {
					fmt.Printf("%d. %s\n", i+1, task)
				}
			}
			
		case "3":
			if len(todos) == 0 {
				fmt.Println("No tasks to delete!")
				continue
			}
			
			fmt.Println("\n--- Your Tasks ---")
			for i, task := range todos {
				fmt.Printf("%d. %s\n", i+1, task)
			}
			
			fmt.Print("Enter task number to delete: ")
			indexStr, _ := reader.ReadString('\n')
			index, err := strconv.Atoi(strings.TrimSpace(indexStr))
			
			if err != nil || index < 1 || index > len(todos) {
				fmt.Println("Invalid task number!")
				continue
			}
			
			todos = append(todos[:index-1], todos[index:]...)
			fmt.Println("✓ Task deleted!")
			
		case "4":
			fmt.Println("Goodbye!")
			return
			
		default:
			fmt.Println("Invalid option! Choose 1-4.")
		}
	}
}
