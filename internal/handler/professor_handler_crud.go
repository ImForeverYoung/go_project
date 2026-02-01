package handler

import (
	"HW_5/internal/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// POST /api/professors
func (h *Handler) CreateProfessor(c echo.Context) error {
	var req model.ProfessorRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ServerResponse{Status: "error", Message: "Invalid request"})
	}

	id, err := h.storage.CreateProfessor(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ServerResponse{Status: "error", Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, model.ServerResponse{
		Status: "success",
		Data:   map[string]int{"id": id},
	})
}

// GET /api/professors/:id
func (h *Handler) GetProfessor(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ServerResponse{Status: "error", Message: "Invalid ID"})
	}

	prof, err := h.storage.GetProfessor(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ServerResponse{Status: "error", Message: "Professor not found"})
	}

	return c.JSON(http.StatusOK, model.ServerResponse{
		Status: "success",
		Data:   prof,
	})
}

// GET /api/professors
func (h *Handler) ListProfessors(c echo.Context) error {
	profs, err := h.storage.ListProfessors(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ServerResponse{Status: "error", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, model.ServerResponse{
		Status: "success",
		Data:   profs,
	})
}

// PUT /api/professors/:id
func (h *Handler) UpdateProfessor(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ServerResponse{Status: "error", Message: "Invalid ID"})
	}

	var req model.ProfessorRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ServerResponse{Status: "error", Message: "Invalid request"})
	}

	if err := h.storage.UpdateProfessor(c.Request().Context(), id, req); err != nil {
		return c.JSON(http.StatusInternalServerError, model.ServerResponse{Status: "error", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, model.ServerResponse{Status: "success", Message: "Professor updated"})
}

// DELETE /api/professors/:id
func (h *Handler) DeleteProfessor(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ServerResponse{Status: "error", Message: "Invalid ID"})
	}

	if err := h.storage.DeleteProfessor(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, model.ServerResponse{Status: "error", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, model.ServerResponse{Status: "success", Message: "Professor deleted"})
}
