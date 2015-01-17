package components

import (
	"fmt"
)

type Calendar struct {
	Version string `ical:",2.0"`
	ProdId  string `ical:",-//taviti/caldav-go//NONSGML v1.0.0//EN"`
	Method  types.Method
	*Event
}

func (c *Calendar) AddEvent(event *Event) error {

	if event.DateStart == nil && c.Method == "" {
		return fmt.Errorf("event date start is required for calendars without a method")
	}

	c.Event = event
	return nil

}
