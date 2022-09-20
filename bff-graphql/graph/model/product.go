package model

type Product struct {
	ID          string  `json:"id"`
	CreatedAt   *string `json:"createdAt"`
	UpdatedAt   *string `json:"updatedAt"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Thumb       string  `json:"thumb"`
	Availables  int     `json:"availables"`
	Price       float64 `json:"price"`
}
