package handlers

import (
	"net/http"

	"bookinfo/ratings/data"
)

// swagger:route GET /Ratings/{id} Ratings listSingleRating
// Return a list of Ratings from the database
// responses:
//	200: RatingResponse
//	404: errorResponse

// ListSingle handles GET requests
func (p *APIContext) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := getRatingID(r)

	p.l.Println("[DEBUG] get record id", id)

	prod := data.GetRatingByID(id)

	err := data.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing Rating", err)
	}
}
