package canontags_test

import (
	"testing"

	"github.com/evanoberholster/exiftools/mknote/canontags"
)

// Test CanonModelID Values
func TestCanonModel(t *testing.T) {
	t.Log(canontags.CanonModel(0x805))
	_, err := canontags.CanonModel(0x809)
	if err != nil {
		t.Fatalf("Could not find CanonModelID: %v", err)
	}

}

// Test CanonLensType Values
func TestCanonLens(t *testing.T) {
	lens1, err := canontags.CanonLens(138)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	if lens1[0] != "Canon EF 28-80mm f/2.8-4L" {
		t.Fatalf("Test Failed: Lens %v not found", lens1[0])
	}
	lens2, err := canontags.CanonLens(1)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	t.Log(lens2)
}
