package book

import (
	"time"
)

type BookResponse struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Rate        int       `json:"rate"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
