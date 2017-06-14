package main

import (
	"fmt"
	"github.com/goadesign/goa"
	"goa-demo/app"
)

// BottleController implements the bottle resource.
type BottleController struct {
	*goa.Controller
}

// NewBottleController creates a bottle controller.
func NewBottleController(service *goa.Service) *BottleController {
	return &BottleController{Controller: service.NewController("BottleController")}
}

// Show runs the show action.
func (c *BottleController) Show(ctx *app.ShowBottleContext) error {
	// BottleController_Show: start_implement
	// Put your logic here
	if ctx.BottleID == 0 {
		return ctx.NotFound()
	}

	bottle := app.GoaExampleBottle{
		ID:   ctx.BottleID,
		Name: fmt.Sprintf("Bottle #%d", ctx.BottleID),
		Href: app.BottleHref(ctx.BottleID),
	}

	return ctx.OK(&bottle)
}
