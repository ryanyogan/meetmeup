package models

// User defines the structure of a user payload
type User struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Meetups  []*Meetup `json:"meetups"`
}
