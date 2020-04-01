package handlers

import (
	"fmt"
	"net/http"

	"bookinfo/ratings/data"
	"bookinfo/ratings/dto"
	"bookinfo/ratings/logger"
)

// swagger:route PUT /Ratings Ratings updateRating
// Update a Ratings details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update Ratings
func (p *APIContext) Update(rw http.ResponseWriter, r *http.Request) {

	// fetch the Rating from the context
	rating := r.Context().Value(KeyRating{}).(dto.NewRating)
	logger.Log(fmt.Sprintf("updating record id", rating.BookID), logger.DebugLevel)

	data.UpdateRating(rating)
	// write the no content success header
	rw.WriteHeader(http.StatusNoContent)
}
