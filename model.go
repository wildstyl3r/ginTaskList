package main

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name        string     `json:"name"`
	Description string     `json:"descr"`
	Started     *time.Time `json:"started"`
	Cancelled   *time.Time `json:"cancelled"`
	Completed   *time.Time `json:"done"`
}
