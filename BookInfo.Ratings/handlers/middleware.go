package handlers

import (
	"context"
	"net/http"

	"bookinfo/ratings/data"
	"bookinfo/ratings/dto"
	"bookinfo/ratings/logger"
)

// MiddlewareValidateProduct validates the product in the request and calls next if ok
func (apiContext *APIContext) MiddlewareValidateNewRating(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rating := &dto.NewRating{}

		err := data.FromJSON(rating, r.Body)
		if err != nil {
			logger.Log("deserializing rating", logger.ErrorLevel, err)

			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		// validate the product
		errs := apiContext.v.Validate(rating)
		if len(errs) != 0 {
			logger.Log("[ERROR] validating product", logger.ErrorLevel, errs[0])

			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}

		// add the rating to the context
		ctx := context.WithValue(r.Context(), KeyRating{}, rating)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
