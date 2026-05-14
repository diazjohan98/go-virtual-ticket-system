package domain

import "time"

type TicketStatus string

const (
	StatusReserved  TicketStatus = "reserved"
	StatusPurchased TicketStatus = "purchased"
	StatusCancelled TicketStatus = "cancelled"
)

type Ticket struct {
	ID        string       `json:"id"`
	EventID   string       `json:"event_id"`
	UserID    string       `json:"user_id"`
	Status    TicketStatus `json:"status"`
	CreatedAt time.Time    `json:"created_at"`
}
