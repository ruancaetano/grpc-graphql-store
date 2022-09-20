package model

type Order struct {
	ID        string   `json:"id"`
	CreatedAt *string  `json:"createdAt"`
	UpdatedAt *string  `json:"updatedAt"`
	User      *User    `json:"user"`
	Product   *Product `json:"product"`
	Quantity  int      `json:"quantity"`
}
