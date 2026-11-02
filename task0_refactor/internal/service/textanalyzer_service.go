// Package service provides business logic services.
//
// TextAnalyzerService wraps the textanalyzer package with logging capabilities.
// It implements the TextAnalyzer interface for easy mocking in tests.
package service

import (
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/textanalyzer"
)

// TextAnalyzer defines the interface for text analysis operations.
// This allows easy mocking and substitution in tests.
type TextAnalyzer interface {
	WordCount(text string) (int, error)
	CharCount(text string) int
}

// TextAnalyzerService wraps textanalyzer operations with additional capabilities
// such as logging or metrics (extensible for future enhancements).
type TextAnalyzerService struct {
	// future: add logger, metrics collector, etc.
}

// NewTextAnalyzerService creates a new TextAnalyzerService instance.
func NewTextAnalyzerService() *TextAnalyzerService {
	return &TextAnalyzerService{}
}

// WordCount returns the count of words in text.
// Returns an error if text is empty (wrapped ErrEmptyText).
func (tas *TextAnalyzerService) WordCount(text string) (int, error) {
	return textanalyzer.WordCount(text)
}

// CharCount returns the count of characters (runes) in text,
// excluding leading/trailing whitespace.
func (tas *TextAnalyzerService) CharCount(text string) int {
	return textanalyzer.CharCount(text)
}
