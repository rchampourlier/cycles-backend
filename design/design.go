package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("cycles", func() {
	Title("Project management for teams working with cycles")
	Description("The sur-mesure project management tool for teams working with cycles")
	Scheme("http")
	Host("localhost:8081")
	Consumes("application/json")
	Produces("application/json")
})

var _ = Resource("state", func() {
	BasePath("/states")
	DefaultMedia(StateMedia)

	Action("show", func() {
		Description("Get latest state")
		Routing(GET("/latest"))
		Response(OK)
		Response(NotFound)
	})

	Action("create", func() {
		Description("Create a new state")
		Routing(POST("/"))
		Payload(StatePayload, func() {
			Required("state")
		})
		Response(Created)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Origin("http://localhost:8080", func() {
		Headers("Content-Type")
	})
})

var StateMedia = MediaType("application/cycles.state+json", func() {
	Description("State of the Cycles frontend application")
	Reference(StatePayload)
	Attributes(func() {
		Attribute("state", String, "JSON state of the application")
	})

	View("default", func() {
		Attribute("state")
	})
})

var StatePayload = Type("StatePayload", func() {
	Attribute("state", func() {
	})
})
