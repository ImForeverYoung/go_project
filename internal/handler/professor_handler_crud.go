package handler

import (
	"HW_5/internal/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateProfessor handles POST /api/professors
// @Summary Create a new professor
// @Tags Professors
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param input body model.ProfessorRequest true "Professor Data"
// @Success 201 {object} model.ServerResponse
// @Router /professors [post]
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

// GetProfessor handles GET /api/professors/:id
// @Summary Get a professor by ID
// @Tags Professors
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "Professor ID"
// @Success 200 {object} model.ServerResponse
// @Router /professors/{id} [get]
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

// ListProfessors handles GET /api/professors
// @Summary List all professors
// @Tags Professors
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} model.ServerResponse
// @Router /professors [get]
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

// UpdateProfessor handles PUT /api/professors/:id
// @Summary Update a professor
// @Tags Professors
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Professor ID"
// @Param input body model.ProfessorRequest true "Updated Data"
// @Success 200 {object} model.ServerResponse
// @Router /professors/{id} [put]
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

// DeleteProfessor handles DELETE /api/professors/:id
// @Summary Delete a professor
// @Tags Professors
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "Professor ID"
// @Success 200 {object} model.ServerResponse
// @Router /professors/{id} [delete]
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
