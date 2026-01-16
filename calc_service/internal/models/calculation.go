package models

import "time"

type Calculation struct {
	Operation string
	Operands  []int
	Result    int
	RequestID string
	UserID    string
	CreatedAt time.Time
}
