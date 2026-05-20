# gentoo-tuibook

TUI app for reading the Gentoo Linux installation handbook in the terminal.

## Build & run

```bash
go build -o gentoo-tuibook.exe .
go run .
```

No test, lint, typecheck, or CI setup exists.

## Architecture

- Single `main` package (`main.go`), no packages or subdirectories.
- Uses `charmbracelet/bubbletea` (TUI framework), `glamour` (Markdown → terminal render), `lipgloss` (styling).
- Content is embedded via `//go:embed data/handbook/amd64/*.md`.
- Two generations of the app exist: the current `main.go` (Gentoo pink theme, link-extraction mode) and an alternative in `build.py` (green theme, filtering, scroll helpers).

## Gotchas

- **`build.py` is codegen**: it overwrites `main.go` with a different version of the app. If you edit `main.go`, you must also update the embedded string in `build.py` or the change is lost on next `build.py` run.
- **No `.gitignore`**: binary `gentoo-tuibook.exe` is checked in.
- **Data has two naming conventions**: numbered files (`01_about.md`, etc.) from `scripts/fetch.go`, unnumbered (`about.md`, etc.) from `scripts/fetch_full.go`. Both coexist in `data/handbook/amd64/`.

## Fetch scripts

Both are standalone `package main` scripts in `scripts/` that pull from the Gentoo wiki API and convert HTML to Markdown:

- `scripts/fetch.go` — fetches 4 specific pages (about, media, network, disks).
- `scripts/fetch_full.go` — fetches all 11 handbook chapters.
- Run from `scripts/` directory: `go run fetch.go` (or `fetch_full.go`).

## Active skills

Locked in `skills-lock.json`: `golang-patterns`, `golang-testing` (both from `affaan-m/everything-claude-code` registry).
