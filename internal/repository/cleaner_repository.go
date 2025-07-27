package repository

import (
	"context"
	"database/sql"

	"github.com/StEvseeva/cleany/internal/models"
)

// CleanerRepository defines the interface for cleaner data operations
type CleanerRepository interface {
	Create(ctx context.Context, cleaner *models.Cleaner) error
	GetByID(ctx context.Context, id int) (*models.Cleaner, error)
	GetAll(ctx context.Context) ([]models.Cleaner, error)
	Update(ctx context.Context, cleaner *models.Cleaner) error
	Delete(ctx context.Context, id int) error
}

// cleanerRepository implements CleanerRepository
type cleanerRepository struct {
	db *sql.DB
}

// NewCleanerRepository creates a new cleaner repository
func NewCleanerRepository(db *sql.DB) CleanerRepository {
	return &cleanerRepository{db: db}
}

// Create inserts a new cleaner into the database
func (r *cleanerRepository) Create(ctx context.Context, cleaner *models.Cleaner) error {
	query := `
		INSERT INTO cleaners (name, surname)
		VALUES ($1, $2)
		RETURNING id`

	return r.db.QueryRowContext(ctx, query,
		cleaner.Name,
		cleaner.Surname,
	).Scan(&cleaner.Id)
}

// GetByID retrieves a cleaner by its ID
func (r *cleanerRepository) GetByID(ctx context.Context, id int) (*models.Cleaner, error) {
	query := `
		SELECT id, name, surname
		FROM cleaners
		WHERE id = $1`

	cleaner := &models.Cleaner{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&cleaner.Id,
		&cleaner.Name,
		&cleaner.Surname,
	)

	if err != nil {
		return nil, err
	}

	return cleaner, nil
}

// GetAll retrieves all cleaners
func (r *cleanerRepository) GetAll(ctx context.Context) ([]models.Cleaner, error) {
	query := `
		SELECT id, name, surname
		FROM cleaners
		ORDER BY id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cleaners []models.Cleaner
	for rows.Next() {
		var cleaner models.Cleaner
		err := rows.Scan(
			&cleaner.Id,
			&cleaner.Name,
			&cleaner.Surname,
		)
		if err != nil {
			return nil, err
		}
		cleaners = append(cleaners, cleaner)
	}

	return cleaners, nil
}

// Update modifies an existing cleaner
func (r *cleanerRepository) Update(ctx context.Context, cleaner *models.Cleaner) error {
	query := `
		UPDATE cleaners
		SET name = $1, surname = $2
		WHERE id = $3`

	result, err := r.db.ExecContext(ctx, query,
		cleaner.Name,
		cleaner.Surname,
		cleaner.Id,
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

// Delete removes a cleaner by its ID
func (r *cleanerRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM cleaners WHERE id = $1`

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
