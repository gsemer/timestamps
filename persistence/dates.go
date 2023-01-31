package persistence

import (
	"strconv"
	"time"
)

type TimestampRepository struct{}

func NewTimestampsRepository() *TimestampRepository {
	return &TimestampRepository{}
}

func ConvertTimeToString(t time.Time, layout string) (string, error) {
	str := t.Format(time.RFC3339)
	parsed, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return "", err
	}
	parsed = parsed.UTC()
	formatted := t.Format(layout)
	return formatted, nil
}

func (tr TimestampRepository) Hour(period string, t1, t2 time.Time, location *time.Location, layout string) ([]string, error) {
	// initialize a slice to store the timestamps
	var timestamps []string
	// iterate over the time period of num of hours
	hour, _ := time.ParseDuration(period)
	for t := t1; t.Before(t2); t = t.Add(hour) {
		formatted, err := ConvertTimeToString(t, layout)
		if err != nil {
			return nil, err
		}
		timestamps = append(timestamps, formatted)
	}
	return timestamps, nil
}

func (tr TimestampRepository) Day(period string, t1, t2 time.Time, location *time.Location, layout string) ([]string, error) {
	numOfDays := period[:len(period)-1]
	numOfDaysToInt, _ := strconv.Atoi(numOfDays)
	// initialize a slice to store the timestamps
	var timestamps []string
	// iterate over the time period of num of days
	for t := t1; t.Before(t2); {
		formatted, err := ConvertTimeToString(t, layout)
		if err != nil {
			return nil, err
		}
		timestamps = append(timestamps, formatted)
		t = t.AddDate(0, 0, numOfDaysToInt)
	}
	return timestamps, nil
}

func (tr TimestampRepository) Month(period string, t1, t2 time.Time, location *time.Location, layout string) ([]string, error) {
	numOfDays := period[:len(period)-2]
	numOfDaysToInt, _ := strconv.Atoi(numOfDays)
	// initialize a slice to store the timestamps
	var timestamps []string
	// iterate over the time period of num of months
	for t := t1; t.Before(t2); t = t.AddDate(0, numOfDaysToInt, 0) {
		formatted, err := ConvertTimeToString(t, layout)
		if err != nil {
			return nil, err
		}
		timestamps = append(timestamps, formatted)
	}
	return timestamps, nil
}

func (tr TimestampRepository) Year(period string, t1, t2 time.Time, location *time.Location, layout string) ([]string, error) {
	numOfDays := period[:len(period)-1]
	numOfDaysToInt, _ := strconv.Atoi(numOfDays)
	// initialize a slice to store the timestamps
	var timestamps []string
	// Iterate over the time period of num of years
	for t := t1; t.Before(t2); {
		formatted, err := ConvertTimeToString(t, layout)
		if err != nil {
			return nil, err
		}
		timestamps = append(timestamps, formatted)
		t = time.Date(t.Year()+numOfDaysToInt, 1, 1, 0, 0, 0, 0, location)
	}
	return timestamps, nil
}
