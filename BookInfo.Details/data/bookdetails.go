package data

import (
	"bookinfo/details/dto"
	"fmt"
	"time"
)

// ErrDetailNotFound is an error raised when a detail can not be found in the database
var ErrDetailNotFound = fmt.Errorf("Detail not found")

// Detail defines the structure for an Book detail
// swagger:model
type Detail struct {
	// the id of the book
	//
	// required: false
	// min: 1
	BookID int `json:"bookid"` // Unique identifier for the book

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

	// the number of books in the stock
	//
	// required: false
	// min: 0
	CurrentStock int `json:"currentstock"` // the number of books in the stock
}

// Details defines a slice of Detail
type Details []*Detail

// GetDetails returns all Details from the database
func GetDetails() Details {
	return detailList
}

// GetDetailByID returns a single Detail which matches the id from the
// database.
// If a Detail is not found this function returns a DetailNotFound error
func GetDetailByID(id int) (*Detail, error) {
	i := findIndexByBookID(id)
	if id == -1 {
		return nil, ErrDetailNotFound
	}

	return detailList[i], nil
}

// UpdateDetail replaces a Detail in the database with the given
// item.
func UpdateDetail(det dto.DetailPrice) error {
	i := findIndexByBookID(det.BookID)
	if i == -1 {
		return ErrDetailNotFound
	} else {
		// update the Detail in the DB
		detailList[i].Price = det.Price
		return nil
	}
}

// AddDetail add a new Detail to the database with the given
// item.
func AddDetail(det dto.Detail) {
	maxID := detailList[len(detailList)-1].BookID
	detail := &Detail{
		BookID:      maxID + 1,
		Name:        det.Name,
		ISBN:        det.ISBN,
		Author:      det.Author,
		PublishDate: det.PublishDate,
		Price:       det.Price,
	}
	detailList = append(detailList, detail)
}

// findIndexByBookID finds the index of a Detail in the database
// returns -1 when no Detail can be found
func findIndexByBookID(id int) int {
	for i, p := range detailList {
		if p.BookID == id {
			return i
		}
	}

	return -1
}

var detailList = []*Detail{
	&Detail{
		BookID:      1,
		Name:        "Lord of the Rings: The Fellowship of the Ring",
		ISBN:        "123AS123",
		Author:      "J.R.R. Tolkien",
		PublishDate: time.Date(1954, time.June, 29, 0, 0, 0, 0, time.UTC),
		Price:       55.29,
	},
	&Detail{
		BookID:      2,
		Name:        "Lord of the Rings: The Two Towers",
		ISBN:        "123AS124",
		Author:      "J.R.R. Tolkien",
		PublishDate: time.Date(1954, time.November, 11, 0, 0, 0, 0, time.UTC),
		Price:       53.29,
	},
}
