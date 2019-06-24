package domain

import "time"

// Wish describes structure of user's wishlist
type Wish struct {
	ID     string
	Title  string    `json:",omitempty"`
	UserID string    `json:",omitempty"`
	Done   bool      `json:",omitempty"`
	Date   time.Time `json:",omitempty"`
}
