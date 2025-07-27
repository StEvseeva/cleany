package server

import (
	"github.com/StEvseeva/cleany/internal/service"
)

type Server struct {
	bookingService       service.BookingService
	cleanerService       service.CleanerService
	roomService          service.RoomService
	cleaningOrderService service.CleaningOrderService
}

func NewServer(
	bookingService service.BookingService,
	cleanerService service.CleanerService,
	roomService service.RoomService,
	cleaningOrderService service.CleaningOrderService,
) *Server {
	return &Server{
		bookingService:       bookingService,
		cleanerService:       cleanerService,
		roomService:          roomService,
		cleaningOrderService: cleaningOrderService,
	}
}
