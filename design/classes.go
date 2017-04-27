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
		Response(OK, func() {
			Status(200)
			Media(CollectionOf("application/studiously.class+json"))
		})
		Response(NotFound, func() {
			Status(404)
			Description("Class does not exist or the user does not have access to it")
		})
	})

	Action("show", func() {
		Description("Get class by ID")
		Routing(GET("/:class_id"))
		Params(func() {
			Param("class_id", UUID, "Class ID")
		})
		Response(OK, func() {
			Status(200)
			Media(ClassMedia)
		})
		Response(NotFound, func() {
			Status(404)
			Description("Class does not exist or the user does not have access to it")
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
			Param("answered", Boolean, "Filter by whether the question has been answered by the member")
		})
		Response(OK, func() {
			Status(200)
			Media(CollectionOf("application/studiously.question+json"), "feed")
		})
		Response(BadRequest, func() {
			Status(400)
			Description("A query parameter is invalid")
		})
		Response(NotFound, func() {
			Status(404)
			Description("Class does not exist or the user does not have access to it")
		})
		//Response("UnknownUnit", func() {
		//	Description("Unit does not exist in the context of the class")
		//	Status(400)
		//})
		//Response("UnknownAuthor", func() {
		//	Description("Author does not exist in the context of the class")
		//	Status(400)
		//})
		//Response("UnknownType", func() {
		//	Description("Question type is not recognized")
		//	Status(400)
		//})
	})

	Action("show_members", func() {
		Description("Get members of a class")
		Routing(GET("/:class_id/members"))
		Params(func() {
			Param("class_id", UUID, "Class ID")
		})
		Response(OK, func() {
			Status(200)
			Media(CollectionOf("application/studiously.member+json"))
		})
		Response(NotFound, func() {
			Status(404)
			Description("Class does not exist or the user does not have access to it")
		})
	})
})

var ClassMedia = MediaType("application/studiously.class+json", func() {
	Attributes(func() {
		Attribute("id", UUID)
		Attribute("name", String)
		Attribute("current_unit", UUID, "Current unit of study")

		Required("id", "name")
	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("current_unit")
	})
})

var MemberMedia = MediaType("application/studiously.member+json", func() {
	Attributes(func() {
		Attribute("id", UUID)
		Attribute("name", String)
		Attribute("role", String, func() {
			Enum("student", "moderator", "teacher", "administrator")
		})
		Required("id", "name", "role")
	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("role")
	})
})
