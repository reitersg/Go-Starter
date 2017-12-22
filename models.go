package main

import "time"

type Task struct {
	Name		string
	Assignee 	string
	Due			time.Time
	Description	string
	taskType	string
}

type Tasks Task[]