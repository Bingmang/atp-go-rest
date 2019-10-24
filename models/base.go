package models

import "time"

type base struct {
	Id int64				`json:"id"`
	CreatedTime time.Time	`json:"created_time"`
	Yn bool					`json:"yn"`
}