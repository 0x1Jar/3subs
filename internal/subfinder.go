package internal

import (
	"fmt"
	"os/exec"
)

// RunSubfinder executes the subfinder command for the given domain and saves the output to a specified file.
func RunSubfinder(domain string) error {
	cmd := exec.Command("subfinder", "-d", domain, "-o", "subdomains_subfinder.txt")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run subfinder: %w", err)
	}
	return nil
}
