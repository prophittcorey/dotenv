// Package dotenv loads local .env files into the environment.
package dotenv

var (
	// ScanPath are the directories that will be scanned for .env files.
	ScanPath = []string{"./"}
)

// Load scans the local environment for .env files and loads each one that is
// found.
func Load() error {
	return nil
}
