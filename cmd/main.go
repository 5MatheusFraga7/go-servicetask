package main

import (
	"task-manager/internal/db/repository"
	"task-manager/internal/models"
	"time"
)

func main() {

	task := models.Task{
		Name:        "Primeira tarefa",
		UserName:    "Usuário 1",
		Description: "Descrição da tarefa",
		Deadline:    time.Now().Add(100),
	}

	err := repository.TaskRepository{}.CreateTask(task)

	if err != nil {
		panic(err)
	}

}
