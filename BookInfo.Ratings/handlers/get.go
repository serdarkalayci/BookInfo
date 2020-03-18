package handlers

import (
	"net/http"

	"bookinfo/ratings/data"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// swagger:route GET /Ratings/{id} Ratings listSingleRating
// Return a list of Ratings from the database
// responses:
//	200: RatingResponse
//	404: errorResponse

// ListSingle handles GET requests
func (p *APIContext) ListSingle(rw http.ResponseWriter, r *http.Request) {
	tracer := opentracing.GlobalTracer()
	spanname := "Ratings.ListSingle"
	var span opentracing.Span

	wireContext, err := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header))
	if err != nil {
		// The method is called without a span context in the http header.
		//
		span = tracer.StartSpan(spanname)
	} else {
		// Create the span referring to the RPC client if available.
		// If wireContext == nil, a root span will be created.
		span = opentracing.StartSpan(spanname, ext.RPCServerOption(wireContext))
	}
	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, r.URL.RequestURI())
	ext.HTTPMethod.Set(span, r.Method)
	defer span.Finish()

	id := getBookID(r)

	p.l.Println("[DEBUG] get record id", id)

	prod := data.GetRatingByID(id)

	err = data.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing Rating", err)
	}
}
