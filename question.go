package main

import (
	"github.com/goadesign/goa"
	"github.com/studiously/core/app"
)

// QuestionController implements the question resource.
type QuestionController struct {
	*goa.Controller
}

// NewQuestionController creates a question controller.
func NewQuestionController(service *goa.Service) *QuestionController {
	return &QuestionController{Controller: service.NewController("QuestionController")}
}

// Show runs the show action.
func (c *QuestionController) Show(ctx *app.ShowQuestionContext) error {
	// QuestionController_Show: start_implement

	// Put your logic here

	// QuestionController_Show: end_implement
	res := &app.StudiouslyQuestion{}
	return ctx.OK(res)
}
