package services

import (
	"errors"
	"strings"
	"time"
	"timestamps/domain"
)

type TimestampService struct {
	tr domain.TimestampsRepository
}

func NewTimestampsService(tr domain.TimestampsRepository) *TimestampService {
	return &TimestampService{tr: tr}
}

func (ts TimestampService) GetMatchingTimestamps(period, tmp1, tmp2, loc, layout string) ([]string, error) {
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
	// check if t1 is after t2
	if t1.After(t2) {
		return nil, errors.New("invalid query parameters: t1 must be less than t2")
	}
	// handle every case separately based on the period
	switch {
	case strings.Contains(period, "h"):
		timestamps, err := ts.tr.Hour(period, t1, t2, location, layout)
		if err != nil {
			return nil, err
		}
		return timestamps, nil
	case strings.Contains(period, "d"):
		timestamps, err := ts.tr.Day(period, t1, t2, location, layout)
		if err != nil {
			return nil, err
		}
		return timestamps, nil
	case strings.Contains(period, "mo"):
		timestamps, err := ts.tr.Month(period, t1, t2, location, layout)
		if err != nil {
			return nil, err
		}
		return timestamps, nil
	case strings.Contains(period, "y"):
		timestamps, err := ts.tr.Year(period, t1, t2, location, layout)
		if err != nil {
			return nil, err
		}
		return timestamps, nil
	default:
		return nil, errors.New("unsupported period")
	}
}
