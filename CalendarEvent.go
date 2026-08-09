package main

import (
	"time"
)

type CalendarEventResponse struct {
	Embedded Embedded `json:"_embedded"`
}

type Embedded struct {
	Events []CalendarEvent `json:"osdi:events"`
}

type CalendarEvent struct {
	Identifiers  []string  `json:"identifiers"`
	Id           string    `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Instructions string    `json:"instructions"`
	CreatedDate  time.Time `json:"created_date"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	SponsorTitle string    `json:"action_network:sponsor.title"`
	Status       string    `json:"status"`
	Visibility   string    `json:"visibility"`
	EventLink    string    `json:"browser_url"`
}
