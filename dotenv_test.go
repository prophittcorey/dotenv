package dotenv

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	ScanFiles = []string{"testfiles/env", "testfiles/env-alt"}

	if err := Load(); err != nil {
		t.Fatal("failed to load testfiles/env")
	}

	if os.Getenv("DOTENV_ENV") != "development" {
		t.Fatalf("failed to set DOTENV_ENV; got %s", os.Getenv("DOTENV_ENV"))
	}

	if os.Getenv("TEST_KEY") != "123-abc" {
		t.Fatalf("failed to set TEST_KEY; got %s", os.Getenv("TEST_KEY"))
	}

	if os.Getenv("ALT_ENV") != "abc-def" {
		t.Fatalf("failed to set ALT_ENV; got %s", os.Getenv("ALT_ENV"))
	}

	if os.Getenv("CHAR_TEST") != "#^&$_./\\#" {
		t.Fatalf("failed to set CHAR_TEST; got %s", os.Getenv("CHAR_TEST"))
	}
}
