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

var _ = Resource("events", func() {
	BasePath("/events")
	DefaultMedia(StateMedia)

	Action("push", func() {
		Description("Push a new event")
		Routing(POST("/push"))
		Payload(EventPayload, func() {
			Required("event")
		})
		Response(OK)
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
		Attribute("state")
	})

	View("default", func() {
		Attribute("state")
	})
})

var StatePayload = Type("StatePayload", func() {
	Attribute("state", Any, "Application state", func() {
	})
})

var EventPayload = Type("EventPayload", func() {
	Attribute("event", Any, "Event", func() {
	})
})

var State = Type("State", func() {
	Attribute("roles", ArrayOf(Role), "Team members roles", func() {})
	Attribute("statuses", ArrayOf(Status), "Status for plans' status records", func() {})

})

var Role = Type("Role", func() {
	Attribute("id", String, "An unique identifier", func() {})
	Attribute("read", String, "The readable representation of the role", func() {})
	Attribute("icon", String, "The identifier of the icon to use to display this role", func() {})
})

var Status = Type("Status", func() {

})
