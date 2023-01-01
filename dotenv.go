// Package dotenv loads local .env files into the environment.
package dotenv

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	// ScanFiles a slice of files to look for. Each will be loaded.
	ScanFiles = []string{".env"}
)

// Load scans the local environment for .env files and loads each one that is
// found.
func Load() error {
	var (
		quoteFixer = regexp.MustCompile(`^"(.*)"$`)
	)

	for _, envFile := range ScanFiles {
		if stat, err := os.Stat(envFile); err != nil || stat.IsDir() {
			continue /* skip if it's a directory, or no file is found */
		}

		err := (func(f string) error {
			file, err := os.Open(filepath.Clean(f))

			if err != nil {
				return fmt.Errorf("main.go: failed to open %s; %s", f, err)
			}

			defer (func() {
				if err := file.Close(); err != nil {
					log.Printf("main.go: failed to close file %s; %s\n", f, err)
				}
			})()

			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())

				if strings.HasPrefix(line, "#") {
					continue
				}

				before, after, found := strings.Cut(line, "=")

				if found {
					if err := os.Setenv(before, quoteFixer.ReplaceAllString(after, `$1`)); err != nil {
						return fmt.Errorf("dotenv: failed to set env value %s; %s", before, err)
					}
				}
			}

			if err := scanner.Err(); err != nil {
				return fmt.Errorf("dotenv: failed to load file %s; %s", f, err)
			}

			return nil
		})(envFile)

		if err != nil {
			return err
		}
	}

	return nil
}
