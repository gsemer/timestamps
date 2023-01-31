package domain

import "time"

type MatchingTimestamps struct {
	Data []string `json:"data"`
}

type Error struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}

type TimestampsService interface {
	GetMatchingTimestamps(period, t1, t2, location, layout string) ([]string, error)
}

type TimestampsRepository interface {
	Hour(period string, t1, t2 time.Time, location *time.Location, layout string) ([]string, error)
	Day(period string, t1, t2 time.Time, location *time.Location, layout string) ([]string, error)
	Month(period string, t1, t2 time.Time, location *time.Location, layout string) ([]string, error)
	Year(period string, t1, t2 time.Time, location *time.Location, layout string) ([]string, error)
}
