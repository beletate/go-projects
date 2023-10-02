package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Task struct {
	ID int
	Description string
	Done bool
}

func addTask(tasks []Task, description string)[]Task {
	newTask := Task {
		ID: len(tasks) + 1,
		Description: description,
		Done: false, 
	}

	return append(tasks, newTask)
}

func main() {
	var tasks []Task

	for {
		fmt.Println("Please, choose an option:")
		fmt.Println("1. Add Task")
		fmt.Println("2. Complete a task")
		fmt.Println("3. Show Tasks")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter you option:")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Create a description for your task:")
			var description string
			fmt.Scanln(&description)
			tasks = addTask(tasks, description)
			fmt.Println("Task added with success!")
		case 2:
			var allDone bool = true
			for _, task := range tasks {
				status := "To do"
				if !task.Done {
					fmt.Printf("%d. %s (%s)\n", task.ID, task.Description, status)
					allDone = false
				}
			}
			if allDone {
				fmt.Println("All tasks marked as completed.")
				break
			}
			fmt.Print("Enter the task id:")
			var id int
			fmt.Scanln(&id)
			for i := range tasks {
				if tasks[i].ID == id{
					tasks[i].Done = true
					fmt.Printf("Task '%s' marked as completed!\n", tasks[i].Description)
				}
			}
		case 3:
			fmt.Println("List")
			for _, task := range tasks {
				status := "To do"
				if task.Done {
					status = "Done"
				}
				fmt.Printf("%d. %s (%s)\n", task.ID, task.Description, status)
			}
		case 4:
			fmt.Println("Exiting")
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}