package models

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	// Test with known date
	dob := "1990-05-10"
	expectedAge := time.Now().Year() - 1990

	age, err := CalculateAge(dob)
	if err != nil {
		t.Fatalf("CalculateAge failed: %v", err)
	}

	if age != expectedAge && age != expectedAge-1 {
		t.Errorf("Expected age around %d, got %d", expectedAge, age)
	}
}
