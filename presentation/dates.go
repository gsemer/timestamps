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
	first, found2 := request.URL.Query()["t1"]
	second, found3 := request.URL.Query()["t2"]
	tz, found4 := request.URL.Query()["tz"]
	var period, t1, t2, location string
	if found1 && found2 && found3 && found4 {
		period = p[0]
		t1 = first[0]
		t2 = second[0]
		location = tz[0]
	} else {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("make sure that you use all required query parameters"))
		return
	}
	timestamps, err := th.ts.GetMatchingTimestamps(period, t1, t2, location, constants.UTC)
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
