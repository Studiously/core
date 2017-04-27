//go:generate goagen bootstrap -d github.com/studiously/core/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/studiously/core/app"
)

func main() {
	// Create service
	service := goa.New("studiously")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "class" controller
	c := NewClassController(service)
	app.MountClassController(service, c)
	// Mount "question" controller
	c2 := NewQuestionController(service)
	app.MountQuestionController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
