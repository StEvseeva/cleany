package server

import (
	"net/http"

	"github.com/StEvseeva/cleany/internal/models"
	"github.com/labstack/echo/v4"
)

// GetCleaningOrders returns all cleaning orders
func (s *Server) GetCleaningOrders(ctx echo.Context) error {
	orders, err := s.cleaningOrderService.GetAllCleaningOrders(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, orders)
}

// PostCleaningOrders creates a new cleaning order
func (s *Server) PostCleaningOrders(ctx echo.Context) error {
	var req models.CleaningOrderCreateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	order, err := s.cleaningOrderService.CreateCleaningOrder(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, order)
}

// DeleteCleaningOrdersId deletes a cleaning order by ID
func (s *Server) DeleteCleaningOrdersId(ctx echo.Context, id int) error {
	err := s.cleaningOrderService.DeleteCleaningOrder(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return ctx.NoContent(http.StatusNoContent)
}

// GetCleaningOrdersId returns a cleaning order by ID
func (s *Server) GetCleaningOrdersId(ctx echo.Context, id int) error {
	order, err := s.cleaningOrderService.GetCleaningOrder(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, order)
}

// PutCleaningOrdersId updates a cleaning order by ID
func (s *Server) PutCleaningOrdersId(ctx echo.Context, id int) error {
	var req models.CleaningOrderUpdateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	order, err := s.cleaningOrderService.UpdateCleaningOrder(ctx.Request().Context(), id, &req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, order)
}

// PostCleaningOrdersIdCleaners assigns a cleaner to a cleaning order
func (s *Server) PostCleaningOrdersIdCleaners(ctx echo.Context, id int) error {
	var req models.CleanerOrderCreateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := s.cleaningOrderService.AssignCleaner(ctx.Request().Context(), id, &req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "cleaner assigned"})
}

// DeleteCleaningOrdersIdCleanersCleanerId removes a cleaner from a cleaning order
func (s *Server) DeleteCleaningOrdersIdCleanersCleanerId(ctx echo.Context, id int, cleanerId int) error {
	err := s.cleaningOrderService.RemoveCleaner(ctx.Request().Context(), id, cleanerId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "cleaner removed"})
}
