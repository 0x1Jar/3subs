# 3subs - Subdomain Enumeration Automation

## Description
3subs is a CLI tool that automates subdomain enumeration using subfinder, assetfinder, and findomain. It merges and deduplicates results from all three tools for comprehensive coverage.

## Features
- Runs subfinder, assetfinder, and findomain automatically
- Merges and deduplicates subdomain results
- Simple CLI with progress banners
- Custom output filename with `-o` flag

## Requirements
- Go 1.18+
- macOS or Linux
- subfinder, assetfinder, and findomain installed and available in your PATH

## Installation
You can install 3subs directly using Go (the binary will be named `3subs` and typically placed in your `$GOPATH/bin`):
```sh
go install github.com/0x1Jar/3subs/cmd/3subs@latest
```

Or clone this repository and build the binary manually:
```sh
git clone https://github.com/0x1Jar/3subs.git
cd 3subs
go build -o 3subs ./cmd
```

After building the binary, you should move the `3subs` binary to a directory that is in your `$PATH` for easy access:
```sh
sudo mv 3subs /usr/local/bin/
```

### Dependencies Installation
You can use the provided script to install all required tools (subfinder, assetfinder, findomain) on macOS/Linux:
```sh
zsh install_tools.sh
```
- This will install `subfinder` and `assetfinder` to `$HOME/go/bin` (ensure this is in your PATH).
- On macOS, `findomain` is installed via Homebrew to `/usr/local/bin`.
- On Linux, `findomain` is downloaded and placed in `/usr/local/bin`.

#### Troubleshooting
- If you get a `command not found` error for any tool, ensure `$HOME/go/bin` and `/usr/local/bin` are in your `PATH`.
- For Go tools: Add this to your shell profile (e.g., `~/.zshrc` or `~/.bashrc`):
  ```sh
  export PATH="$HOME/go/bin:$PATH"
  ```
- For Homebrew (macOS):
  ```sh
  export PATH="/usr/local/bin:$PATH"
  ```

## Usage
```sh
3subs -d example.com -o myresults.txt
```
- Replace `example.com` with your target domain.
- Use `-o` to specify a custom output file for merged results (default: `subdomains_all.txt`).
- The results will be saved in:
  - `subdomains_subfinder.txt`
  - `subdomains_assetfinder.txt`
  - `subdomain_find.txt`
  - Merged and deduplicated results will be saved in your custom output file (default: `subdomains_all.txt`).

## Example Output
```
============================================
   3subs - Subdomain Enumeration Automation
   By: 0x1Jar (github.com/0x1Jar)
   This tool automates subdomain enumeration
   using subfinder, assetfinder, and findomain.
============================================
[~] Starting subdomain enumeration for: example.com
[>] Running subfinder...
[√] subfinder done.
[>] Running assetfinder...
[√] assetfinder done.
[>] Running findomain...
[√] findomain done.
[>] Merging results and removing duplicates...
[√] Merge complete.
[✓] Subdomain enumeration completed in 2.3s.
```

## License
MIT
