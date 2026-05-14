package domain

import "time"

type Event struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	TotalTickests    time.Time `json:"Total_tickests"`
	AvailableTickets time.Time `json:"available_tickets"`
	CreatedAt        time.Time `json:"created_at"`
}
