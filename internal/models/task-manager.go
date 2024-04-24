package models

import (
	models "task-manager/internal/models/taskmanager/file_proccecers"
	"time"
)

type TaskManager struct {
	Tasks []Task
}

func (tm *TaskManager) CreateTasks() {
	csvParser := models.CsvParser{}

	data := csvParser.GetDataFromInternalFile()

	for _, row := range data {
		deadline, _ := time.Parse("2006-01-02 15:04:05.999999999", row[4])
		tm.Tasks = append(tm.Tasks, Task{
			Name:        row[0],
			Description: row[2],
			Duration:    "",
			UserName:    row[1],
			Deadline:    deadline,
			Done:        false,
			FinishedAT:  nil,
		})
	}

}
