package service

import (
	"context"
	"fmt"

	"github.com/StEvseeva/cleany/internal/models"
	"github.com/StEvseeva/cleany/internal/repository"
)

// RoomService defines the interface for room business operations
type RoomService interface {
	CreateRoom(ctx context.Context, req *models.RoomCreateRequest) (*models.Room, error)
	GetRoom(ctx context.Context, id int) (*models.Room, error)
	GetAllRooms(ctx context.Context) ([]models.Room, error)
	UpdateRoom(ctx context.Context, id int, req *models.RoomUpdateRequest) (*models.Room, error)
	DeleteRoom(ctx context.Context, id int) error
}

// roomService implements RoomService
type roomService struct {
	roomRepo repository.RoomRepository
}

// NewRoomService creates a new room service
func NewRoomService(roomRepo repository.RoomRepository) RoomService {
	return &roomService{
		roomRepo: roomRepo,
	}
}

// CreateRoom creates a new room with validation
func (s *roomService) CreateRoom(ctx context.Context, req *models.RoomCreateRequest) (*models.Room, error) {
	// Validate input
	if req.Floor < 0 {
		return nil, fmt.Errorf("floor number must be non-negative")
	}

	// Create room
	room := &models.Room{
		Floor: req.Floor,
		Desc:  req.Desc,
	}

	err := s.roomRepo.Create(ctx, room)
	if err != nil {
		return nil, fmt.Errorf("failed to create room: %w", err)
	}

	return room, nil
}

// GetRoom retrieves a room by ID
func (s *roomService) GetRoom(ctx context.Context, id int) (*models.Room, error) {
	room, err := s.roomRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("room not found: %w", err)
	}

	return room, nil
}

// GetAllRooms retrieves all rooms
func (s *roomService) GetAllRooms(ctx context.Context) ([]models.Room, error) {
	rooms, err := s.roomRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get rooms: %w", err)
	}

	return rooms, nil
}

// UpdateRoom updates an existing room
func (s *roomService) UpdateRoom(ctx context.Context, id int, req *models.RoomUpdateRequest) (*models.Room, error) {
	// Check if room exists
	existingRoom, err := s.roomRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("room not found: %w", err)
	}

	// Update fields if provided
	if req.Floor != nil {
		if *req.Floor < 0 {
			return nil, fmt.Errorf("floor number must be non-negative")
		}
		existingRoom.Floor = *req.Floor
	}
	if req.Desc != nil {
		existingRoom.Desc = req.Desc
	}

	err = s.roomRepo.Update(ctx, existingRoom)
	if err != nil {
		return nil, fmt.Errorf("failed to update room: %w", err)
	}

	return existingRoom, nil
}

// DeleteRoom deletes a room by ID
func (s *roomService) DeleteRoom(ctx context.Context, id int) error {
	// Check if room exists
	_, err := s.roomRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("room not found: %w", err)
	}

	err = s.roomRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete room: %w", err)
	}

	return nil
}
