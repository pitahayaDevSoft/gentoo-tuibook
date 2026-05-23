# gentoo-tuibook

TUI app for reading the Gentoo Linux installation handbook in the terminal.

## Build & run

```bash
# Build natively using the industrial Makefile
make build

# Install globally to /usr/local/bin
sudo make install

# Alternatively, run directly with Go
go run .
```

No test, lint, typecheck, or CI setup exists.

## Architecture

- Single `main` package (`main.go`), no packages or subdirectories.
- Uses `charmbracelet/bubbletea` (TUI framework), `glamour` (Markdown → terminal render), `lipgloss` (styling).
- Content is embedded via `//go:embed data/handbook/amd64/*.md`.
- Features an industrial `Makefile` and a custom Portage ebuild under `ebuild/` for Gentoo integration.
- Colors and theme settings are fully customizable dynamically via standard JSON configuration (`$XDG_CONFIG_HOME/gentoo-tuibook/config.json`).

## Gotchas

- **Data has two naming conventions**: numbered files (`01_about.md`, etc.) from `scripts/fetch.go`, unnumbered (`about.md`, etc.) from `scripts/fetch_full.go`. Both coexist in `data/handbook/amd64/`.

## Fetch scripts

Both are standalone `package main` scripts in `scripts/` that pull from the Gentoo wiki API and convert HTML to Markdown:

- `scripts/fetch.go` — fetches 4 specific pages (about, media, network, disks).
- `scripts/fetch_full.go` — fetches all 11 handbook chapters.
- Run from `scripts/` directory: `go run fetch.go` (or `fetch_full.go`).

## Active skills

Locked in `skills-lock.json`: `golang-patterns`, `golang-testing` (both from `affaan-m/everything-claude-code` registry).
