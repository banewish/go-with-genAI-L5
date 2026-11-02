// Package app defines the application type and dependency wiring.
//
// Dependency hierarchy:
//   main.go
//     └─> App (application wiring)
//           ├─> DemoHandler
//           └─> Service layer
//                 ├─> CalculatorService
//                 ├─> TextAnalyzerService
//                 └─> DemoService
//                       └─> Domain packages (calculator, textanalyzer)
package app

import (
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/handlers"
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/service"
)

// App represents the complete application with all wired dependencies.
// It holds references to all services and handlers.
type App struct {
	DemoHandler *handlers.DemoHandler
	DemoService *service.DemoService
}

// New creates and wires a new App instance with all dependencies.
// All dependencies are created in the correct order to avoid circular dependencies.
func New() *App {
	// Layer 1: Create domain-specific services (wrap raw packages)
	calcService := service.NewCalculatorService()
	textService := service.NewTextAnalyzerService()

	// Layer 2: Create orchestration service (depends on domain services)
	demoService := service.NewDemoService(calcService, textService)

	// Layer 3: Create handlers (depends on services)
	demoHandler := handlers.NewDemoHandler(demoService)

	return &App{
		DemoHandler: demoHandler,
		DemoService: demoService,
	}
}
