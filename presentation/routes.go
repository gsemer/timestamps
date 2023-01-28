package presentation

import "timestamps/domain"

func CreateRoutes(service domain.TimestampsService) map[string]domain.RouteDefinition {
	th := NewTimestampsHandler(service)
	// Define a map
	return map[string]domain.RouteDefinition{
		"/ptlist": {
			Methods:     []string{"GET"},
			HandlerFunc: th.GetMatchingTimestamps,
		},
	}
}
