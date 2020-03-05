package dto

// NewRating defines the structure for an API Rating
// swagger:model
type NewRating struct {
	// the id of the book
	//
	// required: false
	// min: 1
	BookID int `json:"bookid"` // Unique identifier for the book

	// the rating of the book
	//
	// required: true
	// min: 0.01
	Rating float32 `json:"rating" validate:"required,gte=0"`
}
