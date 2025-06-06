package internal

import (
	"fmt"
	"os"
	"os/exec"
)

// RunAssetfinder executes the assetfinder command for the given domain and saves the output to a specified file.
func RunAssetfinder(domain string) error {
	cmd := exec.Command("assetfinder", "--subs-only", domain)
	outputFile := fmt.Sprintf("subdomains_assetfinder.txt")

	// Create the output file
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("could not create output file: %w", err)
	}
	defer file.Close()

	// Set the command's output to the file
	cmd.Stdout = file

	// Run the command
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running assetfinder: %w", err)
	}

	return nil
}
