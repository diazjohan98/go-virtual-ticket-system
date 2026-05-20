package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/diazjohan98/go-virtual-queue-system/internal/domain"
)

type mysqlEventRepository struct {
	db *sql.DB
}

func NewMysqlEventRepository(db *sql.DB) domain.EventRepository {
	return &mysqlEventRepository{db: db}
}

func (r *mysqlEventRepository) GetByID(ctx context.Context, id string) (*domain.Event, error) {
	query := `SELECT id, name, total_tickets, available_tickets, created_at FROM events WHERE id = ?`

	row := r.db.QueryRowContext(ctx, query, id)

	var event domain.Event

	err := row.Scan(&event.ID, &event.Name, &event.TotalTickets, &event.AvailableTickets, &event.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &event, nil
}

func (r *mysqlEventRepository) DecrementAvailableTickets(ctx context.Context, id string) error {
	query := `UPDATE events SET available_tickets = available_tickets - 1 WHERE id = ? AND available_tickets > 0`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no se pudo actualizar el inventario: tickets agotados o evento no existe")
	}
	return nil
}
