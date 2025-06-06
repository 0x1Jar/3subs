package internal

import (
	"io/ioutil"
	"strings"
)

// MergeSubdomains merges subdomain files into a single output file and removes duplicates.
func MergeSubdomains(files []string, outputFile string) error {
	subdomains := make(map[string]struct{})

	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if line != "" {
				subdomains[line] = struct{}{}
			}
		}
	}

	var uniqueSubdomains []string
	for subdomain := range subdomains {
		uniqueSubdomains = append(uniqueSubdomains, subdomain)
	}

	return ioutil.WriteFile(outputFile, []byte(strings.Join(uniqueSubdomains, "\n")), 0644)
}
