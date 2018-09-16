package main

import (
	"time"

	"github.com/goadesign/goa"
	"github.com/rchampourlier/cycles-backend/app"
	"github.com/rchampourlier/cycles-backend/commands"
	"github.com/rchampourlier/cycles-backend/queries"
)

// EventsController implements the state resource.
type EventsController struct {
	*goa.Controller
	*commands.EventsCommands
	*queries.StateQueries
}

// NewEventsController creates a state controller.
func NewEventsController(service *goa.Service, commands *commands.EventsCommands, queries *queries.StateQueries) *EventsController {
	return &EventsController{
		Controller:     service.NewController("EventsController"),
		EventsCommands: commands,
		StateQueries:   queries,
	}
}

// Push pushes a new event
func (c *EventsController) Push(ctx *app.PushEventsContext) error {
	// EventsController_Push: start_implement

	// Put your logic here
	_, err := c.PushEvent(time.Now(), ctx.Payload.Event)
	if err != nil {
		return ctx.InternalServerError(err)
	}
	// TODO: should validate that Events is a valid JSON
	//   and return a `BadRequest` response.

	state, err := c.GetLatestState()
	if err != nil {
		return ctx.InternalServerError(err)
	}

	res := &app.CyclesState{
		State: &state,
	}

	return ctx.OK(res)
	// EventsController_Create: end_implement
}
