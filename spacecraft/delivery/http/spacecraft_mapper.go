package http

import "github.com/dakdikduk/galactic-api/domain"

type ListSpacecraftsResponse struct {
	Data []ListSpacecraftItem
}

type ListSpacecraftItem struct {
	ID     int
	Name   string
	Status string
}

type SuccessResponse struct {
	Success bool
}

func mapSpacecraftToResponse(spacecrafts []domain.Spacecraft) ListSpacecraftsResponse {
	data := []ListSpacecraftItem{}

	for _, v := range spacecrafts {
		data = append(data, ListSpacecraftItem{ID: int(v.ID), Name: v.Name, Status: v.Status})
	}

	return ListSpacecraftsResponse{Data: data}
}
