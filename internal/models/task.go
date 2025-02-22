package model

import "time"

type Task struct {
	ID          int
	Title       string
	Description string
	Content     string
	Created     time.Time
	Author      User
	Completed   bool
	InProgress  bool
}
