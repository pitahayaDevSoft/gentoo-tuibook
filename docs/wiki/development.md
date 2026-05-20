# Development Guide

## Prerequisites

- Go 1.26.2 or later
- A terminal with true-color support

## Local Setup

1. Clone the repository.
2. `go mod download`

## Useful Commands

- `go run .` — Run the TUI application.
- `go build -o gentoo-tuibook.exe .` — Build binary.
- `cd scripts && go run fetch.go` — Fetch 4 handbook chapters.
- `cd scripts && go run fetch_full.go` — Fetch all 11 handbook chapters.
- `python build.py` — Regenerate `main.go` from the alternate (green theme) version.
