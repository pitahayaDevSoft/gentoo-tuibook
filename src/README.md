# Source

This Go project uses a flat `main` package at the repository root (`main.go`).

No `src/` subdirectory is required — the single-binary entrypoint lives at:

- `../main.go` — Application entrypoint (Bubbletea TUI)
- `../build.py` — Alternative codegen version of the app
- `../scripts/fetch.go` / `fetch_full.go` — Wiki content fetch scripts
