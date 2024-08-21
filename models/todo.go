package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Taskname    string    `json:"taskname" validate:"required"`
	Description string    `json:"description"`
	Starttime   time.Time `json:"starttime" validate:"required"`
	Endtime     time.Time `json:"endtime" valdiate:"required"`
	Status      string    `json:"status"`
	Userid      uint      `json:"userid"`
}
