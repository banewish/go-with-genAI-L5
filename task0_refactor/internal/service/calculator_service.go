// Package service provides business logic services.
//
// CalculatorService wraps the calculator package with logging capabilities.
// It implements the Calculator interface for easy mocking in tests.
package service

import (
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/calculator"
)

// Calculator defines the interface for calculator operations.
// This allows easy mocking and substitution in tests.
type Calculator interface {
	Add(a, b float64) float64
	Subtract(a, b float64) float64
	Multiply(a, b float64) float64
	Divide(a, b float64) (float64, error)
}

// CalculatorService wraps calculator operations with additional capabilities
// such as logging or metrics (extensible for future enhancements).
type CalculatorService struct {
	// future: add logger, metrics collector, etc.
}

// NewCalculatorService creates a new CalculatorService instance.
func NewCalculatorService() *CalculatorService {
	return &CalculatorService{}
}

// Add returns the sum of a and b.
func (cs *CalculatorService) Add(a, b float64) float64 {
	return calculator.Add(a, b)
}

// Subtract returns the difference a - b.
func (cs *CalculatorService) Subtract(a, b float64) float64 {
	return calculator.Subtract(a, b)
}

// Multiply returns the product of a and b.
func (cs *CalculatorService) Multiply(a, b float64) float64 {
	return calculator.Multiply(a, b)
}

// Divide returns the quotient a / b.
// Returns an error if b is zero (wrapped ErrDivisionByZero).
func (cs *CalculatorService) Divide(a, b float64) (float64, error) {
	return calculator.Divide(a, b)
}
