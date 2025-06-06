package main

import (
	"3subs/internal"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	banner := `
============================================
   3subs - Subdomain Enumeration Automation
   By: 0xjar (github.com/0x1jar)
   This tool automates subdomain enumeration
   using subfinder, assetfinder, and findomain.
============================================
`
	fmt.Println(banner)

	domain := ""
	flag.StringVar(&domain, "d", "", "Target domain name")
	flag.Usage = func() {
		fmt.Println("Usage: ./3subs -d <domain>")
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
	if err := internal.RunSubfinder(domain); err != nil {
		log.Fatalf("Error running subfinder: %v", err)
	}
	fmt.Println("[√] subfinder done.")

	fmt.Println("[>] Running assetfinder...")
	if err := internal.RunAssetfinder(domain); err != nil {
		log.Fatalf("Error running assetfinder: %v", err)
	}
	fmt.Println("[√] assetfinder done.")

	fmt.Println("[>] Running findomain...")
	if err := internal.RunFindomain(domain); err != nil {
		log.Fatalf("Error running findomain: %v", err)
	}
	fmt.Println("[√] findomain done.")

	fmt.Println("[>] Merging results and removing duplicates...")
	files := []string{"subdomains_subfinder.txt", "subdomains_assetfinder.txt", "subdomain_find.txt"}
	if err := internal.MergeSubdomains(files, "subdomains_all.txt"); err != nil {
		log.Fatalf("Error merging subdomains: %v", err)
	}
	fmt.Println("[√] Merge complete.")

	dur := time.Since(start)
	fmt.Printf("[✓] Subdomain enumeration completed in %s.\n", dur)
}
