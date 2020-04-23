package handlers

import (
	"fmt"
	"net/http"

	"bookinfo/ratings/data"
	"bookinfo/ratings/logger"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// swagger:route GET /Ratings/{id} Ratings listSingleRating
// Return a list of Ratings from the database
// responses:
//	200: RatingResponse
//	404: errorResponse

// ListSingle handles GET requests
func (ctx *DBContext) ListSingle(rw http.ResponseWriter, r *http.Request) {
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

	logger.Log(fmt.Sprintf("get record id %d", id), logger.DebugLevel)

	rating, err := data.GetRatingByID(id, ctx.RedisClient, ctx.DatabaseName)
	if err != nil {
		logger.Log("Error getting Rating", logger.ErrorLevel, err)
	} else {
		logger.Log(fmt.Sprintf("Current Rating: %f", rating.CurrentRating), logger.DebugLevel)
	}
	err = data.ToJSON(rating, rw)
	if err != nil {
		// we should never be here but log the error just incase
		logger.Log("Error serializing Rating", logger.ErrorLevel, err)
	}
}
