package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"sort"

	"github.com/emicklei/go-restful"
)

type CustomerVisit struct {
	VisitId          int       `json:"visit_id"`
	CustomerName     string    `json:"customer_name"`
	CustomerIndustry string    `json:"customer_industry"`
	VisitStart       time.Time `json:"visit_start"`
	VisitEnd         time.Time `json:"visit_end"`
}

type CustomerVisits []CustomerVisit

const iso8601 = "2006-01-02T15:04:05Z0700"

func (s *Schedule) GetVisitors(request *restful.Request, response *restful.Response) {
	startDateStr := request.QueryParameter("start_date")
	endDateStr := request.QueryParameter("end_date")
	if startDateStr == "" || endDateStr == "" {
		response.WriteEntity(s.CustomerVisits)
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
	visits := make(CustomerVisits, 0)
	for _, visit := range s.CustomerVisits {
		if timeRangeIntersect(startDate, endDate, visit.VisitStart, visit.VisitEnd) {
			visits = append(visits, visit)
		}
	}
	response.WriteEntity(visits)
}

func (s *Schedule) GetVisitor(request *restful.Request, response *restful.Response) {
	visitorIdStr := request.PathParameter("visitor-id")
	visitorId, err := strconv.Atoi(visitorIdStr)
	if err != nil {
		response.WriteErrorString(http.StatusBadRequest, fmt.Sprintf("invalid visitor id. %v", err))
		return
	}

	visitor, err := s.CustomerVisits.findVisitor(visitorId)
	if err != nil {
		response.WriteErrorString(http.StatusNotFound, err.Error())
		return
	}

	response.WriteEntity(visitor)
}

func (s *Schedule) GetVisitorAgenda(request *restful.Request, response *restful.Response) {
	visitorIdStr := request.PathParameter("visitor-id")
	visitorId, err := strconv.Atoi(visitorIdStr)
	if err != nil {
		response.WriteErrorString(http.StatusBadRequest, fmt.Sprintf("invalid visitor id. %v", err))
		return
	}

	visitor, err := s.CustomerVisits.findVisitor(visitorId)
	if err != nil {
		response.WriteErrorString(http.StatusNotFound, err.Error())
		return
	}

	events := s.VisitEvents.GetEventsForRange(visitor.VisitStart, visitor.VisitEnd)

	response.WriteEntity(events)
}

func (v *CustomerVisits) findVisitor(id int) (*CustomerVisit, error) {
	index := sort.Search(len(*v), func(i int) bool {
		return (*v)[i].VisitId >= id
	})
	if index == len(*v) || (*v)[index].VisitId != id {
		return nil, fmt.Errorf("failed to find visitor with id %v", id)
	}
	return &(*v)[index], nil
}
