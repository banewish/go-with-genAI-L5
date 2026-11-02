// Command app runs the demo application.
//
// This thin main.go serves only to:
// 1. Create and wire the application (dependency injection)
// 2. Call the appropriate handlers
// 3. Handle top-level errors
//
// All business logic is organized in internal/ packages following a layered architecture:
//   - app/: Application wiring and initialization
//   - handlers/: Request handlers (entry points for operations)
//   - service/: Business logic layer with demo, calculator, and text analyzer services
//   - calculator/ & textanalyzer/: Domain packages with core functionality
package main

import (
	"fmt"
	"os"

	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/app"
)

func main() {
	// Step 1: Wire all dependencies via App factory
	application := app.New()

	// Step 2: Execute the orchestrated demos via handler
	if err := application.DemoHandler.HandleRunAllDemos(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
