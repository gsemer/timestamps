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
			"20180214T204603Z",
			"20190101T000000Z",
			"20200101T000000Z",
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
			"20210214T204603Z",
			"20210316T204603Z",
			"20210415T214603Z",
			"20210515T214603Z",
			"20210614T214603Z",
			"20210714T214603Z",
			"20210813T214603Z",
			"20210912T214603Z",
			"20211012T214603Z",
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
			"20211010T204603Z",
			"20211011T204603Z",
			"20211012T204603Z",
			"20211013T204603Z",
			"20211014T204603Z",
			"20211015T204603Z",
			"20211016T204603Z",
			"20211017T204603Z",
			"20211018T204603Z",
			"20211019T204603Z",
			"20211020T204603Z",
			"20211021T204603Z",
			"20211022T204603Z",
			"20211023T204603Z",
			"20211024T204603Z",
			"20211025T204603Z",
			"20211026T204603Z",
			"20211027T204603Z",
			"20211028T204603Z",
			"20211029T204603Z",
			"20211030T204603Z",
			"20211031T204603Z",
			"20211101T204603Z",
			"20211102T204603Z",
			"20211103T204603Z",
			"20211104T204603Z",
			"20211105T204603Z",
			"20211106T204603Z",
			"20211107T204603Z",
			"20211108T204603Z",
			"20211109T204603Z",
			"20211110T204603Z",
			"20211111T204603Z",
			"20211112T204603Z",
			"20211113T204603Z",
			"20211114T204603Z",
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
			"20210714T204603Z",
			"20210714T214603Z",
			"20210714T224603Z",
			"20210714T234603Z",
			"20210715T004603Z",
			"20210715T014603Z",
			"20210715T024603Z",
			"20210715T034603Z",
			"20210715T044603Z",
			"20210715T054603Z",
			"20210715T064603Z",
			"20210715T074603Z",
			"20210715T084603Z",
			"20210715T094603Z",
			"20210715T104603Z",
			"20210715T114603Z",
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
