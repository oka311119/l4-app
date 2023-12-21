package uuidgen

import (
	"testing"
)

func TestUUIDGenerate(t *testing.T) {
	gen := &UUID{}
	uuid := gen.V4()

	if len(uuid) != 36 { 
		t.Fatalf("Expected 36 characters, got %d", len(uuid))
	}
}
