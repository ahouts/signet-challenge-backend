package main

import (
	"encoding/json"
	"os"
	"sort"
)

type Schedule struct {
	CustomerVisits CustomerVisits `json:"customer_visits"`
	VisitEvents    VisitEvents    `json:"visit_events"`
}

func NewSchedule(filename string) (*Schedule, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	schedule := Schedule{}
	jsonParser := json.NewDecoder(f)
	if err = jsonParser.Decode(&schedule); err != nil {
		return nil, err
	}

	// sort customer visits by ID for searching later
	sort.Slice(schedule.CustomerVisits, func(i, j int) bool {
		return schedule.CustomerVisits[i].VisitId < schedule.CustomerVisits[j].VisitId
	})

	// sort schedules by ID for searching later
	sort.Slice(schedule.VisitEvents, func(i, j int) bool {
		return schedule.VisitEvents[i].VisitId < schedule.VisitEvents[j].VisitId
	})

	return &schedule, nil
}
