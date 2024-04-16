package models

import "time"

type Task struct {
	Name        string
	Description string
	Duration    string
	User        string
	Deadline    time.Time
	Done        bool
	FinishedAT  time.Time
	CreatedAt   time.Time
}

type TaskManager struct {
	Task []Task
}

func (tm TaskManager) addNewTask() {
	// do something
}
