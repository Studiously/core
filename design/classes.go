package design

import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("class", func() {
	BasePath("/classes")
	DefaultMedia(ClassMedia)
	Action("list", func() {
		Description("Get all classes a user is in")
		Routing(GET("/"))
		Response(OK, ArrayOf(ClassMedia))
	})
	Action("show", func() {
		Description("Get class by ID")
		Routing(GET("/:class_id"))
		Params(func() {
			Param("class_id", UUID, "Class ID")
		})
		Response(OK, ClassMedia)
		Response(NotFound, func() {
			Description("Specified class does not exist or the user does not have access to it")
			Status(404)
		})
	})
	Action("show_questions", func() {
		Description("Get questions for a class")
		Routing(GET("/:class_id/questions"))
		Params(func() {
			// Questions are scoped by class, so clients cannot retrieve all questions for all classes a user is in automatically. This is because studying multiple subjects at the same time
			Param("class_id", UUID, "Class ID")

			Param("question_type", String, "Filter by question type")
			Param("author_id", UUID, "Filter by author")
			Param("unit_id", UUID, "Filter by unit")
			Param("answered", Boolean, "Filter by whether the question has been answered by the user")
		})
		Response(OK, ArrayOf(FeedQuestionMedia))
		Response(NotFound, func() {
			Description("Specified class does not exist or the user does not have access to it")
			Status(404)
		})
		Response("UnknownUnit", func() {
			Description("Specified unit does not exist in the context of the class")
			Status(400)
		})
		Response("UnknownAuthor", func() {
			Description("Specified author does not exist in the context of the class")
			Status(400)
		})
		Response("UnknownType", func() {
			Description("Specified question type is not recognized")
		})
	})
})

var ClassMedia = MediaType("application/studiously.class+json", func() {
	Attributes(func() {
		Attribute("id", UUID)
		Attribute("email", String)
		Attribute("name", String)
		Required("email", "name")
	})
	View("default", func() {
		Attribute("id", UUID)
		Attribute("email")
		Attribute("name")
	})
})
var FeedQuestionMedia = MediaType("application/studiously.feed_question+json", func() {
	Attributes(func() {
		Attribute("id", UUID)
		Attribute("question_type", String, "Valid types include 'multiple_choice', 'true_false', and 'short_answer'.")
		Attribute("author_id", UUID)
		Attribute("unit_id", UUID)
		Attribute("votes", Integer)
		Attribute("answered", Boolean, "Whether the current user has answered the question.")

		Required("id", "question_type", "author_id", "unit_id", "votes", "answered")
	})

	View("default", func() {
		Attribute("id")
		Attribute("question_type")
		Attribute("author_id")
		Attribute("unit_id")
		Attribute("votes")
		Attribute("answered")
	})
	View("by_unit", func() {
		Attribute("id")
		Attribute("question_type")
		Attribute("author_id")
		Attribute("votes")
		Attribute("answered")
	})
	View("by_author", func() {
		Attribute("id")
		Attribute("question_type")
		Attribute("unit_id")
		Attribute("votes")
		Attribute("answered")
	})
	View("by_answered", func() {
		Attribute("id")
		Attribute("question_type")
		Attribute("author_id")
		Attribute("unit_id")
		Attribute("votes")
	})
	View("by_type", func() {
		Attribute("id")
		Attribute("author_id")
		Attribute("unit_id")
		Attribute("votes")
		Attribute("answered")
	})
})
