package model

import "time"

type DbEmployee struct {
	ID           int
	Fullname     string
	Position     string
	Reception_at time.Time
}
