package domain

// UserRole keeps system roles of users
type UserRole int

const (
	DefaultUser UserRole = 0
	Admin       UserRole = 1
)

// User describes list of software users
type User struct {
	ID           string
	Name         string   `json:",omitempty"`
	Surname      string   `json:",omitempty"`
	Login        string   `json:",omitempty"`
	PasswordHash string   `json:",omitempty"`
	Role         UserRole `json:",omitempty"`
}
