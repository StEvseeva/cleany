package repository

import (
	"context"
	"database/sql"

	"github.com/StEvseeva/cleany/internal/models"
)

// CleaningOrderRepository defines the interface for cleaning order data operations
type CleaningOrderRepository interface {
	Create(ctx context.Context, order *models.CleaningOrder) error
	GetByID(ctx context.Context, id int) (*models.CleaningOrder, error)
	GetAll(ctx context.Context) ([]models.CleaningOrder, error)
	Update(ctx context.Context, order *models.CleaningOrder) error
	Delete(ctx context.Context, id int) error
	AssignCleaner(ctx context.Context, orderID, cleanerID int) error
	RemoveCleaner(ctx context.Context, orderID, cleanerID int) error
}

// cleaningOrderRepository implements CleaningOrderRepository
type cleaningOrderRepository struct {
	db *sql.DB
}

// NewCleaningOrderRepository creates a new cleaning order repository
func NewCleaningOrderRepository(db *sql.DB) CleaningOrderRepository {
	return &cleaningOrderRepository{db: db}
}

// Create inserts a new cleaning order into the database
func (r *cleaningOrderRepository) Create(ctx context.Context, order *models.CleaningOrder) error {
	query := `
		INSERT INTO cleaning_orders (booking_id, cleaning_ts, cleaning_type, cost, done, notes)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	return r.db.QueryRowContext(ctx, query,
		order.BookingId,
		order.CleaningTs,
		order.CleaningType,
		order.Cost,
		order.Done,
		order.Notes,
	).Scan(&order.Id)
}

// GetByID retrieves a cleaning order by its ID
func (r *cleaningOrderRepository) GetByID(ctx context.Context, id int) (*models.CleaningOrder, error) {
	query := `
		SELECT id, booking_id, cleaning_ts, cleaning_type, cost, done, notes
		FROM cleaning_orders
		WHERE id = $1`

	order := &models.CleaningOrder{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&order.Id,
		&order.BookingId,
		&order.CleaningTs,
		&order.CleaningType,
		&order.Cost,
		&order.Done,
		&order.Notes,
	)

	if err != nil {
		return nil, err
	}

	return order, nil
}

// GetAll retrieves all cleaning orders
func (r *cleaningOrderRepository) GetAll(ctx context.Context) ([]models.CleaningOrder, error) {
	query := `
		SELECT id, booking_id, cleaning_ts, cleaning_type, cost, done, notes
		FROM cleaning_orders
		ORDER BY id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.CleaningOrder
	for rows.Next() {
		var order models.CleaningOrder
		err := rows.Scan(
			&order.Id,
			&order.BookingId,
			&order.CleaningTs,
			&order.CleaningType,
			&order.Cost,
			&order.Done,
			&order.Notes,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

// Update modifies an existing cleaning order
func (r *cleaningOrderRepository) Update(ctx context.Context, order *models.CleaningOrder) error {
	query := `
		UPDATE cleaning_orders
		SET booking_id = $1, cleaning_ts = $2, cleaning_type = $3, cost = $4, done = $5, notes = $6
		WHERE id = $7`

	result, err := r.db.ExecContext(ctx, query,
		order.BookingId,
		order.CleaningTs,
		order.CleaningType,
		order.Cost,
		order.Done,
		order.Notes,
		order.Id,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Delete removes a cleaning order by its ID
func (r *cleaningOrderRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM cleaning_orders WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// AssignCleaner assigns a cleaner to a cleaning order
func (r *cleaningOrderRepository) AssignCleaner(ctx context.Context, orderID, cleanerID int) error {
	query := `
		INSERT INTO cleaner_orders (order_id, cleaner_id)
		VALUES ($1, $2)`

	_, err := r.db.ExecContext(ctx, query, orderID, cleanerID)
	return err
}

// RemoveCleaner removes a cleaner from a cleaning order
func (r *cleaningOrderRepository) RemoveCleaner(ctx context.Context, orderID, cleanerID int) error {
	query := `
		DELETE FROM cleaner_orders
		WHERE order_id = $1 AND cleaner_id = $2`

	result, err := r.db.ExecContext(ctx, query, orderID, cleanerID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
