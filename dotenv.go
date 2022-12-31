// Package dotenv loads local .env files into the environment.
package dotenv

import "fmt"

var (
	// ScanFiles a slice of files to look for. Each will be loaded.
	ScanFiles = []string{".env"}
)

// Load scans the local environment for .env files and loads each one that is
// found.
func Load() error {
	for _, file := range ScanFiles {
		fmt.Println(file)

		// TODO: Load file.
	}

	return nil
}
