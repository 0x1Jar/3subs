# 3subs - Subdomain Enumeration Automation

## Description
3subs is a CLI tool that automates subdomain enumeration using subfinder, assetfinder, and findomain. It merges and deduplicates results from all three tools for comprehensive coverage.

## Features
- Runs subfinder, assetfinder, and findomain automatically
- Merges and deduplicates subdomain results
- Simple CLI with progress banners

## Requirements
- Go 1.18+
- macOS or Linux
- subfinder, assetfinder, and findomain installed and available in your PATH

## Installation
You can install 3subs directly using Go:
```sh
go install github.com/0x1Jar/3subs@latest
```

Or clone this repository and build the binary manually:
```sh
git clone https://github.com/0x1Jar/3subs.git
cd 3subs
go build -o 3subs ./cmd
```

### Dependencies Installation
You can use the provided script to install all required tools (subfinder, assetfinder, findomain) on macOS/Linux:
```sh
zsh install_tools.sh 0r ./install_tools.sh
```
This will install all dependencies and place them in your PATH (Go tools in `$HOME/go/bin`, findomain in `/usr/local/bin`).

## Usage
```sh
./3subs -d example.com
```
- Replace `example.com` with your target domain.
- The results will be saved in:
  - `subdomains_subfinder.txt`
  - `subdomains_assetfinder.txt`
  - `subdomain_find.txt`
  - Merged and deduplicated: `subdomains_all.txt`


## License
MIT