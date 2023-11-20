package book

import "encoding/json"

type BodyPost struct {
	Title       string      `json:"title" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Rate        json.Number `json:"rate" binding:"required,number"`
}

type BodyPut struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       json.Number
	Rate        json.Number
}
