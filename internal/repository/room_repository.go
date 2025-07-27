package repository

import (
	"context"
	"database/sql"

	"github.com/StEvseeva/cleany/internal/models"
)

// RoomRepository defines the interface for room data operations
type RoomRepository interface {
	Create(ctx context.Context, room *models.Room) error
	GetByID(ctx context.Context, id int) (*models.Room, error)
	GetAll(ctx context.Context) ([]models.Room, error)
	Update(ctx context.Context, room *models.Room) error
	Delete(ctx context.Context, id int) error
}

// roomRepository implements RoomRepository
type roomRepository struct {
	db *sql.DB
}

// NewRoomRepository creates a new room repository
func NewRoomRepository(db *sql.DB) RoomRepository {
	return &roomRepository{db: db}
}

// Create inserts a new room into the database
func (r *roomRepository) Create(ctx context.Context, room *models.Room) error {
	query := `
		INSERT INTO rooms (floor, "desc")
		VALUES ($1, $2)
		RETURNING id`

	return r.db.QueryRowContext(ctx, query,
		room.Floor,
		room.Desc,
	).Scan(&room.Id)
}

// GetByID retrieves a room by its ID
func (r *roomRepository) GetByID(ctx context.Context, id int) (*models.Room, error) {
	query := `
		SELECT id, floor, "desc"
		FROM rooms
		WHERE id = $1`

	room := &models.Room{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&room.Id,
		&room.Floor,
		&room.Desc,
	)

	if err != nil {
		return nil, err
	}

	return room, nil
}

// GetAll retrieves all rooms
func (r *roomRepository) GetAll(ctx context.Context) ([]models.Room, error) {
	query := `
		SELECT id, floor, "desc"
		FROM rooms
		ORDER BY id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []models.Room
	for rows.Next() {
		var room models.Room
		err := rows.Scan(
			&room.Id,
			&room.Floor,
			&room.Desc,
		)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

// Update modifies an existing room
func (r *roomRepository) Update(ctx context.Context, room *models.Room) error {
	query := `
		UPDATE rooms
		SET floor = $1, "desc" = $2
		WHERE id = $3`

	result, err := r.db.ExecContext(ctx, query,
		room.Floor,
		room.Desc,
		room.Id,
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

// Delete removes a room by its ID
func (r *roomRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM rooms WHERE id = $1`

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
