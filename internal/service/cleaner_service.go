package service

import (
	"context"
	"fmt"

	"github.com/StEvseeva/cleany/internal/models"
	"github.com/StEvseeva/cleany/internal/repository"
)

// CleanerService defines the interface for cleaner business operations
type CleanerService interface {
	CreateCleaner(ctx context.Context, req *models.CleanerCreateRequest) (*models.Cleaner, error)
	GetCleaner(ctx context.Context, id int) (*models.Cleaner, error)
	GetAllCleaners(ctx context.Context) ([]models.Cleaner, error)
	UpdateCleaner(ctx context.Context, id int, req *models.CleanerUpdateRequest) (*models.Cleaner, error)
	DeleteCleaner(ctx context.Context, id int) error
}

// cleanerService implements CleanerService
type cleanerService struct {
	cleanerRepo repository.CleanerRepository
}

// NewCleanerService creates a new cleaner service
func NewCleanerService(cleanerRepo repository.CleanerRepository) CleanerService {
	return &cleanerService{
		cleanerRepo: cleanerRepo,
	}
}

// CreateCleaner creates a new cleaner with validation
func (s *cleanerService) CreateCleaner(ctx context.Context, req *models.CleanerCreateRequest) (*models.Cleaner, error) {
	// Validate input
	if req.Name == "" {
		return nil, fmt.Errorf("cleaner name is required")
	}
	if req.Surname == "" {
		return nil, fmt.Errorf("cleaner surname is required")
	}

	// Create cleaner
	cleaner := &models.Cleaner{
		Name:    req.Name,
		Surname: req.Surname,
	}

	err := s.cleanerRepo.Create(ctx, cleaner)
	if err != nil {
		return nil, fmt.Errorf("failed to create cleaner: %w", err)
	}

	return cleaner, nil
}

// GetCleaner retrieves a cleaner by ID
func (s *cleanerService) GetCleaner(ctx context.Context, id int) (*models.Cleaner, error) {
	cleaner, err := s.cleanerRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("cleaner not found: %w", err)
	}

	return cleaner, nil
}

// GetAllCleaners retrieves all cleaners
func (s *cleanerService) GetAllCleaners(ctx context.Context) ([]models.Cleaner, error) {
	cleaners, err := s.cleanerRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get cleaners: %w", err)
	}

	return cleaners, nil
}

// UpdateCleaner updates an existing cleaner
func (s *cleanerService) UpdateCleaner(ctx context.Context, id int, req *models.CleanerUpdateRequest) (*models.Cleaner, error) {
	// Check if cleaner exists
	existingCleaner, err := s.cleanerRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("cleaner not found: %w", err)
	}

	// Update fields if provided
	if req.Name != nil {
		if *req.Name == "" {
			return nil, fmt.Errorf("cleaner name cannot be empty")
		}
		existingCleaner.Name = *req.Name
	}
	if req.Surname != nil {
		if *req.Surname == "" {
			return nil, fmt.Errorf("cleaner surname cannot be empty")
		}
		existingCleaner.Surname = *req.Surname
	}

	err = s.cleanerRepo.Update(ctx, existingCleaner)
	if err != nil {
		return nil, fmt.Errorf("failed to update cleaner: %w", err)
	}

	return existingCleaner, nil
}

// DeleteCleaner deletes a cleaner by ID
func (s *cleanerService) DeleteCleaner(ctx context.Context, id int) error {
	// Check if cleaner exists
	_, err := s.cleanerRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("cleaner not found: %w", err)
	}

	err = s.cleanerRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete cleaner: %w", err)
	}

	return nil
}
