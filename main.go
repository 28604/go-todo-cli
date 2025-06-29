package main

import (
	"fmt"
	"os"
)

// Main funcion can run three features: add, list, done.
// $ go run . add "Buy groceries"
// $ go run . list
// $ go run . done 2
func main() {
	args := os.Args // args := []string{".", "add", "Buy groceries"}

	if len(args) < 2 {
		fmt.Println("Invalid input.")
		fmt.Println("* To add tasks: go run . add \"{{task name}}\"")
		fmt.Println("* To list tasks: go run . list")
		fmt.Println("* To mark tasks done go run . done {{task number}}")
		return
	}

	// Switch cases check the second argument and do the task accordingly.
	switch args[1] {
	case "add":
		addTask(args[2:])
	case "list":
		listTask()
	case "done":
		markDone(args[2:])
	default:
		fmt.Println("Unknown command:", args[1])
	}
}
