package domain

type MatchingTimestamps struct {
	Data []string `json:"data"`
}

type Error struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}

type TimestampsService interface {
	GetMatchingTimestamps(period, tmp1, tmp2, loc, layout string) ([]string, error)
}

type TimestampsRepository interface {
	Hour(period, tmp1, tmp2, loc, layout string) ([]string, error)
	Day(period, tmp1, tmp2, loc, layout string) ([]string, error)
	Month(period, tmp1, tmp2, loc, layout string) ([]string, error)
	Year(period, tmp1, tmp2, loc, layout string) ([]string, error)
}
