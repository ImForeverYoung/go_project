package model

type Professor struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Title      string `json:"title"`
	Email      string `json:"email"`
	Dataset    string `json:"dataset,omitempty"` // For internal use or extra data
}

// Request payload for creating/updating a professor
type ProfessorRequest struct {
	Name       string `json:"name"`
	Department string `json:"department"`
	Title      string `json:"title"`
	Email      string `json:"email"`
}
