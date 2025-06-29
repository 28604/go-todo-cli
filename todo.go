package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

const dataFile = "tasks.json"

// loadTasks() reads the tasks from dataFile and returns a slice of tasks and possible error.
func loadTasks() ([]Task, error) {
	data, err := os.ReadFile(dataFile)

	// If dataFile does not exist, then returns an empty slice and a nil error.
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil // No file yet.
		}
	}

	// If dataFile exists, then unmarshals the JSON data into a slice of Task.
	var tasks []Task
	err = json.Unmarshal(data, &tasks) // json.Unmarshal() decodes JSON data into Go data structures.
	return tasks, err
}

// saveTasks() takes a slice of Task and saves it into the dataFile.
func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ") // json.MarshalIndent() encodes Go data structures into readable JSON data.
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644) // 0644 means owner reads/writes, others read.
}

// $ go run . add "Anything you want to do"
func addTask(args []string) {
	if len(args) == 0 {
		fmt.Println("Please provide a task description.")
		return
	}
	desc := args[0]
	tasks, _ := loadTasks()
	tasks = append(tasks, Task{Description: desc, Done: false})
	saveTasks(tasks)
	fmt.Println("Added: ", desc)
}

// $ go run . list
func listTask() {
	tasks, err := loadTasks()
	if err != nil {
		panic("Error occured while loading tasks.")
	}
	for i, task := range tasks {
		var status string
		if task.Done {
			status = "[x]"
		} else {
			status = "[ ]"
		}
		fmt.Printf("%d. %s %s\n", i+1, status, task.Description)
	}
}

// $ go run . done 1
func markDone(args []string) {

	if len(args) == 0 {
		fmt.Println("Please provide task number.")
		return
	}

	index, err := strconv.Atoi(args[0])
	if err != nil || index < 1 {
		fmt.Println("Invalid task number.")
		return
	}

	tasks, _ := loadTasks()
	if index > len(tasks) {
		fmt.Println("Task number out of range.")
		return
	}

	tasks[index-1].Done = true
	saveTasks(tasks)
	fmt.Println("Marked done:", tasks[index-1].Description)
}
