package http

import (
	"net/http"
	"strconv"

	"github.com/dakdikduk/galactic-api/domain"
	"github.com/labstack/echo/v4"
)

type SpacecraftHandler struct {
	spacecraftUseCase domain.SpacecraftUseCase
}

func NewSpacecraftHandler(e *echo.Echo, spacecraftUseCase domain.SpacecraftUseCase) {
	handler := &SpacecraftHandler{
		spacecraftUseCase: spacecraftUseCase,
	}
	e.GET("/spacecrafts", handler.ListSpacecrafts)
	e.POST("/spacecrafts", handler.CreateSpacecraft)
	e.GET("/spacecrafts/:id", handler.GetSpacecraftByID)
	e.DELETE("/spacecrafts/:id", handler.DeleteSpacecraft)
	e.PATCH("/spacecrafts/:id", handler.UpdateSpacecraft)
}

func (h *SpacecraftHandler) ListSpacecrafts(c echo.Context) error {
	// Parse query parameters
	params := domain.ListSpacecraftParams{
		Name:    c.QueryParam("name"),
		Class:   c.QueryParam("class"),
		Status:  c.QueryParam("status"),
		Page:    parseQueryParamInt(c.QueryParam("page")),
		PerPage: parseQueryParamInt(c.QueryParam("per_page")),
	}

	// Call the use case to list spacecrafts
	spacecrafts, err := h.spacecraftUseCase.List(c.Request().Context(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	// Respond with the list of spacecrafts
	return c.JSON(http.StatusOK, mapSpacecraftToResponse(spacecrafts))
}

func (h *SpacecraftHandler) GetSpacecraftByID(c echo.Context) error {
	// Extract spacecraft ID from the URL path
	spacecraftID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid spacecraft ID"})
	}

	// Call the use case to get a spacecraft by ID
	spacecraft, err := h.spacecraftUseCase.GetByID(c.Request().Context(), spacecraftID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Spacecraft not found"})
	}

	// Respond with the spacecraft details
	return c.JSON(http.StatusOK, spacecraft)
}

func (h *SpacecraftHandler) CreateSpacecraft(c echo.Context) error {
	// Parse the request body to get the spacecraft data
	var spacecraft domain.Spacecraft
	if err := c.Bind(&spacecraft); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	// Call the use case to create a new spacecraft
	h.spacecraftUseCase.Create(c.Request().Context(), spacecraft)

	// Respond with a success message
	return c.JSON(http.StatusCreated, map[string]string{"message": "Spacecraft created successfully"})
}

func (h *SpacecraftHandler) UpdateSpacecraft(c echo.Context) error {
	// Extract spacecraft ID from the URL path
	spacecraftID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid spacecraft ID"})
	}

	// Parse the request body to get the updated spacecraft data
	var updatedSpacecraft domain.Spacecraft
	if err := c.Bind(&updatedSpacecraft); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	// Call the use case to update the spacecraft
	updatedSpacecraft.ID = uint(spacecraftID)
	h.spacecraftUseCase.Update(c.Request().Context(), updatedSpacecraft)

	// Respond with a success message
	return c.JSON(http.StatusOK, map[string]string{"message": "Spacecraft updated successfully"})
}

func (h *SpacecraftHandler) DeleteSpacecraft(c echo.Context) error {
	// Extract spacecraft ID from the URL path
	spacecraftID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid spacecraft ID"})
	}

	// Call the use case to delete the spacecraft
	h.spacecraftUseCase.Delete(c.Request().Context(), spacecraftID)

	// Respond with a success message
	return c.JSON(http.StatusOK, map[string]string{"message": "Spacecraft deleted successfully"})
}

// Helper function to parse integer query parameters
func parseQueryParamInt(value string) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return result
}
