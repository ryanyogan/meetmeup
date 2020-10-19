package models

// Meetup defines the structure of a meetup payload
type Meetup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      string `json:"userID"`
}
