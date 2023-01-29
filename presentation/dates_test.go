package presentation

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"timestamps/app/fakes"
	"timestamps/domain"
)

func TestTimestampHandler_GetMatchingTimestamps_Year(t *testing.T) {
	ts := &fakes.FakeTimestampsService{}
	matchingTimestamps := domain.MatchingTimestamps{
		Data: []string{
			"2018-02-14T20:46:03+02:00",
			"2019-01-01T00:00:00+02:00",
			"2020-01-01T00:00:00+02:00",
			"2021-01-01T00:00:00+02:00",
		},
	}
	ts.GetMatchingTimestampsReturns(matchingTimestamps.Data, nil)

	myHandlerFunc := TimestampHandler{ts: ts}

	values := url.Values{}
	values.Set("period", "1y")
	values.Set("tz", "Europe/Athens")
	values.Set("t1", "20180214T204603Z")
	values.Set("t2", "20211115T123456Z")
	request, _ := http.NewRequest("GET", "/ptlist?"+values.Encode(), nil)
	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/ptlist", myHandlerFunc.GetMatchingTimestamps)
	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("[ERROR] Status code is not %v, but is %v", http.StatusOK, response.Code)
		return
	}

	matchingTimestampsList, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
		return
	}

	result := MatchingTimestampsResponse{}
	err = json.Unmarshal(matchingTimestampsList, &result)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(result.Data, matchingTimestamps.Data) {
		t.Error("[ERROR] Not the expected output")
		return
	}
}

func TestTimestampHandler_GetMatchingTimestamps_Month(t *testing.T) {
	ts := &fakes.FakeTimestampsService{}
	matchingTimestamps := domain.MatchingTimestamps{
		Data: []string{
			"2021-02-14T20:46:03+02:00",
			"2021-03-16T20:46:03+02:00",
			"2021-04-15T21:46:03+03:00",
			"2021-05-15T21:46:03+03:00",
			"2021-06-14T21:46:03+03:00",
			"2021-07-14T21:46:03+03:00",
			"2021-08-13T21:46:03+03:00",
			"2021-09-12T21:46:03+03:00",
			"2021-10-12T21:46:03+03:00",
			"2021-11-11T20:46:03+02:00",
		},
	}
	ts.GetMatchingTimestampsReturns(matchingTimestamps.Data, nil)

	myHandlerFunc := TimestampHandler{ts: ts}

	values := url.Values{}
	values.Set("period", "1mo")
	values.Set("tz", "Europe/Athens")
	values.Set("t1", "20210214T204603Z")
	values.Set("t2", "20211115T123456Z")
	request, _ := http.NewRequest("GET", "/ptlist?"+values.Encode(), nil)
	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/ptlist", myHandlerFunc.GetMatchingTimestamps)
	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("[ERROR] Status code is not %v, but is %v", http.StatusOK, response.Code)
		return
	}

	matchingTimestampsList, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
		return
	}

	result := MatchingTimestampsResponse{}
	err = json.Unmarshal(matchingTimestampsList, &result)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(result.Data, matchingTimestamps.Data) {
		t.Error("[ERROR] Not the expected output")
		return
	}
}

