package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/0x1Jar/3subs/internal"
)

func main() {
	banner := `
============================================
   3subs - Subdomain Enumeration Automation
   By: 0xjar (github.com/0x1Jar)
   This tool automates subdomain enumeration
   using subfinder, assetfinder, and findomain.
============================================
`
	fmt.Println(banner)

	domain := ""
	output := "subdomains_all.txt"
	flag.StringVar(&domain, "d", "", "Target domain name")
	flag.StringVar(&output, "o", "subdomains_all.txt", "Output file name for merged results")
	flag.Usage = func() {
		fmt.Println("Usage: 3subs -d <domain> [-o <output_file>]")
		flag.PrintDefaults()
	}
	flag.Parse()

	if domain == "" {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("[~] Starting subdomain enumeration for: %s\n", domain)
	start := time.Now()

	fmt.Println("[>] Running subfinder...")
	subfinderOut := "subdomains_subfinder.txt"
	if err := internal.RunSubfinder(domain, subfinderOut); err != nil {
		log.Fatalf("Error running subfinder: %v", err)
	}
	fmt.Println("[√] subfinder done.")

	fmt.Println("[>] Running assetfinder...")
	assetfinderOut := "subdomains_assetfinder.txt"
	if err := internal.RunAssetfinder(domain, assetfinderOut); err != nil {
		log.Fatalf("Error running assetfinder: %v", err)
	}
	fmt.Println("[√] assetfinder done.")

	fmt.Println("[>] Running findomain...")
	findomainOut := "subdomain_find.txt"
	if err := internal.RunFindomain(domain, findomainOut); err != nil {
		log.Fatalf("Error running findomain: %v", err)
	}
	fmt.Println("[√] findomain done.")

	fmt.Println("[>] Merging results and removing duplicates...")
	files := []string{subfinderOut, assetfinderOut, findomainOut}
	if err := internal.MergeSubdomains(files, output); err != nil {
		log.Fatalf("Error merging subdomains: %v", err)
	}
	fmt.Println("[√] Merge complete.")

	dur := time.Since(start)
	fmt.Printf("[✓] Subdomain enumeration completed in %s.\n", dur)
}
