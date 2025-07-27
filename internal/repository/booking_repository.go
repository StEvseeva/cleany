package repository

import (
	"context"
	"database/sql"

	"github.com/StEvseeva/cleany/internal/models"
)

// BookingRepository defines the interface for booking data operations
type BookingRepository interface {
	Create(ctx context.Context, booking *models.Booking) error
	GetByID(ctx context.Context, id int) (*models.Booking, error)
	GetAll(ctx context.Context) ([]models.Booking, error)
	Update(ctx context.Context, booking *models.Booking) error
	Delete(ctx context.Context, id int) error
}

// bookingRepository implements BookingRepository
type bookingRepository struct {
	db *sql.DB
}

// NewBookingRepository creates a new booking repository
func NewBookingRepository(db *sql.DB) BookingRepository {
	return &bookingRepository{db: db}
}

// Create inserts a new booking into the database
func (r *bookingRepository) Create(ctx context.Context, booking *models.Booking) error {
	query := `
		INSERT INTO bookings (room_id, check_in_ts, check_out_ts, guests)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	return r.db.QueryRowContext(ctx, query,
		booking.RoomId,
		booking.CheckInTs,
		booking.CheckOutTs,
		booking.Guests,
	).Scan(&booking.Id)
}

// GetByID retrieves a booking by its ID
func (r *bookingRepository) GetByID(ctx context.Context, id int) (*models.Booking, error) {
	query := `
		SELECT id, room_id, check_in_ts, check_out_ts, guests
		FROM bookings
		WHERE id = $1`

	booking := &models.Booking{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&booking.Id,
		&booking.RoomId,
		&booking.CheckInTs,
		&booking.CheckOutTs,
		&booking.Guests,
	)

	if err != nil {
		return nil, err
	}

	return booking, nil
}

// GetAll retrieves all bookings
func (r *bookingRepository) GetAll(ctx context.Context) ([]models.Booking, error) {
	query := `
		SELECT id, room_id, check_in_ts, check_out_ts, guests
		FROM bookings
		ORDER BY id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var booking models.Booking
		err := rows.Scan(
			&booking.Id,
			&booking.RoomId,
			&booking.CheckInTs,
			&booking.CheckOutTs,
			&booking.Guests,
		)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

// Update modifies an existing booking
func (r *bookingRepository) Update(ctx context.Context, booking *models.Booking) error {
	query := `
		UPDATE bookings
		SET room_id = $1, check_in_ts = $2, check_out_ts = $3, guests = $4
		WHERE id = $5`

	result, err := r.db.ExecContext(ctx, query,
		booking.RoomId,
		booking.CheckInTs,
		booking.CheckOutTs,
		booking.Guests,
		booking.Id,
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

// Delete removes a booking by its ID
func (r *bookingRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM bookings WHERE id = $1`

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
