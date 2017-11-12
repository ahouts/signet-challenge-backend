package main

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-swagger12"
)

func setupRoutes(s *Schedule) {
	restful.Add(scheduleWs(s))
	config := swagger.Config{
		WebServices:     restful.RegisteredWebServices(),
		ApiPath:         "/apidocs.json",
		SwaggerPath:     "/apidocs/",
		SwaggerFilePath: "./swagger-dist"}
	swagger.InstallSwaggerService(config)
}

func scheduleWs(s *Schedule) *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/schedule").
		Doc("Schedule API Root").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/visitor").To(s.GetVisitors).
		Doc("Get info for all visitors").
		Param(ws.QueryParameter("start_date", "starting date filter by").DataType("string").DefaultValue("")).
		Param(ws.QueryParameter("end_date", "ending date filter by").DataType("string").DefaultValue("")).
		Writes(new(CustomerVisits)))

	ws.Route(ws.GET("/visitor/{visitor-id}").To(s.GetVisitor).
		Doc("Get visitor info").
		Param(ws.PathParameter("visitor-id", "Id of the visitor").DataType("int").DefaultValue("-1")).
		Writes(new(CustomerVisit)))

	ws.Route(ws.GET("/visitor/{visitor-id}/agenda").To(s.GetVisitorAgenda).
		Doc("Get visitor info").
		Param(ws.PathParameter("visitor-id", "Id of the visitor").DataType("int").DefaultValue("-1")).
		Writes(new(VisitEvents)))

	ws.Route(ws.GET("/event").To(s.GetEvents).
		Doc("Get info for all events").
		Param(ws.QueryParameter("start_date", "starting date filter by").DataType("string").DefaultValue("")).
		Param(ws.QueryParameter("end_date", "ending date filter by").DataType("string").DefaultValue("")).
		Writes(new(VisitEvents)))

	ws.Route(ws.GET("/event/{event-id}").To(s.GetEvent).
		Doc("Get event info").
		Param(ws.PathParameter("event-id", "Id of the event").DataType("int").DefaultValue("-1")).
		Writes(new(VisitEvent)))

	return ws
}
