package service

import (
	"context"
	"fmt"

	"github.com/StEvseeva/cleany/internal/models"
)

// BookingService defines the interface for booking business operations
type BookingService interface {
	CreateBooking(ctx context.Context, req *models.BookingCreateRequest) (*models.Booking, error)
	GetBooking(ctx context.Context, id int) (*models.Booking, error)
	GetAllBookings(ctx context.Context) ([]models.Booking, error)
	UpdateBooking(ctx context.Context, id int, req *models.BookingUpdateRequest) (*models.Booking, error)
	DeleteBooking(ctx context.Context, id int) error
}

// CreateBooking creates a new booking with validation
func (s *bookingService) CreateBooking(ctx context.Context, req *models.BookingCreateRequest) (*models.Booking, error) {
	// Validate that the room exists
	_, err := s.roomRepo.GetByID(ctx, req.RoomId)
	if err != nil {
		return nil, fmt.Errorf("room not found: %w", err)
	}

	// Validate check-in/check-out dates
	if req.CheckInTs.After(req.CheckOutTs) {
		return nil, fmt.Errorf("check-in date must be before check-out date")
	}

	// Create booking
	booking := &models.Booking{
		RoomId:     req.RoomId,
		CheckInTs:  &req.CheckInTs,
		CheckOutTs: &req.CheckOutTs,
		Guests:     req.Guests,
	}

	err = s.bookingRepo.Create(ctx, booking)
	if err != nil {
		return nil, fmt.Errorf("failed to create booking: %w", err)
	}

	_, err = s.cleaningOrderService.CreateCleaningOrdersForBooking(ctx, *booking)
	if err != nil {
		return nil, fmt.Errorf("failed to create cleaning orders for booking: %w", err)
	}

	return booking, nil
}

// GetBooking retrieves a booking by ID
func (s *bookingService) GetBooking(ctx context.Context, id int) (*models.Booking, error) {
	booking, err := s.bookingRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("booking not found: %w", err)
	}

	return booking, nil
}

// GetAllBookings retrieves all bookings
func (s *bookingService) GetAllBookings(ctx context.Context) ([]models.Booking, error) {
	bookings, err := s.bookingRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get bookings: %w", err)
	}

	return bookings, nil
}

// UpdateBooking updates an existing booking
func (s *bookingService) UpdateBooking(ctx context.Context, id int, req *models.BookingUpdateRequest) (*models.Booking, error) {
	// Check if booking exists
	existingBooking, err := s.bookingRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("booking not found: %w", err)
	}

	// Validate that the room exists
	_, err = s.roomRepo.GetByID(ctx, req.RoomId)
	if err != nil {
		return nil, fmt.Errorf("room not found: %w", err)
	}

	// Validate check-in/check-out dates
	if req.CheckInTs.After(req.CheckOutTs) {
		return nil, fmt.Errorf("check-in date must be before check-out date")
	}

	// Update booking
	existingBooking.RoomId = req.RoomId
	existingBooking.CheckInTs = &req.CheckInTs
	existingBooking.CheckOutTs = &req.CheckOutTs
	existingBooking.Guests = req.Guests

	err = s.bookingRepo.Update(ctx, existingBooking)
	if err != nil {
		return nil, fmt.Errorf("failed to update booking: %w", err)
	}

	return existingBooking, nil
}

// DeleteBooking deletes a booking by ID
func (s *bookingService) DeleteBooking(ctx context.Context, id int) error {
	// Check if booking exists
	_, err := s.bookingRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("booking not found: %w", err)
	}

	err = s.bookingRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete booking: %w", err)
	}

	return nil
}
