package server

import (
	"net/http"

	"github.com/StEvseeva/cleany/internal/models"
	"github.com/labstack/echo/v4"
)

// GetRooms returns all rooms
func (s *Server) GetRooms(ctx echo.Context) error {
	rooms, err := s.roomService.GetAllRooms(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, rooms)
}

// PostRooms creates a new room
func (s *Server) PostRooms(ctx echo.Context) error {
	var req models.RoomCreateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	room, err := s.roomService.CreateRoom(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, room)
}

// DeleteRoomsId deletes a room by ID
func (s *Server) DeleteRoomsId(ctx echo.Context, id int) error {
	err := s.roomService.DeleteRoom(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return ctx.NoContent(http.StatusNoContent)
}

// GetRoomsId returns a room by ID
func (s *Server) GetRoomsId(ctx echo.Context, id int) error {
	room, err := s.roomService.GetRoom(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, room)
}

// PutRoomsId updates a room by ID
func (s *Server) PutRoomsId(ctx echo.Context, id int) error {
	var req models.RoomUpdateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	room, err := s.roomService.UpdateRoom(ctx.Request().Context(), id, &req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, room)
}
