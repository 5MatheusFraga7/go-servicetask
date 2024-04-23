package models

import (
	models "task-manager/internal/models/taskmanager/file_proccecers"
	"time"
)

type TaskManager struct {
	Tasks []Task
}

func (tm TaskManager) CreateTasks() {
	csvParser := models.CsvParser{}

	data := csvParser.GetDataFromInternalFile()

	for _, row := range data {
		deadline, _ := time.Parse("2006-01-02 15:04:05", row[1])
		tm.Tasks = append(tm.Tasks, Task{
			Name:        row[0],
			Description: row[0],
			Duration:    row[0],
			UserName:    row[0],
			Deadline:    deadline,
			Done:        false,
			FinishedAT:  nil,
		})
	}
}
