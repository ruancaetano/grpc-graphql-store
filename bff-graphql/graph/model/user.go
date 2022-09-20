package model

type User struct {
	ID        string  `json:"id"`
	CreatedAt *string `json:"createdAt"`
	UpdatedAt *string `json:"updatedAt"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
}
