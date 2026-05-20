package worker

import (
	"context"
	"log"
	"time"

	"github.com/tu-usuario-github/go-virtual-queue-system/internal/domain"
)

type QueueWorker struct {
	queueRepo domain.QueueRepository
	eventRepo domain.EventRepository
}

func NewQueueWorker(qr domain.QueueRepository, er domain.EventRepository) *QueueWorker {
	return &QueueWorker{
		queueRepo: qr,
		eventRepo: er,
	}
}

func (w *QueueWorker) Start(ctx context.Context, eventID string) {
	log.Printf("Worker iniciado y vigilando la fila del evento: %s", eventID)

	for {
		select {
		case <-ctx.Done():
			log.Printf("Worker detenido para el evento: %s", eventID)
			return

		default:
			userID, err := w.queueRepo.Dequeue(ctx, eventID)
			if err != nil {
				log.Printf("Error al leer la fila de Redis: %c", err)
				time.Sleep(1 * time.Second)
				continue
			}
			if userID == "" {
				time.Sleep(1 * time.Second)
				continue
			}

			log.Printf("Fila avanzando: Procesando compra para el usuario %s:", &userID)

			err = w.eventRepo.DecrementAvailableTickets(ctx, eventID)
			if err != nil {
				log.Printf("Falló la compra del usuario %s: %v", userID, err)
				continue
			}

			log.Printf("Compra exitosa confirmada en MySQL para el usuario %s:", &userID)
		}
	}
}
