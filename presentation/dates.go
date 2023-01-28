package presentation

import (
	"encoding/json"
	"net/http"
	"timestamps/domain"
)

type TimestampHandler struct {
	ts domain.TimestampsService
}

func NewTimestampsHandler(ts domain.TimestampsService) *TimestampHandler {
	return &TimestampHandler{ts: ts}
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
	}

	timestamps, err := th.ts.GetMatchingTimestamps(period, tmp1, tmp2, loc)
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
	matchingTimestamps := domain.MatchingTimestamps{Data: timestamps}
	result, _ := json.Marshal(matchingTimestamps)
	writer.Write(result)
	return
}
