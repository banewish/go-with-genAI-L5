// Package service provides business logic services.
//
// DemoService orchestrates demo workflows using calculator and text analyzer services.
// It contains the core business logic that was previously in main.go's run() function.
package service

import (
	"errors"
	"fmt"

	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/calculator"
)

// DemoService orchestrates demo operations using calculator and textanalyzer services.
// It encapsulates all the demo logic that was previously scattered in main.go.
type DemoService struct {
	calculatorService   Calculator
	textAnalyzerService TextAnalyzer
}

// NewDemoService creates a new DemoService with its dependencies.
func NewDemoService(calc Calculator, textAnalyzer TextAnalyzer) *DemoService {
	return &DemoService{
		calculatorService:   calc,
		textAnalyzerService: textAnalyzer,
	}
}

// RunDivisionDemo demonstrates division operation with error handling.
// Returns an error if division fails (except for handled cases).
func (ds *DemoService) RunDivisionDemo() error {
	quotient, err := ds.calculatorService.Divide(10, 0)
	if err != nil {
		// Check if it's the expected division by zero error
		if errors.Is(err, calculator.ErrDivisionByZero) {
			fmt.Println("division by zero was correctly detected:", err)
			return nil
		}
		return fmt.Errorf("division demo: %w", err)
	}
	fmt.Println("10 / 0 =", quotient)
	return nil
}

// RunArithmeticDemo demonstrates basic arithmetic operations.
// Does not return an error; prints result to stdout.
func (ds *DemoService) RunArithmeticDemo() {
	sum := ds.calculatorService.Add(2, 3)
	fmt.Println("2 + 3 =", sum)
}

// RunWordCountDemo demonstrates word counting with error handling.
// Returns an error if word count fails.
func (ds *DemoService) RunWordCountDemo() error {
	words, err := ds.textAnalyzerService.WordCount("the quick brown fox")
	if err != nil {
		return fmt.Errorf("word count demo: %w", err)
	}
	fmt.Println("word count:", words)
	return nil
}

// RunAll orchestrates all demo workflows in sequence.
// Returns an error if any demo fails (except for expected errors like division by zero).
func (ds *DemoService) RunAll() error {
	ds.RunArithmeticDemo()

	if err := ds.RunDivisionDemo(); err != nil {
		return fmt.Errorf("demo service: %w", err)
	}

	if err := ds.RunWordCountDemo(); err != nil {
		return fmt.Errorf("demo service: %w", err)
	}

	return nil
}
