package models

import "time"

type Event struct {
	ID          int
	Name        string `binding:"required"` //for generating the error if this is missing
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserId      int
}

var events = []Event{}

// Save adds the event to the slice
func (e Event) Save() {
	events = append(events, e)
}

// GetAllEvents retrieves all events
func GetAllEvents() []Event {
	return events
}
