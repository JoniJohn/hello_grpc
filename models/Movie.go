package models

import "time"

type Movie struct {
	ID       string
	Title    string
	Genre    string
	CreateAt time.Time
	UpdateAt time.Time
}
