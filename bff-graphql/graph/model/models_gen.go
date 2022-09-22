// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type DeleteProductInput struct {
	ID string `json:"ID"`
}

type GenericResponse struct {
	Success bool `json:"success"`
}

type NewOrderInput struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type NewProductInput struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Thumb       string  `json:"thumb"`
	Availables  int     `json:"availables"`
	Price       float64 `json:"price"`
}

type NewUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResponse struct {
	Token string `json:"Token"`
}

type UpdateProductAvailablesInput struct {
	ID         string `json:"ID"`
	ValueToAdd int    `json:"valueToAdd"`
}

type UpdateProductInput struct {
	ID          string  `json:"ID"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Thumb       string  `json:"thumb"`
	Price       float64 `json:"price"`
}
