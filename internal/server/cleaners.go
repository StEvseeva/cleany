package server

import (
	"net/http"

	"github.com/StEvseeva/cleany/internal/models"
	"github.com/labstack/echo/v4"
)

// GetCleaners returns all cleaners
func (s *Server) GetCleaners(ctx echo.Context) error {
	cleaners, err := s.service.GetAllCleaners(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, cleaners)
}

// PostCleaners creates a new cleaner
func (s *Server) PostCleaners(ctx echo.Context) error {
	var req models.CleanerCreateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	cleaner, err := s.service.CreateCleaner(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, cleaner)
}

// DeleteCleanersId deletes a cleaner by ID
func (s *Server) DeleteCleanersId(ctx echo.Context, id int) error {
	err := s.service.DeleteCleaner(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return ctx.NoContent(http.StatusNoContent)
}

// GetCleanersId returns a cleaner by ID
func (s *Server) GetCleanersId(ctx echo.Context, id int) error {
	cleaner, err := s.service.GetCleaner(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, cleaner)
}

// PutCleanersId updates a cleaner by ID
func (s *Server) PutCleanersId(ctx echo.Context, id int) error {
	var req models.CleanerUpdateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	cleaner, err := s.service.UpdateCleaner(ctx.Request().Context(), id, &req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, cleaner)
}
