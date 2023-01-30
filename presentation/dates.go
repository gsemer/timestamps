package presentation

import (
	"encoding/json"
	"net/http"
	"timestamps/domain"
	"timestamps/domain/constants"
)

type TimestampHandler struct {
	ts domain.TimestampsService
}

func NewTimestampsHandler(ts domain.TimestampsService) *TimestampHandler {
	return &TimestampHandler{ts: ts}
}

type MatchingTimestampsResponse struct {
	Data []string `json:"data"`
}

func (th TimestampHandler) GetMatchingTimestamps(writer http.ResponseWriter, request *http.Request) {
	p, found1 := request.URL.Query()["period"]
	t1, found2 := request.URL.Query()["t1"]
	t2, found3 := request.URL.Query()["t2"]
	tz, found4 := request.URL.Query()["tz"]
	var period, tmp1, tmp2, loc string

	if found1 && found2 && found3 && found4 {
		period = p[0]
		tmp1 = t1[0]
		tmp2 = t2[0]
		loc = tz[0]
	} else {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("make sure that you use all required query parameters"))
		return
	}
	timestamps, err := th.ts.GetMatchingTimestamps(period, tmp1, tmp2, loc, constants.UTC)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err2 := domain.Error{
			Status:      "error",
			Description: err.Error(),
		}
		resultErr, _ := json.Marshal(err2)
		writer.Write(resultErr)
		return
	}

	writer.WriteHeader(http.StatusOK)
	matchingTimestamps := MatchingTimestampsResponse{Data: timestamps}
	result, _ := json.Marshal(matchingTimestamps)
	writer.Write(result)
	return
}
