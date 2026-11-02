// Package handlers provides HTTP-style request handlers for demos.
//
// DemoHandler wraps demo service calls in a handler pattern,
// making it easy to add middleware or additional processing in the future.
package handlers

import (
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/service"
)

// DemoHandler handles demo-related requests/operations.
// It acts as an interface between main.go and business logic (DemoService).
// This layer allows for future middleware, request validation, or response formatting.
type DemoHandler struct {
	demoService *service.DemoService
}

// NewDemoHandler creates a new DemoHandler with its dependencies.
func NewDemoHandler(demoService *service.DemoService) *DemoHandler {
	return &DemoHandler{
		demoService: demoService,
	}
}

// HandleRunAllDemos executes all demos orchestrated by DemoService.
// Returns an error if any demo fails.
func (dh *DemoHandler) HandleRunAllDemos() error {
	return dh.demoService.RunAll()
}
