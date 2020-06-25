package handlers

import (
	"bookinfo/ratings/data"
	"bookinfo/ratings/logger"
	"net/http"
)

// swagger:route GET /health/live Live Live
// Return 200 if the api is up and running
// responses:
//	200: OK
//	404: errorResponse

// Live handles GET requests
func (ctx *APIContext) Live(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
}

// swagger:route GET /health/ready Ready Ready
// Return 200 if the api is up and running and connected to the database
// responses:
//	200: OK
//	404: errorResponse

// Ready handles GET requests
func (ctx *DBContext) Ready(rw http.ResponseWriter, r *http.Request) {
	err := data.GetHealth(ctx.RedisClient, ctx.DatabaseName)
	if err != nil {
		logger.Log("Error connecting to database", logger.ErrorLevel, err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}
