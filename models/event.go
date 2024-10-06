package models

import "time"

type Event struct {
	ID          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserID      int
}

var events []Event = []Event{}

func (event Event) Save() {
	events = append(events, event)
}

func GetAllEvents() []Event {
	return events
}
