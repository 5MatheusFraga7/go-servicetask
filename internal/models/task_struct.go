package models

import "time"

type Task struct {
	Name        string
	Description string
	Duration    string
	UserName    string
	Deadline    time.Time
	Done        bool
	FinishedAT  time.Time
	CreatedAt   time.Time
}
