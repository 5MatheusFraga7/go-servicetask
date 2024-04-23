package models

import "time"

type Task struct {
	Id          int
	Name        string
	Description string
	Duration    string
	UserName    string
	Deadline    time.Time
	Done        bool
	FinishedAT  *time.Time
	CreatedAt   time.Time
}
