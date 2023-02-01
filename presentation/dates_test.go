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
)

func TestTimestampHandler_GetMatchingTimestamps_Year(t *testing.T) {
	ts := &fakes.FakeTimestampsService{}
	matchingTimestamps := []string{
		"20180214T210000Z",
		"20190214T210000Z",
		"20200214T210000Z",
	}
	ts.GetMatchingTimestampsReturns(matchingTimestamps, nil)

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

	var result []string
	err = json.Unmarshal(matchingTimestampsList, &result)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(result, matchingTimestamps) {
		t.Error("[ERROR] Not the expected output")
		return
	}
}

func TestTimestampHandler_GetMatchingTimestamps_Month(t *testing.T) {
	ts := &fakes.FakeTimestampsService{}
	matchingTimestamps := []string{
		"20210214T210000Z",
		"20210314T210000Z",
		"20210414T210000Z",
		"20210514T210000Z",
		"20210614T210000Z",
		"20210714T210000Z",
		"20210814T210000Z",
		"20210914T210000Z",
		"20211014T210000Z",
	}
	ts.GetMatchingTimestampsReturns(matchingTimestamps, nil)

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

	var result []string
	err = json.Unmarshal(matchingTimestampsList, &result)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(result, matchingTimestamps) {
		t.Error("[ERROR] Not the expected output")
		return
	}
}

func TestTimestampHandler_GetMatchingTimestamps_Day(t *testing.T) {
	ts := &fakes.FakeTimestampsService{}
	matchingTimestamps := []string{
		"20211010T210000Z",
		"20211011T210000Z",
		"20211012T210000Z",
		"20211013T210000Z",
		"20211014T210000Z",
		"20211015T210000Z",
		"20211016T210000Z",
		"20211017T210000Z",
		"20211018T210000Z",
		"20211019T210000Z",
		"20211020T210000Z",
		"20211021T210000Z",
		"20211022T210000Z",
		"20211023T210000Z",
		"20211024T210000Z",
		"20211025T210000Z",
		"20211026T210000Z",
		"20211027T210000Z",
		"20211028T210000Z",
		"20211029T210000Z",
		"20211030T210000Z",
		"20211031T210000Z",
		"20211101T210000Z",
		"20211102T210000Z",
		"20211103T210000Z",
		"20211104T210000Z",
		"20211105T210000Z",
		"20211106T210000Z",
		"20211107T210000Z",
		"20211108T210000Z",
		"20211109T210000Z",
		"20211110T210000Z",
		"20211111T210000Z",
		"20211112T210000Z",
		"20211113T210000Z",
		"20211114T210000Z",
	}
	ts.GetMatchingTimestampsReturns(matchingTimestamps, nil)

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

	var result []string
	err = json.Unmarshal(matchingTimestampsList, &result)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(result, matchingTimestamps) {
		t.Error("[ERROR] Not the expected output")
		return
	}
}

func TestTimestampHandler_GetMatchingTimestamps_Hour(t *testing.T) {
	ts := &fakes.FakeTimestampsService{}
	matchingTimestamps := []string{
		"20210714T210000Z",
		"20210714T220000Z",
		"20210714T230000Z",
		"20210715T000000Z",
		"20210715T010000Z",
		"20210715T020000Z",
		"20210715T030000Z",
		"20210715T040000Z",
		"20210715T050000Z",
		"20210715T060000Z",
		"20210715T070000Z",
		"20210715T080000Z",
		"20210715T090000Z",
		"20210715T100000Z",
		"20210715T110000Z",
		"20210715T120000Z",
	}
	ts.GetMatchingTimestampsReturns(matchingTimestamps, nil)

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

	var result []string
	err = json.Unmarshal(matchingTimestampsList, &result)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(result, matchingTimestamps) {
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
