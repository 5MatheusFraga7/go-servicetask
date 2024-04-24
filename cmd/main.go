package main

import (
	"fmt"
	"strconv"
	"task-manager/internal/db/repository"
	"task-manager/internal/models"
)

func main() {
	tm := models.TaskManager{}
	taskRepository := repository.TaskRepository{}

	tm.CreateTasks()

	size := len(tm.Tasks)

	fmt.Println(strconv.Itoa(size))

	for _, task := range tm.Tasks {
		fmt.Println(task)
		taskRepository.CreateTask(task)
	}
}
