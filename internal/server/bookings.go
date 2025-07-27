package server

import (
	"net/http"

	"github.com/StEvseeva/cleany/internal/models"
	"github.com/labstack/echo/v4"
)

// GetBookings returns all bookings
func (s *Server) GetBookings(ctx echo.Context) error {
	bookings, err := s.bookingService.GetAllBookings(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, bookings)
}

// PostBookings creates a new booking
func (s *Server) PostBookings(ctx echo.Context) error {
	var req models.BookingCreateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	booking, err := s.bookingService.CreateBooking(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, booking)
}

// DeleteBookingsId deletes a booking by ID
func (s *Server) DeleteBookingsId(ctx echo.Context, id int) error {
	err := s.bookingService.DeleteBooking(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return ctx.NoContent(http.StatusNoContent)
}

// GetBookingsId returns a booking by ID
func (s *Server) GetBookingsId(ctx echo.Context, id int) error {
	booking, err := s.bookingService.GetBooking(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, booking)
}

// PutBookingsId updates a booking by ID
func (s *Server) PutBookingsId(ctx echo.Context, id int) error {
	var req models.BookingUpdateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	booking, err := s.bookingService.UpdateBooking(ctx.Request().Context(), id, &req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, booking)
}
