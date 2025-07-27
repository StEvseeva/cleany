package service

import (
	"context"
	"fmt"
	"time"

	"github.com/StEvseeva/cleany/internal/models"
	"github.com/StEvseeva/cleany/internal/repository"
)

// CleaningOrderService defines the interface for cleaning order business operations
type CleaningOrderService interface {
	CreateCleaningOrder(ctx context.Context, req *models.CleaningOrderCreateRequest) (*models.CleaningOrder, error)
	CreateCleaningOrdersForBooking(ctx context.Context, booking models.Booking) ([]models.CleaningOrderCreateRequest, error)
	GetCleaningOrder(ctx context.Context, id int) (*models.CleaningOrder, error)
	GetAllCleaningOrders(ctx context.Context) ([]models.CleaningOrder, error)
	GetAllCleaningOrdersByCleanerId(ctx context.Context, cleaner_id int) ([]models.CleaningOrder, error)
	UpdateCleaningOrder(ctx context.Context, id int, req *models.CleaningOrderUpdateRequest) (*models.CleaningOrder, error)
	DeleteCleaningOrder(ctx context.Context, id int) error
	AssignCleaner(ctx context.Context, orderID int, req *models.CleanerOrderCreateRequest) error
	RemoveCleaner(ctx context.Context, orderID, cleanerID int) error
}

// cleaningOrderService implements CleaningOrderService
type cleaningOrderService struct {
	cleaningOrderRepo repository.CleaningOrderRepository
	bookingRepo       repository.BookingRepository
	cleanerRepo       repository.CleanerRepository
}

// NewCleaningOrderService creates a new cleaning order service
func NewCleaningOrderService(
	cleaningOrderRepo repository.CleaningOrderRepository,
	bookingRepo repository.BookingRepository,
	cleanerRepo repository.CleanerRepository,
) CleaningOrderService {
	return &cleaningOrderService{
		cleaningOrderRepo: cleaningOrderRepo,
		bookingRepo:       bookingRepo,
		cleanerRepo:       cleanerRepo,
	}
}

// CreateCleaningOrder creates a new cleaning order with validation
func (s *cleaningOrderService) CreateCleaningOrder(ctx context.Context, req *models.CleaningOrderCreateRequest) (*models.CleaningOrder, error) {
	// Validate that the booking exists
	booking, err := s.bookingRepo.GetByID(ctx, req.BookingId)
	if err != nil {
		return nil, fmt.Errorf("booking not found: %w", err)
	}

	// Validate cost
	if req.Cost < 0 {
		return nil, fmt.Errorf("cost must be non-negative")
	}
	if req.Cost == 0 {
		req.Cost = countOrderCost(*booking, *req.CleaningType)
	}

	// Create cleaning order
	order := &models.CleaningOrder{
		BookingId:    req.BookingId,
		CleaningTs:   &req.CleaningTs,
		CleaningType: req.CleaningType,
		Cost:         req.Cost,
		Done:         req.Done,
		Notes:        req.Notes,
	}

	err = s.cleaningOrderRepo.Create(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("failed to create cleaning order: %w", err)
	}

	return order, nil
}

// GetCleaningOrder retrieves a cleaning order by ID
func (s *cleaningOrderService) GetCleaningOrder(ctx context.Context, id int) (*models.CleaningOrder, error) {
	order, err := s.cleaningOrderRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("cleaning order not found: %w", err)
	}

	return order, nil
}

// GetAllCleaningOrders retrieves all cleaning orders
func (s *cleaningOrderService) GetAllCleaningOrders(ctx context.Context) ([]models.CleaningOrder, error) {
	orders, err := s.cleaningOrderRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get cleaning orders: %w", err)
	}

	return orders, nil
}

// GetAllCleaningOrders retrieves all cleaning orders assigned to a cleaner
func (s *cleaningOrderService) GetAllCleaningOrdersByCleanerId(ctx context.Context, cleaner_id int) ([]models.CleaningOrder, error) {
	orders, err := s.cleaningOrderRepo.GetAllByCleanerId(ctx, cleaner_id)
	if err != nil {
		return nil, fmt.Errorf("failed to get cleaning orders: %w", err)
	}

	return orders, nil
}

// UpdateCleaningOrder updates an existing cleaning order
func (s *cleaningOrderService) UpdateCleaningOrder(ctx context.Context, id int, req *models.CleaningOrderUpdateRequest) (*models.CleaningOrder, error) {
	// Check if cleaning order exists
	existingOrder, err := s.cleaningOrderRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("cleaning order not found: %w", err)
	}

	// Validate that the booking exists
	_, err = s.bookingRepo.GetByID(ctx, req.BookingId)
	if err != nil {
		return nil, fmt.Errorf("booking not found: %w", err)
	}

	// Validate cost
	if req.Cost < 0 {
		return nil, fmt.Errorf("cost must be non-negative")
	}

	// Update cleaning order
	existingOrder.BookingId = req.BookingId
	existingOrder.CleaningTs = &req.CleaningTs
	existingOrder.CleaningType = req.CleaningType
	existingOrder.Cost = req.Cost
	existingOrder.Done = req.Done
	existingOrder.Notes = req.Notes

	err = s.cleaningOrderRepo.Update(ctx, existingOrder)
	if err != nil {
		return nil, fmt.Errorf("failed to update cleaning order: %w", err)
	}

	return existingOrder, nil
}

