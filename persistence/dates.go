package persistence

import (
	"fmt"
	"time"
)

type TimestampRepository struct{}

func NewTimestampsRepository() *TimestampRepository {
	return &TimestampRepository{}
}

func (tr TimestampRepository) Hour(period, tmp1, tmp2, loc string) ([]time.Time, error) {
	layout := "20060102T150405Z"

	// Parse the timestamps
	t1, err := time.Parse(layout, tmp1)
	if err != nil {
		fmt.Println(err)
	}
	t2, err := time.Parse(layout, tmp2)
	if err != nil {
		fmt.Println(err)
	}

	// Initialize a slice to store the timestamps
	var timestamps []time.Time

	// Iterate over the time period of one hour
	hour, _ := time.ParseDuration(period)
	for t := t1; t.Before(t2); t = t.Add(hour) {
		timestamps = append(timestamps, t)
	}

	return timestamps, nil
}

func (tr TimestampRepository) Day() error {
	return nil
}

func (tr TimestampRepository) Month() error {
	return nil
}

func (tr TimestampRepository) Year() error {
	return nil
}
