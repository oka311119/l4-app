package helpers

import (
	"testing"
)

func TestSaltGeneration(t *testing.T) {
	gen := NewSaltGenerator()
	salt, err := gen.Generate()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(salt) != 32 { // 16 bytes = 32 hex characters
		t.Fatalf("Expected 32 characters, got %d", len(salt))
	}
}
