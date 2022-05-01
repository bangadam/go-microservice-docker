package models

type Product struct {
	ID		int    `json:"id,omitempty"`
	Name	string `json:"name"`
	Price	int    `json:"price"`
}