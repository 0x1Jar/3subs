package internal

import (
	"bytes"
	"fmt"
	"os/exec"
)

// RunFindomain executes the findomain command for the given domain
// and saves the output to a specified file.
func RunFindomain(domain string) error {
	cmd := exec.Command("findomain", "--unique-output", "subdomain_find.txt", "--target", domain)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run findomain: %v, stderr: %s", err, stderr.String())
	}
	return nil
}
