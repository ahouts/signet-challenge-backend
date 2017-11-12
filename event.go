package main

import (
	"time"

	"fmt"
	"net/http"

	"sort"
	"strconv"

	"github.com/emicklei/go-restful"
)

type VisitEvent struct {
	VisitId      int       `json:"visit_id"`
	Event        string    `json:"event"`
	ShowOnAgenda bool      `json:"show_on_agenda"`
	EventStart   time.Time `json:"event_start"`
	EventEnd     time.Time `json:"event_end"`
}

type VisitEvents []VisitEvent

func (e *VisitEvents) GetEventsForRange(t1, t2 time.Time) VisitEvents {
	events := make(VisitEvents, 0)
	for _, event := range *e {
		if event.ShowOnAgenda && timeRangeIntersect(t1, t2, event.EventStart, event.EventEnd) {
			events = append(events, event)
		}
	}
	return events
}

func (s *Schedule) GetEvents(request *restful.Request, response *restful.Response) {
	startDateStr := request.QueryParameter("start_date")
	endDateStr := request.QueryParameter("end_date")
	if startDateStr == "" || endDateStr == "" {
		response.WriteEntity(s.VisitEvents)
		return
	}
	startDate, err := time.Parse(iso8601, startDateStr)
	if err != nil {
		response.WriteErrorString(http.StatusBadRequest, fmt.Sprintf("failed to parse starting date. %v", err))
		return
	}
	endDate, err := time.Parse(iso8601, endDateStr)
	if err != nil {
		response.WriteErrorString(http.StatusBadRequest, fmt.Sprintf("failed to parse ending date. %v", err))
		return
	}
	events := make(VisitEvents, 0)
	for _, event := range s.VisitEvents {
		if timeRangeIntersect(startDate, endDate, event.EventStart, event.EventEnd) {
			events = append(events, event)
		}
	}
	response.WriteEntity(events)
}

func (s *Schedule) GetEvent(request *restful.Request, response *restful.Response) {
	eventIdStr := request.PathParameter("event-id")
	eventId, err := strconv.Atoi(eventIdStr)
	if err != nil {
		response.WriteErrorString(http.StatusBadRequest, fmt.Sprintf("invalid event id. %v", err))
		return
	}

	event, err := s.VisitEvents.findEvent(eventId)
	if err != nil {
		response.WriteErrorString(http.StatusNotFound, err.Error())
		return
	}

	response.WriteEntity(event)
}

func (e *VisitEvents) findEvent(id int) (*VisitEvent, error) {
	index := sort.Search(len(*e), func(i int) bool {
		return (*e)[i].VisitId >= id
	})
	if index == len(*e) || (*e)[index].VisitId != id {
		return nil, fmt.Errorf("failed to find event with id %v", id)
	}
	return &(*e)[index], nil
}
