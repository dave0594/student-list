package models

type Student struct {
	ID          int      `json:"id"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	Email       string   `json:"email"`
	PhoneNumber string   `json:"phone_number"`
	CreatedAt   NullTime `json:"created_at"`
	UpdatedAt   NullTime `json:"updated_at"`
}
