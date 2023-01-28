package services

import (
	"errors"
	"time"
	"timestamps/domain"
)

type TimestampService struct {
	tr domain.TimestampsRepository
}

func NewTimestampsService(tr domain.TimestampsRepository) *TimestampService {
	return &TimestampService{tr: tr}
}

func (ts TimestampService) GetMatchingTimestamps(period, tmp1, tmp2, loc string) ([]time.Time, error) {
	switch period {
	case "1h":
		timestamps, err := ts.tr.Hour(period, tmp1, tmp2, loc)
		if err != nil {
			return nil, err
		}
		return timestamps, nil
	case "1d":
		return nil, nil
	case "1mo":
		return nil, nil
	case "1y":
		return nil, nil
	default:
		return nil, errors.New("invalid period")
	}
}
