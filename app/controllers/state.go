package controllers

import (
	"time"

	"github.com/goadesign/goa"
	"github.com/rchampourlier/cycles-backend/app"
	"github.com/rchampourlier/cycles-backend/commands"
	"github.com/rchampourlier/cycles-backend/queries"
)

// StateController implements the state resource.
type StateController struct {
	*goa.Controller
	*commands.StateCommands
	*queries.StateQueries
}

// NewStateController creates a state controller.
func NewStateController(service *goa.Service, commands *commands.StateCommands, queries *queries.StateQueries) *StateController {
	return &StateController{
		Controller:    service.NewController("StateController"),
		StateCommands: commands,
		StateQueries:  queries,
	}
}

// Create runs the create action.
func (c *StateController) Create(ctx *app.CreateStateContext) error {
	// StateController_Create: start_implement

	_, err := c.StoreState(time.Now(), ctx.Payload.State)
	if err != nil {
		return ctx.InternalServerError(err)
	}
	// TODO: should validate that State is a valid JSON
	//   and return a `BadRequest` response.

	return ctx.Created()
	// StateController_Create: end_implement
}

// Show runs the show action.
func (c *StateController) Show(ctx *app.ShowStateContext) error {
	// StateController_Show: start_implement

	// TODO: handle the internal server error case
	state, err := c.GetLatestState()
	if err != nil {
		return ctx.NotFound()
	}

	res := &app.CyclesState{
		State: &state,
	}

	return ctx.OK(res)
	// StateController_Show: end_implement
}
