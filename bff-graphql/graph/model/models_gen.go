// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID        string  `json:"id"`
	CreatedAt *string `json:"createdAt"`
	UpdatedAt *string `json:"updatedAt"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
}
