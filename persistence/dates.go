package persistence

import (
	"time"
)

type TimestampRepository struct{}

func NewTimestampsRepository() *TimestampRepository {
	return &TimestampRepository{}
}

func (tr TimestampRepository) Hour(period, tmp1, tmp2, loc string) ([]string, error) {
	layout := "20060102T150405Z"
	// set timezone to input parameter loc
	location, err := time.LoadLocation(loc)
	if err != nil {
		return nil, err
	}
	// parse the timestamps
	t1, err := time.ParseInLocation(layout, tmp1, location)
	if err != nil {
		return nil, err
	}
	t2, err := time.ParseInLocation(layout, tmp2, location)
	if err != nil {
		return nil, err
	}
	// initialize a slice to store the timestamps
	var timestamps []string
	// iterate over the time period of one hour
	hour, _ := time.ParseDuration(period)
	for t := t1; t.Before(t2); t = t.Add(hour) {
		str := t.Format(time.RFC3339)
		timestamps = append(timestamps, str)
	}
	return timestamps, nil
}

func (tr TimestampRepository) Day(period, tmp1, tmp2, loc string) ([]string, error) {
	layout := "20060102T150405Z"
	// set timezone to input parameter loc
	location, err := time.LoadLocation(loc)
	if err != nil {
		return nil, err
	}
	// parse the timestamps
	t1, err := time.ParseInLocation(layout, tmp1, location)
	if err != nil {
		return nil, err
	}
	t2, err := time.ParseInLocation(layout, tmp2, location)
	if err != nil {
		return nil, err
	}
	// initialize a slice to store the timestamps
	var timestamps []string
	// Iterate over the time period of one day
	for t := t1; t.Before(t2); {
		str := t.Format(time.RFC3339)
		timestamps = append(timestamps, str)
		t = t.AddDate(0, 0, 1)
	}
	return timestamps, nil
}

func (tr TimestampRepository) Month(period, tmp1, tmp2, loc string) ([]string, error) {
	layout := "20060102T150405Z"
	// set timezone to input parameter loc
	location, err := time.LoadLocation(loc)
	if err != nil {
		return nil, err
	}
	// parse the timestamps
	t1, err := time.ParseInLocation(layout, tmp1, location)
	if err != nil {
		return nil, err
	}
	t2, err := time.ParseInLocation(layout, tmp2, location)
	if err != nil {
		return nil, err
	}
	// initialize a slice to store the timestamps
	var timestamps []string
	// iterate over the time period of one month
	month := time.Duration(30 * 24 * time.Hour)
	for t := t1; t.Before(t2); t = t.Add(month) {
		str := t.Format(time.RFC3339)
		timestamps = append(timestamps, str)
	}
	return timestamps, nil
}

func (tr TimestampRepository) Year(period, tmp1, tmp2, loc string) ([]string, error) {
	layout := "20060102T150405Z"
	// set timezone to input parameter loc
	location, err := time.LoadLocation(loc)
	if err != nil {
		return nil, err
	}
	// parse the timestamps
	t1, err := time.ParseInLocation(layout, tmp1, location)
	if err != nil {
		return nil, err
	}
	t2, err := time.ParseInLocation(layout, tmp2, location)
	if err != nil {
		return nil, err
	}
	// initialize a slice to store the timestamps
	var timestamps []string
	// Iterate over the time period of one day
	for t := t1; t.Before(t2); {
		str := t.Format(time.RFC3339)
		timestamps = append(timestamps, str)
		t = time.Date(t.Year()+1, 1, 1, 0, 0, 0, 0, location)
	}
	return timestamps, nil
}
