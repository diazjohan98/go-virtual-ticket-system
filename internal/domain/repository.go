package domain

import "context"

type EventRepository interface {
	GetBYID(ctx context.Context, id string) (*Event, error)
	DecrementAvailableTickets(ctx context.Context, id string) error
}

type TicketRepository interface {
	Enqueue(ctx context.Context, eventID string, userID string) error
	Dequeue(ctx context.Context, eventID string) (string, error)
	GetPosition(ctx context.Context, eventID string, userID string) (int, error)
}
