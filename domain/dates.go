package domain

import "time"

type MatchingTimestamps struct {
	Data []time.Time `json:"data"`
}

type Error struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}

type TimestampsService interface {
	GetMatchingTimestamps(period, tmp1, tmp2, loc string) ([]time.Time, error)
}

type TimestampsRepository interface {
	Hour(period, tmp1, tmp2, loc string) ([]time.Time, error)
	Day() error
	Month() error
	Year() error
}
