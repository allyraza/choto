package models

import "time"

// URL :
type URL struct {
	ID        int       `json:"id"`
	Short     string    `json:"short_url"`
	Long      string    `json:"long_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
