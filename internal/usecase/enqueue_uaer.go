package usecase

import (
	"context"
	"errors"

	"github.com/diazjohan98/go-virtual-queue-system/internal/domain"
)

type EnqueueUserUseCase struct {
	eventRepo domain.EventRepository
	queueRepo domain.QueueRepository
}

func NewEnqueueUserUseCase(er domain.EventRepository, qr domain.QueueRepository) *EnqueueUserUseCase {
	return &EnqueueUserUseCase{
		eventRepo: er,
		queueRepo: qr,
	}
}

func (uc *EnqueueUserUseCase) Execute(ctx context.Context, eventID string, userID string) error {
	// Check if event exists
	event, err := uc.eventRepo.GetBYID(ctx, eventID)
	if err != nil {
		return err
	}

	if event == nil {
		return errors.New("event not found")
	}

	if event.AvailableTickets <= 0 {
		return errors.New("sold out: ya no quedan entradas disponibles")
	}

	err = uc.queueRepo.Enqueue(ctx, eventID, userID)
	if err != nil {
		return err
	}

	err = uc.queueRepo.Enqueue(ctx, eventID, userID)
	if err != nil {
		return err
	}

	return nil
}
