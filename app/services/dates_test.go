package services

import (
	"reflect"
	"testing"
	"timestamps/persistence/fakes"
)

func TestTimestampService_GetMatchingTimestamps_Year(t *testing.T) {
	period := "1y"
	tz := "Europe/Athens"
	t1 := "20180214T204603Z"
	t2 := "20211115T123456Z"
	layout := "20060102T150405Z"

	mockMatchingTimestamps := []string{
		"20180214T210000Z",
		"20190214T210000Z",
		"20200214T210000Z",
	}

	tr := &fakes.FakeTimestampsRepository{}

	tr.YearReturns(mockMatchingTimestamps, nil)

	ts := TimestampService{tr: tr}
	timestamps, err := ts.GetMatchingTimestamps(period, t1, t2, tz, layout)
	if err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(timestamps, mockMatchingTimestamps) {
		t.Errorf("[ERROR] we expected to have %v as output,\n but we got %v", mockMatchingTimestamps, timestamps)
		return
	}
}

func TestTimestampService_GetMatchingTimestamps_Month(t *testing.T) {
	period := "1mo"
	tz := "Europe/Athens"
	t1 := "20210214T204603Z"
	t2 := "20211115T123456Z"
	layout := "20060102T150405Z"

	mockMatchingTimestamps := []string{
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

	tr := &fakes.FakeTimestampsRepository{}

	tr.MonthReturns(mockMatchingTimestamps, nil)

	ts := TimestampService{tr: tr}
	timestamps, err := ts.GetMatchingTimestamps(period, t1, t2, tz, layout)
	if err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(timestamps, mockMatchingTimestamps) {
		t.Errorf("[ERROR] we expected to have %v as output,\n but we got %v", mockMatchingTimestamps, timestamps)
		return
	}
}

func TestTimestampService_GetMatchingTimestamps_Day(t *testing.T) {
	period := "1d"
	tz := "Europe/Athens"
	t1 := "20211010T204603Z"
	t2 := "20211115T123456Z"
	layout := "20060102T150405Z"

	mockMatchingTimestamps := []string{
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

	tr := &fakes.FakeTimestampsRepository{}

	tr.DayReturns(mockMatchingTimestamps, nil)

	ts := TimestampService{tr: tr}
	timestamps, err := ts.GetMatchingTimestamps(period, t1, t2, tz, layout)
	if err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(timestamps, mockMatchingTimestamps) {
		t.Errorf("[ERROR] we expected to have %v as output,\n but we got %v", mockMatchingTimestamps, timestamps)
		return
	}
}

func TestTimestampService_GetMatchingTimestamps_Hour(t *testing.T) {
	period := "1h"
	tz := "Europe/Athens"
	t1 := "20210714T204603Z"
	t2 := "20210715T123456Z"
	layout := "20060102T150405Z"

	mockMatchingTimestamps := []string{
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

	tr := &fakes.FakeTimestampsRepository{}

	tr.HourReturns(mockMatchingTimestamps, nil)

	ts := TimestampService{tr: tr}
	timestamps, err := ts.GetMatchingTimestamps(period, t1, t2, tz, layout)
	if err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(timestamps, mockMatchingTimestamps) {
		t.Errorf("[ERROR] we expected to have %v as output,\n but we got %v", mockMatchingTimestamps, timestamps)
		return
	}
}

func TestTimestampService_GetMatchingTimestamps_FAIL_1(t *testing.T) {
	// THIS IS THE CASE WHERE t2 IS LESS THAN t1
	period := "1y"
	tz := "Europe/Athens"
	t1 := "20211115T123456Z"
	t2 := "20180214T204603Z"
	layout := "20060102T150405Z"

	tr := &fakes.FakeTimestampsRepository{}
	ts := TimestampService{tr: tr}

	_, err := ts.GetMatchingTimestamps(period, t1, t2, tz, layout)
	if err == nil {
		t.Error(err)
		return
	}
}

func TestTimestampService_GetMatchingTimestamps_FAIL_2(t *testing.T) {
	// THIS IS THE CASE WHERE WE GOT AN UNEXPECTED PERIOD
	period := "1w"
	tz := "Europe/Athens"
	t1 := "20180214T204603Z"
	t2 := "20211115T123456Z"
	layout := "20060102T150405Z"

	tr := &fakes.FakeTimestampsRepository{}
	ts := TimestampService{tr: tr}

	_, err := ts.GetMatchingTimestamps(period, t1, t2, tz, layout)
	if err == nil {
		t.Error(err)
		return
	}
}
