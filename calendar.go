package googlewrapper

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/oauth2"
	calendar "google.golang.org/api/calendar/v3"
)

func AddAttendee(srv *calendar.Service, calID string, event *calendar.Event, email string) {
	//events, err := srv.Events.List("primary").ShowDeleted(false).TimeMin(t).MaxResults(10).Do()

	event.Attendees = append(event.Attendees, &calendar.EventAttendee{
		Email: email,
	})

	events, err := srv.Events.Patch(calID, event.Id, event).Do()
	fmt.Println(events)
	//events, err := srv.Events.List("primary").ShowDeleted(false).SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("update err: %v", err)
	}
}

func ChangeOrganizer(srv *calendar.Service, calID string, event *calendar.Event, email, username string) {
	//events, err := srv.Events.List("primary").ShowDeleted(false).TimeMin(t).MaxResults(10).Do()
	newOrg := &calendar.EventOrganizer{
		DisplayName: username,
		Email:       email,
	}
	event.Organizer = newOrg

	events, err := srv.Events.Patch(calID, event.Id, event).Do()
	fmt.Println(events)
	//events, err := srv.Events.List("primary").ShowDeleted(false).SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("update err: %v", err)
	}
}

func UpdateDesc(srv *calendar.Service, calID string, eventID string, desc string) {
	//events, err := srv.Events.List("primary").ShowDeleted(false).TimeMin(t).MaxResults(10).Do()
	event := calendar.Event{
		Description: desc,
	}
	events, err := srv.Events.Patch(calID, eventID, &event).Do()
	fmt.Println(events)
	//events, err := srv.Events.List("primary").ShowDeleted(false).SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("update err: %v", err)
	}

}

func GetEvent(srv *calendar.Service, calendarID, eventID string) *calendar.Event {
	//events, err := srv.Events.List("primary").ShowDeleted(false).TimeMin(t).MaxResults(10).Do()
	event, err := srv.Events.Get(calendarID, eventID).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve the event: %v", err)
	}
	return event
}

func Watch(srv *calendar.Service) *calendar.Channel {
	newChan := calendar.Channel{}
	//events, err := srv.Events.List("primary").ShowDeleted(false).TimeMin(t).MaxResults(10).Do()
	chanx, err := srv.Events.Watch("primary", &newChan).Do()
	if err != nil {
		panic(err)
	}
	return chanx
}

func GetEvents(srv *calendar.Service, startDate time.Time, endDate time.Time, googleCalID string) *calendar.Events {
	//events, err := srv.Events.List("primary").ShowDeleted(false).TimeMin(t).MaxResults(10).Do()
	events, err := srv.Events.List(googleCalID).
		ShowDeleted(false).
		SingleEvents(true).
		TimeMin(startDate.Format(time.RFC3339)).
		TimeMax(endDate.Format(time.RFC3339)).
		OrderBy("startTime").Do()

	//events, err := srv.Events.List("primary").ShowDeleted(false).SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	return events
}

func GetCalendar(config *oauth2.Config, calendarToken string) *calendar.Service {

	srv, err := calendar.New(GetClient(config, calendarToken))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	return srv
}
