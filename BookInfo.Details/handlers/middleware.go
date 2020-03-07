package handlers

import (
	"context"
	"net/http"

	"bookinfo/details/data"
	"bookinfo/details/dto"
)

// MiddlewareValidateNewDetail validates new book detail in the request and calls next if ok
func (apiContext *APIContext) MiddlewareValidateNewDetail(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		detail := &dto.Detail{}

		err := data.FromJSON(detail, r.Body)
		if err != nil {
			apiContext.l.Println("[ERROR] deserializing book detail", err)

			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		// validate the product
		errs := apiContext.v.Validate(detail)
		if len(errs) != 0 {
			apiContext.l.Println("[ERROR] validating book detail", errs)

			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}

		// add the rating to the context
		ctx := context.WithValue(r.Context(), KeyDetail{}, detail)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}

// MiddlewareValidateDetailPrice validates new book detail in the request and calls next if ok
func (apiContext *APIContext) MiddlewareValidateDetailPrice(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		detprice := &dto.DetailPrice{}

		err := data.FromJSON(detprice, r.Body)
		if err != nil {
			apiContext.l.Println("[ERROR] deserializing price data", err)

			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		// validate the product
		errs := apiContext.v.Validate(detprice)
		if len(errs) != 0 {
			apiContext.l.Println("[ERROR] validating price data", errs)

			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}

		// add the rating to the context
		ctx := context.WithValue(r.Context(), KeyDetail{}, detprice)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
