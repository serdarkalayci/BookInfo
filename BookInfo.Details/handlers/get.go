package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"bookinfo/details/data"
	"bookinfo/details/dto"
)

// swagger:route GET /Ratings/{id} Ratings listSingleRating
// Return a list of Ratings from the database
// responses:
//	200: RatingResponse
//	404: errorResponse

// ListSingle handles GET requests
func (p *APIContext) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := getBookID(r)

	p.l.Println("[DEBUG] get record id", id)

	detail, err := data.GetDetailByID(id)
	if err != nil {
		p.l.Println("[ERROR] fetching book detail", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
	// Call stocks service
	netClient := &http.Client{Timeout: time.Second * 10}
	stockresponse, err := netClient.Get("http://localhost:5114/stocks/1")
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
		p.l.Println("[ERROR] serializing Rating", err)
	}
}
