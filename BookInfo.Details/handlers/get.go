package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"bookinfo/details/data"
	"bookinfo/details/dto"
	"bookinfo/details/logger"

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
	spanname := "Details.ListSingle"
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


	detail, err := data.GetDetailByID(id, ctx.MongoClient, ctx.DatabaseName)
	if err != nil {
		logger.Log("Error getting Detail", logger.ErrorLevel, err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
	// Call stocks service
	url := os.Getenv("STOCK_URL")
	if url == "" {
		url = "http://localhost:5114"
	}

	url = url + "/stocks/" + strconv.Itoa(id)
	// First prepare the tracing info
	netClient := &http.Client{Timeout: time.Second * 10}
	req, _ := http.NewRequest("GET", url, nil)
	// Inject the client span context into the headers
	tracer.Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
	stockresponse, err := netClient.Do(req)

	stockInfo := &dto.StockInfo{
		CurrentStock: 0,
	}
	if err == nil {
		buf, _ := ioutil.ReadAll(stockresponse.Body)
		json.Unmarshal(buf, &stockInfo)
	}
	detail.CurrentStock = stockInfo.CurrentStock
	err = data.ToJSON(detail, rw)
	if err != nil {
		// we should never be here but log the error just incase
		logger.Log("Error serializing Rating", logger.ErrorLevel, err)
	}
}
