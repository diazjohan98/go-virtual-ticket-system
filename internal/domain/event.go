package domain

import "time"

type Event struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	TotalTickets     int       `json:"total_tickets"`
	AvailableTickets int       `json:"available_tickets"`
	CreatedAt        time.Time `json:"created_at"`
}
