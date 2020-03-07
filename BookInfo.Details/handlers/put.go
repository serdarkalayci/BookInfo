package handlers

import (
	"net/http"

	"bookinfo/details/data"
	"bookinfo/details/dto"
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

	// fetch the DetailPrice from the context
	detail := r.Context().Value(KeyDetail{}).(dto.DetailPrice)
	p.l.Println("[DEBUG] updating record id", detail.BookID)

	data.UpdateDetail(detail)
	// write the no content success header
	rw.WriteHeader(http.StatusNoContent)
}
