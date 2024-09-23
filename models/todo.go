package models

import (
	"time"
)

type (
	TodoModel struct {
		Title     string
		Status    bool
		CreatedAt time.Time
	}

	Todo struct {
		Title  string `json:"title"`
		TodoId int    `json:"id"`
		Status bool   `json:"status"`
	}
)
