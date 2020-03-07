package dto

import "time"

// DetailPrice defines the structure for an Book to update just the price
// swagger:model
type DetailPrice struct {
	// the id of the book
	//
	// required: false
	// min: 1
	BookID int `json:"bookid"` // Unique identifier for the book

	// the price of the book
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"required,gte=0"`
}

// Detail defines the structure for an Book detail
// swagger:model
type Detail struct {
	// the name of the book
	//
	// required: true
	Name string `json:"name" validate:"required"`

	// the ISBN of the book
	//
	// required: true
	ISBN string `json:"isbn" validate:"required"`

	// the author of the book
	//
	// required: true
	Author string `json:"author" validate:"required"`

	// the publish date of the book
	//
	// required: true
	PublishDate time.Time `json:"publishdate" validate:"required"`

	// the price of the book
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"required,gte=0"`
}
