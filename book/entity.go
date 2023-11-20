package book

import "time"

type Book struct {
	Id          int
	Title       string
	Description string
	Price       int
	Rate        int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
