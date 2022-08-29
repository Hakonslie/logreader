package events

import (
	"fmt"
	"strings"
)

type EventHistory struct {
	count  int
	events map[int]Event
}

type Event struct {
	Name   string
	EventJ string
}

func Initialize() (e EventHistory) {
	e.events = make(map[int]Event)
	e.count = 0
	return
}

func (e *EventHistory) GiveHistory() {
	for i, k := range e.events {
		if strings.Contains(k.EventJ, "HoardDragon") {
			fmt.Printf("event %d: %s", i, k.EventJ)
		}
	}
}

func (e *EventHistory) RecordEvent(event Event) {
	e.events[e.count] = event
	e.count++

	fmt.Println(event)
}
