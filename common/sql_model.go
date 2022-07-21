package common

import "time"

type SQLModel struct {
	Id        int        `json:"id"`
	Status    int        `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
