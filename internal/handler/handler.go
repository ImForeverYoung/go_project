package handler

import (
	"HW_5/internal/model"
	"HW_5/internal/storage"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	storage *storage.Storage
}

func NewHandler(storage *storage.Storage) *Handler {
	return &Handler{storage: storage}
}

// GetStudent handles GET /student/:id
// @Summary Get student details
// @Tags Student
// @Description Get detailed information about a student
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Success 200 {object} model.ServerResponse
// @Router /student/{id} [get]
func (h *Handler) GetStudent(c echo.Context) error {
	id := c.Param("id")

	student, err := h.storage.GetStudent(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ServerResponse{
			Status:  "error",
			Message: "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, model.ServerResponse{
		Status: "success",
		Data:   student,
	})
}

// GetAllSchedule handles GET /all_class_schedule
// @Summary Get all schedules
// @Tags Schedule
// @Description Get the entire schedule
// @Accept json
// @Produce json
// @Success 200 {object} model.ServerResponse
// @Router /all_class_schedule [get]
func (h *Handler) GetAllSchedule(c echo.Context) error {
	results, err := h.storage.GetAllSchedule(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ServerResponse{
			Status:  "error",
			Message: "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, model.ServerResponse{
		Status: "success",
		Data:   results,
	})
}

// GetGroupSchedule handles GET /schedule/group/:id
// @Summary Get group schedule
// @Tags Schedule
// @Description Get schedule for a specific group
// @Accept json
// @Produce json
// @Param id path int true "Group ID"
// @Success 200 {object} model.ServerResponse
// @Router /schedule/group/{id} [get]
func (h *Handler) GetGroupSchedule(c echo.Context) error {
	id := c.Param("id")

	results, err := h.storage.GetGroupSchedule(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ServerResponse{
			Status:  "error",
			Message: "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, model.ServerResponse{
		Status: "success",
		Data:   results,
	})
}

// MarkAttendance handles POST /attendance/subject
// @Summary Mark attendance
// @Tags Attendance
// @Description Record attendance for a student
// @Accept json
// @Produce json
// @Param input body model.Attendance true "Attendance Data"
// @Success 200 {object} model.ServerResponse
// @Router /attendance/subject [post]
func (h *Handler) MarkAttendance(ctx echo.Context) error {
	var request model.Attendance

	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	id, err := h.storage.MarkAttendance(ctx.Request().Context(), request)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ServerResponse{
			Status:  "error",
			Message: "Internal Server Error",
		})
	}
	message := fmt.Sprintf("google sheet with id=%d succefully wroten in database", id)
	return ctx.JSON(http.StatusOK, model.ServerResponse{
		Status: "success",
		Data:   message,
	})
}

// GetAttendanceBySubjectId handles GET /attendanceBySubjectId/:id
// @Summary Get attendance by subject
// @Tags Attendance
// @Description Get attendance records for a specific subject
// @Accept json
// @Produce json
// @Param id path int true "Subject ID"
// @Success 200 {object} model.ServerResponse
// @Router /attendanceBySubjectId/{id} [get]
func (h *Handler) GetAttendanceBySubjectId(c echo.Context) error {
	id := c.Param("id")

	results, err := h.storage.GetAttendanceBySubjectId(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ServerResponse{
			Status:  "error",
			Message: "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, model.ServerResponse{
		Status: "success",
		Data:   results,
	})
}

// GetAttendanceByStudentId handles GET /attendanceByStudentId/:id
// @Summary Get attendance by student
// @Tags Attendance
// @Description Get attendance records for a specific student
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Success 200 {object} model.ServerResponse
// @Router /attendanceByStudentId/{id} [get]
func (h *Handler) GetAttendanceByStudentId(c echo.Context) error {
	id := c.Param("id")

	results, err := h.storage.GetAttendanceByStudentId(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ServerResponse{
			Status:  "error",
			Message: "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, model.ServerResponse{
		Status: "success",
		Data:   results,
	})
}

// GetGroupStudentCounts handles GET /analytics/groups
// @Summary Get group sizes
// @Tags Analytics
// @Description Get student counts per group
// @Accept json
// @Produce json
// @Success 200 {object} model.ServerResponse
// @Router /analytics/groups [get]
func (h *Handler) GetGroupStudentCounts(c echo.Context) error {
	stats, err := h.storage.GetGroupStudentCounts(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ServerResponse{
			Status:  "error",
			Message: "Failed to fetch group statistics",
		})
	}

	return c.JSON(http.StatusOK, model.ServerResponse{
		Status: "success",
		Data:   stats,
	})
}

// GetStudentsWithAbsences handles GET /analytics/absences
// @Summary Get students with absences
// @Tags Analytics
// @Description Get students who have more than a minimum number of absences
// @Accept json
// @Produce json
// @Param min query int false "Minimum Absences" default(1)
// @Success 200 {object} model.ServerResponse
// @Router /analytics/absences [get]
func (h *Handler) GetStudentsWithAbsences(c echo.Context) error {
	minStr := c.QueryParam("min")
	min := 1
	if minStr != "" {
		var err error
		min, err = strconv.Atoi(minStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, model.ServerResponse{
				Status:  "error",
				Message: "Invalid 'min' parameter",
			})
		}
	}

	stats, err := h.storage.GetStudentsWithAbsences(c.Request().Context(), min)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ServerResponse{
			Status:  "error",
			Message: "Failed to fetch attendance statistics",
		})
	}

	return c.JSON(http.StatusOK, model.ServerResponse{
		Status: "success",
		Data:   stats,
	})
}
