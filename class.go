package main

import (
	"github.com/goadesign/goa"
	"github.com/studiously/core/app"
	"database/sql"
	"github.com/studiously/core/models"
)

// ClassController implements the class resource.
type ClassController struct {
	*goa.Controller
}

// NewClassController creates a class controller.
func NewClassController(service *goa.Service) *ClassController {
	return &ClassController{Controller: service.NewController("ClassController")}
}

// List runs the list action.
func (c *ClassController) List(ctx *app.ListClassContext) error {
	// ClassController_List: start_implement

	// Put your logic here

	// ClassController_List: end_implement
	res := app.StudiouslyClassCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *ClassController) Show(ctx *app.ShowClassContext) error {
	// ClassController_Show: start_implement

	// Put your logic here

	class, err := models.ClassByID(ctx.Value(DatabaseKey).(*sql.DB), int64(ctx.ClassID))
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	// ClassController_Show: end_implement
	res := &app.StudiouslyClass{
		ID: int(class.ID),
		Name: class.Name,
	}
	if class.CurrentUnit.Valid {
		cu := int(class.CurrentUnit.Int64)
		res.CurrentUnit = &cu
	}
	return ctx.OK(res)
}

// ShowMembers runs the show_members action.
func (c *ClassController) ShowMembers(ctx *app.ShowMembersClassContext) error {
	// ClassController_ShowMembers: start_implement

	// Put your logic here

	// ClassController_ShowMembers: end_implement
	res := app.StudiouslyMemberCollection{}
	return ctx.OK(res)
}

// ShowQuestions runs the show_questions action.
func (c *ClassController) ShowQuestions(ctx *app.ShowQuestionsClassContext) error {
	// ClassController_ShowQuestions: start_implement

	// Put your logic here

	// ClassController_ShowQuestions: end_implement
	res := app.StudiouslyQuestionFeedCollection{}
	return ctx.OKFeed(res)
}
