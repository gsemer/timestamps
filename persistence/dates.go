package persistence

import (
	"time"
)

type TimestampRepository struct{}

func NewTimestampsRepository() *TimestampRepository {
	return &TimestampRepository{}
}

func ConvertTimeToString(t time.Time, layout string) (string, error) {
	timeToStr := t.Format(layout)
	parsed, err := time.Parse(layout, timeToStr)
	if err != nil {
		return "", err
	}
	parsed = parsed.UTC()
	formatted := t.Format(layout)
	return formatted, nil
}

func (tr TimestampRepository) Hour(period string, t1, t2 time.Time, location *time.Location, layout string) ([]string, error) {
	var timestamps []string
	hour, _ := time.ParseDuration(period)
	for t := t1; t.Before(t2); t = t.Add(hour) {
		t = t.Round(time.Hour)
		formatted, err := ConvertTimeToString(t, layout)
		if err != nil {
			return nil, err
		}
		timestamps = append(timestamps, formatted)
	}
	return timestamps, nil
}

func (tr TimestampRepository) Day(period string, t1, t2 time.Time, location *time.Location, layout string) ([]string, error) {
	var timestamps []string
	for t := t1; t.Before(t2); t = t.AddDate(0, 0, 1) {
		t = t.Round(time.Hour)
		formatted, err := ConvertTimeToString(t, layout)
		if err != nil {
			return nil, err
		}
		timestamps = append(timestamps, formatted)
	}
	return timestamps, nil
}

func (tr TimestampRepository) Month(period string, t1, t2 time.Time, location *time.Location, layout string) ([]string, error) {
	var timestamps []string
	for t := t1; t.Before(t2); t = t.AddDate(0, 1, 0) {
		if t.Year() != t2.Year() || t.Month() != t2.Month() {
			t = t.Round(time.Hour)
			formatted, err := ConvertTimeToString(t, layout)
			if err != nil {
				return nil, err
			}
			timestamps = append(timestamps, formatted)
		}
	}
	return timestamps, nil
}

func (tr TimestampRepository) Year(period string, t1, t2 time.Time, location *time.Location, layout string) ([]string, error) {
	var timestamps []string
	for t := t1; t.Before(t2); t = t.AddDate(1, 0, 0) {
		if t.Year() != t2.Year() {
			t = t.Round(time.Hour)
			formatted, err := ConvertTimeToString(t, layout)
			if err != nil {
				return nil, err
			}
			timestamps = append(timestamps, formatted)
		}
	}
	return timestamps, nil
}
