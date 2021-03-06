package design

import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("question", func() {
	BasePath("/questions")
	DefaultMedia(QuestionMedia)

	Action("show", func() {
		Description("Get a question by ID")
		Routing(GET("/:questionID"))
		Params(func() {
			Param("questionID", Integer, "Question ID")
		})
		Response(OK, QuestionMedia)
		Response(NotFound, func() {
			Description("Question not found")
			Status(404)
		})
		Response(InternalServerError)
	})
})

var QuestionMedia = MediaType("application/studiously.question+json", func() {
	// Does not return the class ID because strict isolation is maintained, thus it's irrelevant
	Attributes(func() {
		Attribute("id", Integer)
		Attribute("question_type", String)
		Attribute("author_id", Integer)
		Attribute("unit_id", Integer)
		Attribute("votes", Integer)
		Attribute("answered", Boolean)
		Attribute("data", Any)
		Required("id", "question_type", "unit_id")
	})
	View("default", func() {
		Attribute("id")
		Attribute("question_type")
		Attribute("author_id")
		Attribute("unit_id")
		Attribute("votes")
		Attribute("answered")
		Attribute("data")
	})
	View("by_unit", func() {
		Attribute("id")
		Attribute("question_type")
		Attribute("author_id")
		Attribute("votes")
		Attribute("answered")
		Attribute("data")
	})
	View("by_author", func() {
		Attribute("id")
		Attribute("question_type")
		Attribute("unit_id")
		Attribute("votes")
		Attribute("answered")
		Attribute("data")
	})
	View("by_type", func() {
		Attribute("id")
		Attribute("author_id")
		Attribute("unit_id")
		Attribute("votes")
		Attribute("answered")
		Attribute("data")
	})
	View("feed", func() {
		Attribute("id")
		Attribute("author_id")
		Attribute("unit_id")
		Attribute("votes")
		Attribute("answered")
	})
})