func TestTimestampHandler_GetMatchingTimestamps_Day(t *testing.T) {
	ts := &fakes.FakeTimestampsService{}
	matchingTimestamps := domain.MatchingTimestamps{
		Data: []string{
			"2021-10-10T20:46:03+03:00",
			"2021-10-11T20:46:03+03:00",
			"2021-10-12T20:46:03+03:00",
			"2021-10-13T20:46:03+03:00",
			"2021-10-14T20:46:03+03:00",
			"2021-10-15T20:46:03+03:00",
			"2021-10-16T20:46:03+03:00",
			"2021-10-17T20:46:03+03:00",
			"2021-10-18T20:46:03+03:00",
			"2021-10-19T20:46:03+03:00",
			"2021-10-20T20:46:03+03:00",
			"2021-10-21T20:46:03+03:00",
			"2021-10-22T20:46:03+03:00",
			"2021-10-23T20:46:03+03:00",
			"2021-10-24T20:46:03+03:00",
			"2021-10-25T20:46:03+03:00",
			"2021-10-26T20:46:03+03:00",
			"2021-10-27T20:46:03+03:00",
			"2021-10-28T20:46:03+03:00",
			"2021-10-29T20:46:03+03:00",
			"2021-10-30T20:46:03+03:00",
			"2021-10-31T20:46:03+02:00",
			"2021-11-01T20:46:03+02:00",
			"2021-11-02T20:46:03+02:00",
			"2021-11-03T20:46:03+02:00",
			"2021-11-04T20:46:03+02:00",
			"2021-11-05T20:46:03+02:00",
			"2021-11-06T20:46:03+02:00",
			"2021-11-07T20:46:03+02:00",
			"2021-11-08T20:46:03+02:00",
			"2021-11-09T20:46:03+02:00",
			"2021-11-10T20:46:03+02:00",
			"2021-11-11T20:46:03+02:00",
			"2021-11-12T20:46:03+02:00",
			"2021-11-13T20:46:03+02:00",
			"2021-11-14T20:46:03+02:00",
		},
	}
	ts.GetMatchingTimestampsReturns(matchingTimestamps.Data, nil)

	myHandlerFunc := TimestampHandler{ts: ts}

	values := url.Values{}
	values.Set("period", "1d")
	values.Set("tz", "Europe/Athens")
	values.Set("t1", "20211010T204603Z")
	values.Set("t2", "20211115T123456Z")
	request, _ := http.NewRequest("GET", "/ptlist?"+values.Encode(), nil)
	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/ptlist", myHandlerFunc.GetMatchingTimestamps)
	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("[ERROR] Status code is not %v, but is %v", http.StatusOK, response.Code)
		return
	}

	matchingTimestampsList, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
		return
	}

	result := MatchingTimestampsResponse{}
	err = json.Unmarshal(matchingTimestampsList, &result)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(result.Data, matchingTimestamps.Data) {
		t.Error("[ERROR] Not the expected output")
		return
	}
}

func TestTimestampHandler_GetMatchingTimestamps_Hour(t *testing.T) {
	ts := &fakes.FakeTimestampsService{}
	matchingTimestamps := domain.MatchingTimestamps{
		Data: []string{
			"2021-07-14T20:46:03+03:00",
			"2021-07-14T21:46:03+03:00",
			"2021-07-14T22:46:03+03:00",
			"2021-07-14T23:46:03+03:00",
			"2021-07-15T00:46:03+03:00",
			"2021-07-15T01:46:03+03:00",
			"2021-07-15T02:46:03+03:00",
			"2021-07-15T03:46:03+03:00",
			"2021-07-15T04:46:03+03:00",
			"2021-07-15T05:46:03+03:00",
			"2021-07-15T06:46:03+03:00",
			"2021-07-15T07:46:03+03:00",
			"2021-07-15T08:46:03+03:00",
			"2021-07-15T09:46:03+03:00",
			"2021-07-15T10:46:03+03:00",
			"2021-07-15T11:46:03+03:00",
		},
	}
	ts.GetMatchingTimestampsReturns(matchingTimestamps.Data, nil)

	myHandlerFunc := TimestampHandler{ts: ts}

	values := url.Values{}
	values.Set("period", "1h")
	values.Set("tz", "Europe/Athens")
	values.Set("t1", "20210714T204603Z")
	values.Set("t2", "20210715T123456Z")
	request, _ := http.NewRequest("GET", "/ptlist?"+values.Encode(), nil)
	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/ptlist", myHandlerFunc.GetMatchingTimestamps)
	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("[ERROR] Status code is not %v, but is %v", http.StatusOK, response.Code)
		return
	}

	matchingTimestampsList, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
		return
	}

	result := MatchingTimestampsResponse{}
	err = json.Unmarshal(matchingTimestampsList, &result)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(result.Data, matchingTimestamps.Data) {
		t.Error("[ERROR] Not the expected output")
		return
	}
}

func TestTimestampHandler_GetMatchingTimestamps_FAIL(t *testing.T) {
	ts := &fakes.FakeTimestampsService{}
	ts.GetMatchingTimestampsReturns(nil, errors.New("something went wrong"))

	myHandlerFunc := TimestampHandler{ts: ts}

	server := httptest.NewServer(http.HandlerFunc(myHandlerFunc.GetMatchingTimestamps))
	defer server.Close()

	res, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
		return
	}
	res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected %v but got %v", http.StatusBadRequest, res.StatusCode)
	}
}
