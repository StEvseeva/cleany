package service

import "github.com/StEvseeva/cleany/internal/repository"

type Service interface {
	BookingService
	CleanerService
	RoomService
	CleaningOrderService
}

type service struct {
	BookingService
	CleanerService
	RoomService
	CleaningOrderService
}

// roomService implements RoomService
type roomService struct {
	roomRepo repository.RoomRepository
}

// bookingService implements BookingService
type bookingService struct {
	bookingRepo          repository.BookingRepository
	roomRepo             repository.RoomRepository
	cleaningOrderService CleaningOrderService
}

// cleanerService implements CleanerService
type cleanerService struct {
	cleanerRepo repository.CleanerRepository
}

// cleaningOrderService implements CleaningOrderService
type cleaningOrderService struct {
	cleaningOrderRepo repository.CleaningOrderRepository
	bookingRepo       repository.BookingRepository
	cleanerRepo       repository.CleanerRepository
}

// NewCleanerService creates a new cleaner service
func NewCleanerService(cleanerRepo repository.CleanerRepository) CleanerService {
	return &cleanerService{
		cleanerRepo: cleanerRepo,
	}
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

// NewRoomService creates a new room service
func NewRoomService(roomRepo repository.RoomRepository) RoomService {
	return &roomService{
		roomRepo: roomRepo,
	}
}

// NewBookingService creates a new booking service
func NewBookingService(
	bookingRepo repository.BookingRepository,
	roomRepo repository.RoomRepository,
	cleaningOrderService CleaningOrderService) BookingService {
	return &bookingService{
		bookingRepo:          bookingRepo,
		roomRepo:             roomRepo,
		cleaningOrderService: cleaningOrderService,
	}
}

func NewService(
	cleanerRepo repository.CleanerRepository,
	bookingRepo repository.BookingRepository,
	roomRepo repository.RoomRepository,
	cleaningOrderRepo repository.CleaningOrderRepository) Service {
	order := NewCleaningOrderService(cleaningOrderRepo, bookingRepo, cleanerRepo)
	return &service{
		BookingService:       NewBookingService(bookingRepo, roomRepo, order),
		CleanerService:       NewCleanerService(cleanerRepo),
		RoomService:          NewRoomService(roomRepo),
		CleaningOrderService: order,
	}
}

type Server struct {
	service Service
}