// DeleteCleaningOrder deletes a cleaning order by ID
func (s *cleaningOrderService) DeleteCleaningOrder(ctx context.Context, id int) error {
	// Check if cleaning order exists
	_, err := s.cleaningOrderRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("cleaning order not found: %w", err)
	}

	err = s.cleaningOrderRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete cleaning order: %w", err)
	}

	return nil
}

// AssignCleaner assigns a cleaner to a cleaning order
func (s *cleaningOrderService) AssignCleaner(ctx context.Context, orderID int, req *models.CleanerOrderCreateRequest) error {
	// Validate that the cleaning order exists
	_, err := s.cleaningOrderRepo.GetByID(ctx, orderID)
	if err != nil {
		return fmt.Errorf("cleaning order not found: %w", err)
	}

	// Validate that the cleaner exists
	_, err = s.cleanerRepo.GetByID(ctx, req.CleanerId)
	if err != nil {
		return fmt.Errorf("cleaner not found: %w", err)
	}

	err = s.cleaningOrderRepo.AssignCleaner(ctx, orderID, req.CleanerId)
	if err != nil {
		return fmt.Errorf("failed to assign cleaner: %w", err)
	}

	return nil
}

// RemoveCleaner removes a cleaner from a cleaning order
func (s *cleaningOrderService) RemoveCleaner(ctx context.Context, orderID, cleanerID int) error {
	// Validate that the cleaning order exists
	_, err := s.cleaningOrderRepo.GetByID(ctx, orderID)
	if err != nil {
		return fmt.Errorf("cleaning order not found: %w", err)
	}

	// Validate that the cleaner exists
	_, err = s.cleanerRepo.GetByID(ctx, cleanerID)
	if err != nil {
		return fmt.Errorf("cleaner not found: %w", err)
	}

	err = s.cleaningOrderRepo.RemoveCleaner(ctx, orderID, cleanerID)
	if err != nil {
		return fmt.Errorf("failed to remove cleaner: %w", err)
	}

	return nil
}

func (s *cleaningOrderService) CreateCleaningOrdersForBooking(ctx context.Context, booking models.Booking) ([]models.CleaningOrderCreateRequest, error) {
	// Validate that the booking exists
	if _, err := s.bookingRepo.GetByID(ctx, booking.Id); err != nil {
		return nil, fmt.Errorf("booking not found: %w", err)
	}

	orders_queue, err := collectOrdersQueue(booking)
	if err != nil {
		return nil, fmt.Errorf("failed to collect queue for cleaning orders: %w", err)
	}
	_, err = s.cleaningOrderRepo.CreateMany(ctx, orders_queue)

	if err != nil {
		return nil, fmt.Errorf("failed to create cleaning orders: %w", err)
	}

	return orders_queue, nil
}

func collectOrdersQueue(booking models.Booking) ([]models.CleaningOrderCreateRequest, error) {

	baseCleaningTime := 13 * time.Hour

	orders_queue := []models.CleaningOrderCreateRequest{}

	checkInDate := booking.CheckInTs.Truncate(24 * time.Hour)
	checkOutDate := booking.CheckOutTs.Truncate(24 * time.Hour).Add(-24 * time.Hour)

	for date := checkInDate; date.Before(checkOutDate); date = date.Add(24 * time.Hour) {
		cleaningType := "periodic"
		orders_queue = append(orders_queue, models.CleaningOrderCreateRequest{
			BookingId:    booking.Id,
			CleaningTs:   date.Add(baseCleaningTime),
			CleaningType: &cleaningType,
			Cost:         countOrderCost(booking, cleaningType),
		})
	}
	cleaningType := "general"
	orders_queue = append(orders_queue, models.CleaningOrderCreateRequest{
		BookingId:    booking.Id,
		CleaningTs:   booking.CheckOutTs.Add(1 * time.Hour),
		CleaningType: &cleaningType,
		Cost:         countOrderCost(booking, "general"),
	})

	return orders_queue, nil
}

// countOrderCost count cost of one order
// TODO: add internal logic & table "cleaning_type"
func countOrderCost(booking models.Booking, cleaningType string) int {
	cost := 100
	if cleaningType == "general" {
		cost *= 2
	}
	return cost
}
