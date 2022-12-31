package dotenv

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	ScanFiles = []string{"testfiles/env"}

	if err := Load(); err != nil {
		t.Fatal("failed to load testfiles/env")
	}

	if os.Getenv("DOTENV_ENV") != "development" {
		t.Fatalf("failed to set DOTENV_ENV; got %s", os.Getenv("DOTENV_ENV"))
	}

	if os.Getenv("TEST_KEY") != "abc-123" {
		t.Fatalf("failed to set TEST_KEY; got %s", os.Getenv("TEST_KEY"))
	}
}
