// Package classification of Details API
//
// Documentation for Details API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import (
	"bookinfo/details/data"
	"bookinfo/details/dto"
)

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// A list of ratings
// swagger:response RatingsResponse
type detailsResponseWrapper struct {
	// All current ratings
	// in: body
	Body []data.Detail
}

// Data structure representing a single rating
// swagger:response RatingResponse
type detailResponseWrapper struct {
	// Newly created rating
	// in: body
	Body data.Detail
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters updateDetail
type detailParamsWrapper struct {
	// rating data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body dto.DetailPrice
}

// swagger:parameters updateDetail
type detailIDParamsWrapper struct {
	// The id of the rating for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}
